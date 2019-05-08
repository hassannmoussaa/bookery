package apiCtrls

import (
	"strconv"

	"github.com/hassannmoussaa/bookery/pkg/appCtx"
	"github.com/hassannmoussaa/bookery/pkg/models"
	"github.com/hassannmoussaa/pill.go/fastmux"
	"github.com/hassannmoussaa/pill.go/helpers"
	"github.com/valyala/fasthttp"
)

type BookCtrl struct {
	*APICtrl //inheritance
}

func (this *BookCtrl) Add(requestCtx *fasthttp.RequestCtx) {
	ctx := appCtx.Get(requestCtx)
	if book := ParseBookFromRequest(requestCtx); book != nil {
		if err := models.PrepareAndValidateBook(book); err == nil {
			userBook := &models.UserBook{}
			userBook.SetBook(models.GetBookById(book.ID()))
			userBook.SetUser(models.GetUserById(ctx.LoggedUser.ID()))
			if userBook := models.AddUserBook(userBook); userBook == nil {
				if book = models.AddBook(book); book != nil {
					this.Success(requestCtx, book, "new_book_was_added", 201)
				} else {
					this.Fail(requestCtx, nil, "book_cannot_be_added", 400)
				}
			} else {
				this.ServerError(requestCtx, "server_error")
			}
		} else {
			this.ValidationError(requestCtx, nil, err.Error())
		}
	} else {
		this.ServerError(requestCtx, "server_error")
	}
}

func (this *BookCtrl) Delete(requestCtx *fasthttp.RequestCtx) {
	bookId, _ := strconv.Atoi(fastmux.GetParam(requestCtx, "book_id"))
	if ok := models.DeleteUserBook(int32(bookId)); ok {
		if ok = models.DeleteOrderByBookId(int32(bookId)); ok {
			if err := models.DeleteBook(int32(bookId)); err {
				this.Success(requestCtx, nil, "book_was_deleted_successfully", 200)
			} else {
				this.Fail(requestCtx, nil, "book_cannot_be_deleted", 400)
			}
		} else {
			this.Fail(requestCtx, nil, "book_cannot_be_deleted", 400)
		}
	} else {
		this.Fail(requestCtx, nil, "book_cannot_be_deleted", 400)
	}
}
func ParseBookFromRequest(requestCtx *fasthttp.RequestCtx) *models.Book {
	book := &models.Book{}
	book.SetBookName(helpers.BytesToString(requestCtx.PostArgs().Peek("book_name")))
	book.SetAuthorName(helpers.BytesToString(requestCtx.PostArgs().Peek("author_name")))
	page_count, _ := strconv.Atoi(helpers.BytesToString(requestCtx.PostArgs().Peek("page_count")))
	book.SetPageCount(page_count)
	book.SetQuality(helpers.BytesToString(requestCtx.PostArgs().Peek("quality")))
	book.SetFrontImage(helpers.BytesToString(requestCtx.PostArgs().Peek("front_image")))
	book.SetBackImage(helpers.BytesToString(requestCtx.PostArgs().Peek("back_image")))
	book.SetSideImage(helpers.BytesToString(requestCtx.PostArgs().Peek("side_image")))
	book.SetIsRecived(false)
	book.SetIsVerified(false)
	category_id, _ := strconv.Atoi(helpers.BytesToString(requestCtx.PostArgs().Peek("category_id")))
	book.SetCategory(models.GetCategoryById(int32(category_id)))
	return book
}
