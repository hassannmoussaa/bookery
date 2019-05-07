package viewsmodels

import (
	"github.com/hassannmoussaa/bookery/pkg/appCtx"
)

type CPPage struct {
	*Page
}

func NewCPPage(pageName string, ctx *appCtx.Context, active ...interface{}) *CPPage {
	cpPage := &CPPage{}
	cpPage.Page = NewPage(pageName, ctx, active...)
	return cpPage
}
