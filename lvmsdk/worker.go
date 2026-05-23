package lvmsdk

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/tetratelabs/wazero/api"
	"strings"
)

type worker struct {
	addr   string
	token  string
	header Header
	mod    api.Module
}

func (w *worker) do(ctx context.Context, req *DoRequest) (*Response, error) {
	input, _ := json.Marshal(map[string]interface{}{
		"addr":  w.addr,
		"token": w.token,
		"header": func() map[string]string {
			m := map[string]string{}
			if w.header != nil {
				m = w.header(ctx)
			}
			if m == nil {
				m = map[string]string{}
			}
			if _, ok := m["X-Request-Id"]; !ok {
				m["X-Request-Id"] = strings.Replace(uuid.New().String(), "-", "", -1)
			}
			return m
		}(),
		"fn":      req.Fn,
		"content": req.Content,
	})
	results, err := w.mod.ExportedFunction("malloc").Call(ctx, uint64(len(input)))
	if err != nil {
		return nil, err
	}
	if !w.mod.Memory().Write(uint32(results[0]), input) {
		return nil, fmt.Errorf("write failed")
	}
	results, err = w.mod.ExportedFunction("do").Call(ctx, results[0], uint64(len(input)))
	if err != nil {
		return nil, err
	}
	output, ok := w.mod.Memory().Read(uint32(results[0]>>32), uint32(results[0]&0xFFFFFFFF))
	if !ok {
		return nil, fmt.Errorf("read failed")
	}
	rep := &Response{}
	if err := json.Unmarshal(output, rep); err != nil {
		return nil, err
	}
	return rep, nil
}

func (w *worker) call(ctx context.Context, req *CallRequest) (*Response, error) {
	input, _ := json.Marshal(map[string]interface{}{
		"addr":  w.addr,
		"token": w.token,
		"header": func() map[string]string {
			m := map[string]string{}
			if w.header != nil {
				m = w.header(ctx)
			}
			if m == nil {
				m = map[string]string{}
			}
			if _, ok := m["X-Request-Id"]; !ok {
				m["X-Request-Id"] = strings.Replace(uuid.New().String(), "-", "", -1)
			}
			return m
		}(),
		"fn":       req.Fn,
		"content":  req.Content,
		"function": req.Function,
		"params":   req.Params,
	})
	results, err := w.mod.ExportedFunction("malloc").Call(ctx, uint64(len(input)))
	if err != nil {
		return nil, err
	}
	if !w.mod.Memory().Write(uint32(results[0]), input) {
		return nil, fmt.Errorf("write failed")
	}
	results, err = w.mod.ExportedFunction("call").Call(ctx, results[0], uint64(len(input)))
	if err != nil {
		return nil, err
	}
	output, ok := w.mod.Memory().Read(uint32(results[0]>>32), uint32(results[0]&0xFFFFFFFF))
	if !ok {
		return nil, fmt.Errorf("read failed")
	}
	rep := &Response{}
	if err := json.Unmarshal(output, rep); err != nil {
		return nil, err
	}
	return rep, nil
}

func newWorker(addr string, token string, header Header, mod api.Module) *worker {
	return &worker{
		addr:   addr,
		token:  token,
		header: header,
		mod:    mod,
	}
}
