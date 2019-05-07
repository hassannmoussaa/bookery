package viewsmodels

import (
	"github.com/hassannmoussaa/bookery/pkg/appCtx"
	"github.com/hassannmoussaa/bookery/pkg/textualContent"
)

type NotFoundPage struct {
	*Page
}

func GetNotFoundPage(ctx *appCtx.Context) *NotFoundPage {
	notFoundPage := &NotFoundPage{}
	notFoundPage.Page = NewPage("not-found", ctx)
	notFoundPage.Page.SetMetas(textualContent.OfTitle("page_not_found"), textualContent.OfTitle("page_not_found"))
	return notFoundPage
}
