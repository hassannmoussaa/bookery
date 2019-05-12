package webCtrls

import (
	"github.com/hassannmoussaa/bookery/pkg/appCtx"
	"github.com/hassannmoussaa/bookery/pkg/viewsmodels"
	"github.com/hassannmoussaa/pill.go/fastmux"
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
func (this *CPCtrl) Users(requestCtx *fasthttp.RequestCtx) {
	ctx := appCtx.Get(requestCtx)
	search := fastmux.GetParam(requestCtx, "search")
	vm := viewsmodels.GetCPUsers(ctx, search)
	this.Write(requestCtx, vm)
}
func (this *CPCtrl) Orders(requestCtx *fasthttp.RequestCtx) {
	ctx := appCtx.Get(requestCtx)
	vm := viewsmodels.GetCPOrders(ctx)
	this.Write(requestCtx, vm)
}
func (this *CPCtrl) Categories(requestCtx *fasthttp.RequestCtx) {
	ctx := appCtx.Get(requestCtx)
	vm := viewsmodels.GetCPCategories(ctx)
	this.Write(requestCtx, vm)
}
func (this *CPCtrl) AddCategory(requestCtx *fasthttp.RequestCtx) {
	ctx := appCtx.Get(requestCtx)
	vm := viewsmodels.GetCPAddCategory(ctx)
	this.Write(requestCtx, vm)
}
func (this *CPCtrl) Books(requestCtx *fasthttp.RequestCtx) {
	ctx := appCtx.Get(requestCtx)
	search := fastmux.GetParam(requestCtx, "search")
	vm := viewsmodels.GetCPBooks(ctx, search)
	this.Write(requestCtx, vm)
}
func (this *CPCtrl) Admins(requestCtx *fasthttp.RequestCtx) {
	ctx := appCtx.Get(requestCtx)
	vm := viewsmodels.GetCPAdmins(ctx)
	this.Write(requestCtx, vm)
}
func (this *CPCtrl) AddAdmin(requestCtx *fasthttp.RequestCtx) {
	ctx := appCtx.Get(requestCtx)
	vm := viewsmodels.GetCPAddAdmin(ctx)
	this.Write(requestCtx, vm)
}
func (this *CPCtrl) CardOrders(requestCtx *fasthttp.RequestCtx) {
	ctx := appCtx.Get(requestCtx)
	vm := viewsmodels.GetCPCardOrders(ctx)
	this.Write(requestCtx, vm)
}
