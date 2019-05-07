package webCtrls

import (
	"github.com/hassannmoussaa/pill.go/helpers"
	"github.com/hassannmoussaa/bookery/pkg/appCtx"
	"github.com/hassannmoussaa/bookery/pkg/models"
	"github.com/hassannmoussaa/bookery/pkg/viewsmodels"
	"github.com/valyala/fasthttp"
)

type AccountCtrl struct {
	*WebCtrl
}

func (this *AccountCtrl) UnlockAdmin(requestCtx *fasthttp.RequestCtx) {
	ctx := appCtx.Get(requestCtx)
	adminID := requestCtx.QueryArgs().GetUintOrZero("admin_id")
	hash := helpers.BytesToString(requestCtx.QueryArgs().Peek("hash"))
	success := models.UnlockAdminAccount(adminID, hash)
	vm := viewsmodels.GetUnlockAdminPage(ctx, success)
	this.Write(requestCtx, vm)
}
