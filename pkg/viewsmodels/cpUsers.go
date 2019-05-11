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

func GetCPUsers(ctx *appCtx.Context, search string) *CPUsers {
	var nextPagesCount int32
	var users []*models.User
	cpUsers := &CPUsers{}
	cpUsers.CPPage = NewCPPage("cp-users", ctx, true)
	cpUsers.Page.SetMetas(textualContent.OfTitle("users"), textualContent.OfTitle("users"))
	if search == "null" || search == "" {
		users, _, nextPagesCount = models.GetUsers(ctx.QueryParams.Page, 4, 0)
	} else {
		users, _, nextPagesCount = models.GetSimilariUsers(ctx.QueryParams.Page, 4, 0, search)
	}
	if ctx.QueryParams.Page <= 0 {
		ctx.QueryParams.Page = 1
	}
	pagination := models.NewPagination(ctx.QueryParams.Page, nextPagesCount)
	cpUsers.Pagination = pagination
	cpUsers.Users = users
	return cpUsers
}
