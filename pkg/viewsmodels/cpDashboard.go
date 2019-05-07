package viewsmodels

import (
	"github.com/hassannmoussaa/bookery/pkg/appCtx"
	"github.com/hassannmoussaa/bookery/pkg/textualContent"
)

type CPDashboard struct {
	*CPPage
}

func GetCPDashboard(ctx *appCtx.Context) *CPDashboard {
	cpDashboard := &CPDashboard{}
	cpDashboard.CPPage = NewCPPage("cp-dashboard", ctx, true)
	cpDashboard.Page.SetMetas(textualContent.OfTitle("dashboard"), textualContent.OfTitle("dashboard"))
	return cpDashboard
}
