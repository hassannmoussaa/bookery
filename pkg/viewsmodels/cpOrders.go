package viewsmodels

import (
	"github.com/hassannmoussaa/bookery/pkg/appCtx"
	"github.com/hassannmoussaa/bookery/pkg/models"
	"github.com/hassannmoussaa/bookery/pkg/textualContent"
)

type CPOrders struct {
	*CPPage
	Orders     []*models.Order
	Pagination *models.Pagination
}

func GetCPOrders(ctx *appCtx.Context) *CPOrders {
	var nextPagesCount int32
	var orders []*models.Order
	cpOrders := &CPOrders{}
	cpOrders.CPPage = NewCPPage("cp-orders", ctx, true)
	cpOrders.Page.SetMetas(textualContent.OfTitle("orders"), textualContent.OfTitle("orders"))
	orders, _, nextPagesCount = models.GetOrders(ctx.QueryParams.Page, 4, 0)
	if ctx.QueryParams.Page <= 0 {
		ctx.QueryParams.Page = 1
	}
	pagination := models.NewPagination(ctx.QueryParams.Page, nextPagesCount)
	cpOrders.Pagination = pagination
	cpOrders.Orders = orders
	return cpOrders
}
