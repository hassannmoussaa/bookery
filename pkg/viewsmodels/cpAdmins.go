package viewsmodels

import (
	"github.com/hassannmoussaa/bookery/pkg/appCtx"
	"github.com/hassannmoussaa/bookery/pkg/models"
	"github.com/hassannmoussaa/bookery/pkg/textualContent"
)

type CPAdmins struct {
	*CPPage
	Admins []*models.Admin
}

func GetCPAdmins(ctx *appCtx.Context) *CPAdmins {
	cpAdmins := &CPAdmins{}
	cpAdmins.CPPage = NewCPPage("cp-admins", ctx, true)
	cpAdmins.Page.SetMetas(textualContent.OfTitle("admins"), textualContent.OfTitle("admins"))
	admins, _, _ := models.GetAdmins(ctx.QueryParams.Page, 4, 0)
	cpAdmins.Admins = admins
	return cpAdmins
}
