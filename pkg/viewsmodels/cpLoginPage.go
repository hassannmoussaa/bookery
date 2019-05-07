package viewsmodels

import (
	"github.com/hassannmoussaa/bookery/pkg/appCtx"
	"github.com/hassannmoussaa/bookery/pkg/textualContent"
)

type CPLoginPage struct {
	*Page
}

func GetCPLoginPage(ctx *appCtx.Context) *CPLoginPage {
	cpLoginPage := &CPLoginPage{}
	cpLoginPage.Page = NewPage("cp-login", ctx)
	cpLoginPage.Page.SetMetas(textualContent.OfTitle("login"), textualContent.OfTitle("login"))
	return cpLoginPage
}
