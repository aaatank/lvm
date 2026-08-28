package lvmsdk

import "context"

type proxy struct {
	addr  string
	token string
}

func WithProxy(ctx context.Context, addr string, token string) context.Context {
	return context.WithValue(ctx, "lvm_proxy", &proxy{
		addr:  addr,
		token: token,
	})
}

func getProxy(ctx context.Context) *proxy {
	if v := ctx.Value("lvm_proxy"); v != nil {
		if p, ok := v.(*proxy); ok {
			return p
		}
	}
	return nil
}
