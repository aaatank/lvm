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

func WithAsync(ctx context.Context, url string) context.Context {
	if url == "" {
		url = "no"
	}
	return context.WithValue(ctx, "lvm_callback_url", url)
}

func getCallbackUrl(ctx context.Context) string {
	if v := ctx.Value("lvm_callback_url"); v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}
