package viewsmodels

import (
	"github.com/hassannmoussaa/bookery/pkg/appCtx"
	"github.com/hassannmoussaa/bookery/pkg/textualContent"
)

type Index struct {
	*Page
}

func GetIndex(ctx *appCtx.Context) *Index {
	index := &Index{}
	index.Page = NewPage("index", ctx)
	index.Page.SetMetas(textualContent.OfTitle("index"), textualContent.OfTitle("index"))
	return index
}
