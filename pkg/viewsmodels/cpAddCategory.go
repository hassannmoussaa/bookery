package viewsmodels

import (
	"github.com/hassannmoussaa/bookery/pkg/appCtx"
	"github.com/hassannmoussaa/bookery/pkg/textualContent"
)

type CPAddCategory struct {
	*CPPage
}

func GetCPAddCategory(ctx *appCtx.Context) *CPAddCategory {
	cpAddCategory := &CPAddCategory{}
	cpAddCategory.CPPage = NewCPPage("cp-addcat", ctx, true)
	cpAddCategory.Page.SetMetas(textualContent.OfTitle("addcat"), textualContent.OfTitle("addcat"))
	return cpAddCategory
}
