package middlewares

import "github.com/valyala/fasthttp"

func SetResponseContentTypeAsJSON(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	fn := func(requestCtx *fasthttp.RequestCtx) {
		requestCtx.SetContentType("application/json")
		if next != nil {
			next(requestCtx)
		}
	}
	return fn
}
