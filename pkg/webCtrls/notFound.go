package webCtrls

import (
	"github.com/hassannmoussaa/bookery/pkg/appCtx"
	"github.com/hassannmoussaa/bookery/pkg/viewsmodels"
	"github.com/valyala/fasthttp"
)

type NotFoundCtrl struct {
	*WebCtrl
}

func (this *NotFoundCtrl) Get(requestCtx *fasthttp.RequestCtx) {
	ctx := appCtx.Get(requestCtx)
	vm := viewsmodels.GetNotFoundPage(ctx)
	this.Write(requestCtx, vm)
}
