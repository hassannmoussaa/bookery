package middlewares

import (
	"strings"

	"github.com/hassannmoussaa/pill.go/helpers"
	"github.com/valyala/fasthttp"
)

func AllowCORSRequest(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	fn := func(requestCtx *fasthttp.RequestCtx) {
		origin := helpers.BytesToString(requestCtx.Request.Header.Peek("Origin"))
		if isAllowedOrigin(origin) {
			requestCtx.Response.Header.Set("Access-Control-Allow-Origin", origin)
			requestCtx.Response.Header.Set("Access-Control-Allow-Methods", "HEAD,OPTIONS,POST,PUT,PATCH,GET,DELETE")
			requestCtx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
			requestCtx.Response.Header.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, X-Csrf-Token, X-Api-Auth")
		}
		if next != nil {
			next(requestCtx)
		}
	}
	return fn
}

func isAllowedOrigin(origin string) bool {
	if corsAllowedOrigins != nil {
		origin = strings.TrimSpace(strings.ToLower(origin))
		for _, o := range corsAllowedOrigins {
			if o == origin {
				return true
			}
		}
	}
	return false
}
