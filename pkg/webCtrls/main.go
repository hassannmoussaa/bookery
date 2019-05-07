package webCtrls

import (
	"bytes"
	"text/template"

	"github.com/hassannmoussaa/pill.go/templates"
	"github.com/valyala/fasthttp"
)

type WebCtrl struct{}

func (this *WebCtrl) Write(requestCtx *fasthttp.RequestCtx, viewmodel interface{}, options ...interface{}) {
	var statusCode int = 200
	if options != nil && len(options) > 0 {
		statusCode, _ = options[0].(int)
	}
	templateName := "index.html"
	if options != nil && len(options) > 1 {
		templateName, _ = options[1].(string)
	}
	var tmpl *template.Template
	if templateName != "" {
		tmpl = templates.GetTemplate(templateName)
	}
	if tmpl == nil {
		PageNotFound(requestCtx)
	} else {
		requestCtx.SetStatusCode(statusCode)
		requestCtx.SetContentType("text/html")
		var tpl bytes.Buffer
		tmpl.Execute(&tpl, viewmodel)
		requestCtx.Response.SetBodyStream(bytes.NewReader(tpl.Bytes()), tpl.Len())
	}
}

func Redirect(requestCtx *fasthttp.RequestCtx, url string) {
	requestCtx.Redirect(url, 307)
}

func PageNotFound(requestCtx *fasthttp.RequestCtx) {
	Redirect(requestCtx, "/404")
}
