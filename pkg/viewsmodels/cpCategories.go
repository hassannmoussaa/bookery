package viewsmodels

import (
	"github.com/hassannmoussaa/bookery/pkg/appCtx"
	"github.com/hassannmoussaa/bookery/pkg/models"
	"github.com/hassannmoussaa/bookery/pkg/textualContent"
)

type CPCategories struct {
	*CPPage
	Categories []*models.Category
}

func GetCPCategories(ctx *appCtx.Context) *CPCategories {
	cpCategories := &CPCategories{}
	cpCategories.CPPage = NewCPPage("cp-categories", ctx, true)
	cpCategories.Page.SetMetas(textualContent.OfTitle("categories"), textualContent.OfTitle("categories"))
	categories, _, _ := models.GetCategories(ctx.QueryParams.Page, 4, 0)
	cpCategories.Categories = categories
	return cpCategories
}
