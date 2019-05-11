package viewsmodels

import (
	"github.com/hassannmoussaa/bookery/pkg/appCtx"
	"github.com/hassannmoussaa/bookery/pkg/models"
	"github.com/hassannmoussaa/bookery/pkg/textualContent"
)

type CPUsers struct {
	*CPPage
	Users      []*models.User
	Pagination *models.Pagination
}

func GetCPUsers(ctx *appCtx.Context) *CPUsers {
	cpUsers := &CPUsers{}
	cpUsers.CPPage = NewCPPage("cp-users", ctx, true)
	cpUsers.Page.SetMetas(textualContent.OfTitle("users"), textualContent.OfTitle("users"))
	users, _, nextPagesCount := models.GetUsers(ctx.QueryParams.Page, 4, 0)
	if ctx.QueryParams.Page <= 0 {
		ctx.QueryParams.Page = 1
	}
	pagination := models.NewPagination(ctx.QueryParams.Page, nextPagesCount)
	cpUsers.Pagination = pagination
	cpUsers.Users = users
	return cpUsers
}
