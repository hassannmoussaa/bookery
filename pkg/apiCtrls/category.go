package apiCtrls

import (
	"strconv"

	"github.com/hassannmoussaa/bookery/pkg/models"
	"github.com/hassannmoussaa/pill.go/fastmux"
	"github.com/hassannmoussaa/pill.go/helpers"
	"github.com/valyala/fasthttp"
)

type CategoryCtrl struct {
	*APICtrl //inheritance
}

func (this *CategoryCtrl) Add(requestCtx *fasthttp.RequestCtx) {
	if category := ParseCategoryFromRequest(requestCtx); category != nil {
		if err := models.PrepareAndValidateCategory(category); err == nil {
			if category = models.AddCategory(category); category != nil {
				this.Success(requestCtx, category, "new_category_was_added", 201)
			} else {
				this.Fail(requestCtx, nil, "category_cannot_be_added", 400)
			}
		} else {
			this.ValidationError(requestCtx, nil, err.Error())
		}
	} else {
		this.ServerError(requestCtx, "server_error")
	}
}

func (this *CategoryCtrl) Delete(requestCtx *fasthttp.RequestCtx) {
	categoryId, _ := strconv.Atoi(fastmux.GetParam(requestCtx, "category_id"))
	if err := models.DeleteCategory(int32(categoryId)); err {
		this.Success(requestCtx, nil, "category_was_deleted_successfully", 200)
	} else {
		this.Fail(requestCtx, nil, "category_cannot_be_deleted", 400)
	}
}
func ParseCategoryFromRequest(requestCtx *fasthttp.RequestCtx) *models.Category {
	category := &models.Category{}
	category.SetCategoryName(helpers.BytesToString(requestCtx.PostArgs().Peek("category_name")))

	return category
}
