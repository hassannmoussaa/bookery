package middlewares

import (
	"net/http"

	"github.com/hassannmoussaa/pill.go/antiCSRF"
	"github.com/hassannmoussaa/pill.go/helpers"
	"github.com/hassannmoussaa/bookery/pkg/appCtx"
	"github.com/hassannmoussaa/bookery/pkg/textualContent"
	"github.com/valyala/fasthttp"
)

func CSRFToken(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	fn := func(requestCtx *fasthttp.RequestCtx) {
		csrfToken := antiCSRF.GetFastHttpRequestCSRFToken(requestCtx)
		if csrfToken == nil {
			csrfToken = antiCSRF.NewCSRFToken()
			csrfToken.SetFastHttpCookie(requestCtx)
			csrfToken.Expired = true
		} else if csrfToken.IsExpired() {
			csrfToken = antiCSRF.NewCSRFToken()
			csrfToken.SetFastHttpCookie(requestCtx)
			csrfToken.Expired = true
		}

		ctx := appCtx.Get(requestCtx)
		ctx.CSRFToken = csrfToken

		if next != nil {
			next(requestCtx)
		}
	}
	return fn
}

func CSRFProtector(requestCtx *fasthttp.RequestCtx) bool {
	if !antiCSRF.IsSafeMethod(helpers.BytesToString(requestCtx.Method())) {
		ctx := appCtx.Get(requestCtx)
		if ctx.CSRFToken != nil {
			if ctx.CSRFToken.Expired {
				CSRFError(requestCtx)
				return false
			} else if !ctx.CSRFToken.IsValidRequestToken() {
				csrfToken := antiCSRF.NewCSRFToken()
				csrfToken.SetFastHttpCookie(requestCtx)
				CSRFError(requestCtx)
				return false
			}
		} else {
			CSRFError(requestCtx)
			return false
		}
	}
	return true
}

func CSRFError(requestCtx *fasthttp.RequestCtx) {
	if helpers.BytesToString(requestCtx.Response.Header.Peek("Content-Type")) == "application/json" {
		requestCtx.SetStatusCode(http.StatusForbidden)
		requestCtx.Write([]byte(`{"status": "error", "message": "` + textualContent.OfErrorMsg("csrf_error") + `"}`))
	} else {
		requestCtx.Redirect("/404", 307)
	}
	return
}
