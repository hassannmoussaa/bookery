package viewsmodels

import (
	"github.com/hassannmoussaa/bookery/pkg/appCtx"
	"github.com/hassannmoussaa/bookery/pkg/models"
	"github.com/hassannmoussaa/bookery/pkg/textualContent"
)

type CPBooks struct {
	*CPPage
	Books      []*models.Book
	Pagination *models.Pagination
}

func GetCPBooks(ctx *appCtx.Context, search string) *CPBooks {
	var nextPagesCount int32
	var books []*models.Book
	cpBooks := &CPBooks{}
	cpBooks.CPPage = NewCPPage("cp-books", ctx, true)
	cpBooks.Page.SetMetas(textualContent.OfTitle("books"), textualContent.OfTitle("books"))
	if search == "null" || search == "" {
		books, _, nextPagesCount = models.GetBooks(ctx.QueryParams.Page, 4, 0)
	} else {
		books, _, nextPagesCount = models.GetSimilariBooks(ctx.QueryParams.Page, 4, 0, search)
	}
	if ctx.QueryParams.Page <= 0 {
		ctx.QueryParams.Page = 1
	}
	pagination := models.NewPagination(ctx.QueryParams.Page, nextPagesCount)
	cpBooks.Pagination = pagination
	cpBooks.Books = books
	return cpBooks
}
