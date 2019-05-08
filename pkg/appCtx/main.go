package appCtx

import (
	"github.com/hassannmoussaa/bookery/pkg/models"
	"github.com/hassannmoussaa/fasthttpcontext"
	"github.com/hassannmoussaa/pill.go/antiCSRF"
	"github.com/hassannmoussaa/pill.go/helpers"
	"github.com/valyala/fasthttp"
)

type Context struct {
	LoggedAdmin *models.Admin
	LoggedUser  *models.User
	QueryParams *QueryParams
	URLPath     string
	CSRFToken   *antiCSRF.CSRFToken
}

type QueryParams struct {
	Page    int32
	Count   int32
	SinceID int32
	From    int32
	Search  string
	params  map[string]string
}

func (this *QueryParams) Get(paramName string) string {
	if this.params != nil && paramName != "" {
		if v, ok := this.params[paramName]; ok {
			return v
		}
	}
	return ""
}
func (this *QueryParams) Set(paramName string, value string) {
	if paramName != "" && value != "" {
		if this.params == nil {
			this.params = map[string]string{}
		}
		this.params[paramName] = value
	}
}

func Get(requestCtx *fasthttp.RequestCtx) *Context {
	if ctx, ok := fasthttpcontext.GetOk(requestCtx, "context"); ok {
		return ctx.(*Context)
	} else {
		ctx := &Context{}
		ctx.QueryParams = &QueryParams{}
		ctx.URLPath = helpers.BytesToString(requestCtx.Path())
		fasthttpcontext.Set(requestCtx, "context", ctx)
		return ctx
	}
}
func Clear(requestCtx *fasthttp.RequestCtx) {
	fasthttpcontext.Clear(requestCtx)
}
