package viewsmodels

import (
	"github.com/hassannmoussaa/bookery/pkg/appCtx"
	"github.com/hassannmoussaa/bookery/pkg/textualContent"
)

type CPAddAdmin struct {
	*CPPage
}

func GetCPAddAdmin(ctx *appCtx.Context) *CPAddAdmin {
	cpAddAdmin := &CPAddAdmin{}
	cpAddAdmin.CPPage = NewCPPage("cp-addadmin", ctx, true)
	cpAddAdmin.Page.SetMetas(textualContent.OfTitle("addadmin"), textualContent.OfTitle("addadmin"))
	return cpAddAdmin
}
