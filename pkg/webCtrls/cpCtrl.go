package webCtrls

import (
	"github.com/hassannmoussaa/bookery/pkg/appCtx"
	"github.com/hassannmoussaa/bookery/pkg/viewsmodels"
	"github.com/valyala/fasthttp"
)

type CPCtrl struct {
	*WebCtrl
}

func (this *CPCtrl) Login(requestCtx *fasthttp.RequestCtx) {
	ctx := appCtx.Get(requestCtx)
	if ctx.LoggedAdmin != nil {
		Redirect(requestCtx, "/cp")
	} else {
		vm := viewsmodels.GetCPLoginPage(ctx)
		this.Write(requestCtx, vm)
	}
}

func (this *CPCtrl) Dashboard(requestCtx *fasthttp.RequestCtx) {
	ctx := appCtx.Get(requestCtx)
	vm := viewsmodels.GetCPDashboard(ctx)
	this.Write(requestCtx, vm)
}
