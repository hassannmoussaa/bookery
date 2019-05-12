package viewsmodels

import (
	"github.com/hassannmoussaa/bookery/pkg/appCtx"
	"github.com/hassannmoussaa/bookery/pkg/models"
	"github.com/hassannmoussaa/bookery/pkg/textualContent"
)

type CPCardOrders struct {
	*CPPage
	CardOrders []*models.CardOrder
	Pagination *models.Pagination
}

func GetCPCardOrders(ctx *appCtx.Context) *CPCardOrders {
	var nextPagesCount int32
	var cardorders []*models.CardOrder
	cpCardOrders := &CPCardOrders{}
	cpCardOrders.CPPage = NewCPPage("cp-cardorders", ctx, true)
	cpCardOrders.Page.SetMetas(textualContent.OfTitle("cardorders"), textualContent.OfTitle("cardorders"))
	cardorders, _, nextPagesCount = models.GetCardOrders(ctx.QueryParams.Page, 4, 0)
	if ctx.QueryParams.Page <= 0 {
		ctx.QueryParams.Page = 1
	}
	pagination := models.NewPagination(ctx.QueryParams.Page, nextPagesCount)
	cpCardOrders.Pagination = pagination
	cpCardOrders.CardOrders = cardorders
	return cpCardOrders
}
