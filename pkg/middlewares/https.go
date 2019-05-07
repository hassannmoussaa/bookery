package middlewares

import (
	"net/http"
	"strings"

	"github.com/hassannmoussaa/pill.go/helpers"
	"github.com/valyala/fasthttp"
)

func OnlyHttps(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	fn := func(requestCtx *fasthttp.RequestCtx) {
		host := helpers.BytesToString(requestCtx.Host())
		if onlyHttps && helpers.BytesToString(requestCtx.Request.Header.Peek("X-Forwarded-Proto")) != "https" && host != strings.Replace(webAddress, ":80", "", -1) && host != webAddress && host != strings.Replace(apiAddress, ":80", "", -1) && host != apiAddress {
			requestCtx.Redirect(
				"https://"+host+helpers.BytesToString(requestCtx.Path()),
				http.StatusMovedPermanently)
		} else {
			if next != nil {
				next(requestCtx)
			}
		}
	}
	return fn
}
