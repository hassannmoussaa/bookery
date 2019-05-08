package apiCtrls

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/hassannmoussaa/bookery/pkg/appCtx"
	"github.com/hassannmoussaa/bookery/pkg/models"
	"github.com/hassannmoussaa/pill.go/fastmux"
	"github.com/hassannmoussaa/pill.go/helpers"
	"github.com/valyala/fasthttp"
)

type CardOrderCtrl struct {
	*APICtrl //inheritance
}

func (this *CardOrderCtrl) Add(requestCtx *fasthttp.RequestCtx) {
	if cardorder := ParseCardOrderFromRequest(requestCtx); cardorder != nil {
		transaction := &models.Transaction{}
		transaction.SetUser(cardorder.User())
		transaction.SetUserOldCredit(cardorder.User().UserCredit())
		transaction.SetUserNewCredit(cardorder.User().UserCredit() + cardorder.Credit())
		transaction.SetDate(cardorder.Date())
		if transaction = models.AddTransaction(transaction); transaction != nil {
			if err := models.PrepareAndValidateCardOrder(cardorder); err == nil {
				if cardorder = models.AddCardOrder(cardorder); cardorder != nil {
					this.Success(requestCtx, cardorder, "new_card_order_was_added", 201)
				} else {
					this.Fail(requestCtx, nil, "card_order_cannot_be_added", 400)
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

func (this *CardOrderCtrl) Delete(requestCtx *fasthttp.RequestCtx) {
	cardorderid, _ := strconv.Atoi(fastmux.GetParam(requestCtx, "card_order_id"))
	if err := models.DeleteCardOrder(int32(cardorderid)); err {
		this.Success(requestCtx, nil, "card_order_was_deleted_successfully", 200)
	} else {
		this.Fail(requestCtx, nil, "card_order_cannot_be_deleted", 400)
	}
}

func ParseCardOrderFromRequest(requestCtx *fasthttp.RequestCtx) *models.CardOrder {
	cardorder := &models.CardOrder{}
	t := time.Now()
	ctx := appCtx.Get(requestCtx)
	cardorder.SetUser(models.GetUserById(ctx.LoggedUser.ID()))
	cardorder.SetCardNumber(rand.Intn(1000000000))
	credit, _ := strconv.Atoi(helpers.BytesToString(requestCtx.PostArgs().Peek("credit")))
	cardorder.SetCredit(credit)
	cardorder.SetDate(t.Format("Mon Jan _2 15:04:05 2006"))
	return cardorder
}
