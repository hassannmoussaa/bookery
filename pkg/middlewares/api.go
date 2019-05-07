package middlewares

import "github.com/valyala/fasthttp"

func APIProtector(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	fn := func(requestCtx *fasthttp.RequestCtx) {
		success := false
		if requestCtx.Request.Header.Peek("X-Api-Auth") != nil {
			success = APIServerBasicAuth(requestCtx)
		} else {
			success = CSRFProtector(requestCtx)
		}
		if next != nil && success {
			next(requestCtx)
		}
	}
	return fn
}
