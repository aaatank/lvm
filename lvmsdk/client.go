package lvmsdk

import (
	"context"
	"encoding/json"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	cli              *http.Client
	workers          chan *worker
	candidates       chan *worker
	memoryLimitPages uint32
}

func (c *Client) post(ctx context.Context, m api.Module, in_ptr uint32, in_len uint32) (out_ptr uint64) {
	req := struct {
		Url    string            `json:"url"`
		Header map[string]string `json:"header"`
		Body   string            `json:"body"`
	}{}
	var rep struct {
		Code   int         `json:"code"`
		Msg    string      `json:"msg"`
		Reason string      `json:"reason"`
		Data   interface{} `json:"data"`
	}
	defer func() {
		content, _ := json.Marshal(rep)
		results, err := m.ExportedFunction("malloc").Call(ctx, uint64(len(content)))
		if err != nil {
			panic(err)
		}
		if !m.Memory().Write(uint32(results[0]), content) {
			panic("write failed")
		}
		out_ptr = results[0]<<32 | uint64(len(content))
	}()
	content, ok := m.Memory().Read(in_ptr, in_len)
	if !ok {
		rep.Code = 500
		rep.Msg = "unknown error"
		rep.Reason = "read failed"
		return
	}
	if err := json.Unmarshal(content, &req); err != nil {
		rep.Code = 500
		rep.Msg = "unknown error"
		rep.Reason = err.Error()
		return
	}
	hreq, err := http.NewRequest("POST", req.Url, strings.NewReader(req.Body))
	if err != nil {
		rep.Code = 500
		rep.Msg = "unknown error"
		rep.Reason = err.Error()
		return
	}
	for k, v := range req.Header {
		hreq.Header.Set(k, v)
	}
	resp, err := c.cli.Do(hreq)
	if err != nil {
		rep.Code = 500
		rep.Msg = "unknown error"
		rep.Reason = err.Error()
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		content, _ := io.ReadAll(resp.Body)
		rep.Code = resp.StatusCode
		rep.Msg = "unknown error"
		rep.Reason = string(content)
		return
	}
	hrep := struct {
		Content string `json:"content"`
	}{}
	if err := json.NewDecoder(resp.Body).Decode(&hrep); err != nil {
		rep.Code = 500
		rep.Msg = "unknown error"
		rep.Reason = err.Error()
		return
	}
	rep.Data = hrep
	return
}

func (c *Client) Do(ctx context.Context, req *DoRequest) (*Response, error) {
	w := <-c.workers
	if w.mod.Memory().Size()/65536 > c.memoryLimitPages {
		w.mod.Close(ctx)
		w = <-c.candidates
	}
	defer func() {
		c.workers <- w
	}()
	return w.do(ctx, req)
}

func (c *Client) Call(ctx context.Context, req *CallRequest) (*Response, error) {
	w := <-c.workers
	if w.mod.Memory().Size()/65536 > c.memoryLimitPages {
		w.mod.Close(ctx)
		w = <-c.candidates
	}
	defer func() {
		c.workers <- w
	}()
	return w.call(ctx, req)
}

func NewClient(addr string, parallelism int, memoryLimitPages uint32) *Client {
	c := &Client{
		cli: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:        6000,
				MaxConnsPerHost:     1200,
				MaxIdleConnsPerHost: 1200,
			},
		},
		workers:          make(chan *worker, parallelism),
		candidates:       make(chan *worker, 1),
		memoryLimitPages: memoryLimitPages,
	}
	ctx := context.Background()
	r := wazero.NewRuntime(ctx)
	builder := r.NewHostModuleBuilder("net")
	builder.NewFunctionBuilder().WithFunc(c.post).Export("post")
	if _, err := builder.Instantiate(ctx); err != nil {
		panic(err.Error())
	}
	wasi_snapshot_preview1.MustInstantiate(ctx, r)
	config := wazero.NewModuleConfig().WithStartFunctions("_initialize").WithSysWalltime()
	for i := 0; i < parallelism; i++ {
		mod, err := r.InstantiateWithConfig(ctx, wasm, config)
		if err != nil {
			panic(err)
		}
		c.workers <- newWorker(addr, mod)
	}
	go func() {
		for {
			mod, err := r.InstantiateWithConfig(ctx, wasm, config)
			if err != nil {
				panic(err)
			}
			c.candidates <- newWorker(addr, mod)
		}
	}()
	return c
}
