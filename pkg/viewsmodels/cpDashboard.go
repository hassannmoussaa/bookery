package viewsmodels

import (
	"github.com/hassannmoussaa/bookery/pkg/appCtx"
	"github.com/hassannmoussaa/bookery/pkg/models"
	"github.com/hassannmoussaa/bookery/pkg/textualContent"
)

type CPDashboard struct {
	*CPPage
	UsersCount      int
	OrderCount      int
	BooksCount      int
	CategoriesCount int
	CardOrdersCount int
	AdminsCount     int
}

func GetCPDashboard(ctx *appCtx.Context) *CPDashboard {
	cpDashboard := &CPDashboard{}
	cpDashboard.CPPage = NewCPPage("cp-dashboard", ctx, true)
	cpDashboard.Page.SetMetas(textualContent.OfTitle("dashboard"), textualContent.OfTitle("dashboard"))
	cpDashboard.UsersCount = models.GetUsersCount()
	cpDashboard.OrderCount = models.GetOrdersCount()
	cpDashboard.BooksCount = models.GetBooksCount()
	cpDashboard.CategoriesCount = models.GetCategoriesCount()
	cpDashboard.CardOrdersCount = models.GetCardOrdersCount()
	cpDashboard.AdminsCount = models.GetAdminsCount()
	return cpDashboard
}
