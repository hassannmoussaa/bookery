package apiCtrls

import (
	"strconv"
	"time"

	"github.com/hassannmoussaa/bookery/pkg/appCtx"
	"github.com/hassannmoussaa/bookery/pkg/models"
	"github.com/hassannmoussaa/pill.go/fastmux"
	"github.com/hassannmoussaa/pill.go/helpers"
	"github.com/valyala/fasthttp"
)

type OrderCtrl struct {
	*APICtrl //inheritance
}

func (this *OrderCtrl) Add(requestCtx *fasthttp.RequestCtx) {
	if order := ParseOrderFromRequest(requestCtx); order != nil {
		transaction := &models.Transaction{}
		transaction.SetUser(order.User())
		transaction.SetUserOldCredit(order.User().UserCredit())
		transaction.SetUserNewCredit(order.User().UserCredit() - order.Book().Price())
		transaction.SetDate(order.Date())
		if transaction = models.AddTransaction(transaction); transaction != nil {
			if err := models.PrepareAndValidateOrder(order); err == nil {
				if order = models.AddOrder(order); order != nil {
					this.Success(requestCtx, order, "new_order_was_added", 201)
				} else {
					this.Fail(requestCtx, nil, "order_cannot_be_added", 400)
				}
			} else {
				this.ValidationError(requestCtx, nil, err.Error())

			}

		} else {
			this.Fail(requestCtx, nil, "order_cannot_be_added", 400)
		}
	} else {
		this.ServerError(requestCtx, "server_error")
	}
}

func (this *OrderCtrl) Delete(requestCtx *fasthttp.RequestCtx) {
	orderid, _ := strconv.Atoi(fastmux.GetParam(requestCtx, "order_id"))
	if err := models.DeleteOrder(int32(orderid)); err {
		this.Success(requestCtx, nil, "order_was_deleted_successfully", 200)
	} else {
		this.Fail(requestCtx, nil, "order_cannot_be_deleted", 400)
	}
}
func (this *OrderCtrl) Complete(requestCtx *fasthttp.RequestCtx) {
	orderid, _ := strconv.Atoi(fastmux.GetParam(requestCtx, "order_id"))
	if err := models.CompleteOrder(models.GetOrderById(int32(orderid))); err {
		this.Success(requestCtx, nil, "order_was_set_completed_successfully", 200)
	} else {
		this.Fail(requestCtx, nil, "order_cannot_be_set_completed", 400)
	}
}
func ParseOrderFromRequest(requestCtx *fasthttp.RequestCtx) *models.Order {
	order := &models.Order{}
	t := time.Now()
	ctx := appCtx.Get(requestCtx)
	bookid, _ := strconv.Atoi(helpers.BytesToString(requestCtx.PostArgs().Peek("book_id")))
	order.SetBook(models.GetBookById(int32(bookid)))
	order.SetUser(models.GetUserById(ctx.LoggedUser.ID()))
	order.SetOrderStatus("packed")
	order.SetDeliveryMethod(helpers.BytesToString(requestCtx.PostArgs().Peek("delivery_status")))
	order.SetDate(t.Format("Mon Jan _2 15:04:05 2006"))
	return order
}
