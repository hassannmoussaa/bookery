package apiCtrls

import (
	"strconv"

	"github.com/hassannmoussaa/bookery/pkg/appCtx"
	"github.com/hassannmoussaa/bookery/pkg/models"
	"github.com/hassannmoussaa/pill.go/fastmux"
	"github.com/hassannmoussaa/pill.go/helpers"
	"github.com/valyala/fasthttp"
)

type UserCtrl struct {
	*APICtrl //inheritance
}

func (this *UserCtrl) Add(requestCtx *fasthttp.RequestCtx) {
	if user := ParseUserFromRequest(requestCtx); user != nil {
		if err := models.PrepareAndValidateUser(user); err == nil {
			if user = models.AddUser(user); user != nil {
				this.Success(requestCtx, user, "new_user_was_added", 201)
			} else {
				this.Fail(requestCtx, nil, "user_cannot_be_added", 400)
			}
		} else {
			this.ValidationError(requestCtx, nil, err.Error())
		}
	} else {
		this.ServerError(requestCtx, "server_error")
	}
}

func (this *UserCtrl) GetMe(requestCtx *fasthttp.RequestCtx) {
	context := appCtx.Get(requestCtx)
	fields, excluded := this.SetFields(requestCtx)
	data := context.LoggedUser.ToMap("", "", excluded, fields...)
	this.Success(requestCtx, data, "")
}

func (this *UserCtrl) Delete(requestCtx *fasthttp.RequestCtx) {
	userId, _ := strconv.Atoi(fastmux.GetParam(requestCtx, "user_id"))
	if ok := models.DeleteUserBookByUserId(int32(userId)); ok {
		if ok = models.DeleteOrderByUserId(int32(userId)); ok {
			if ok = models.DeleteTransactionByUserId(int32(userId)); ok {
				if ok = models.DeleteCardOrderByUserId(int32(userId)); ok {
					if err := models.DeleteUser(int32(userId)); err {
						this.Success(requestCtx, nil, "user_was_deleted_successfully", 200)
					} else {
						this.Fail(requestCtx, nil, "user_cannot_be_deleted", 400)
					}
				} else {
					this.Fail(requestCtx, nil, "user_cannot_be_deleted", 400)
				}
			} else {
				this.Fail(requestCtx, nil, "user_cannot_be_deleted", 400)
			}
		} else {
			this.Fail(requestCtx, nil, "user_cannot_be_deleted", 400)
		}

	} else {
		this.Fail(requestCtx, nil, "user_cannot_be_deleted", 400)
	}
}
func (this *UserCtrl) Block(requestCtx *fasthttp.RequestCtx) {
	userId, _ := strconv.Atoi(fastmux.GetParam(requestCtx, "user_id"))
	if err := models.BlockUser(models.GetUserById(int32(userId))); err {
		this.Success(requestCtx, nil, "user_was_blocked_successfully", 200)
	} else {
		this.Fail(requestCtx, nil, "user_cannot_be_blocked", 400)
	}
}
func (this *UserCtrl) UnBlock(requestCtx *fasthttp.RequestCtx) {
	userId, _ := strconv.Atoi(fastmux.GetParam(requestCtx, "user_id"))
	if err := models.UnBlockUser(models.GetUserById(int32(userId))); err {
		this.Success(requestCtx, nil, "user_was_unblocked_successfully", 200)
	} else {
		this.Fail(requestCtx, nil, "user_cannot_be_unblocked", 400)
	}
}

func ParseUserFromRequest(requestCtx *fasthttp.RequestCtx) *models.User {
	user := &models.User{}
	user.SetFullName(helpers.BytesToString(requestCtx.PostArgs().Peek("full_name")))
	user.SetEmail(helpers.BytesToString(requestCtx.PostArgs().Peek("email")))
	user.SetPassword(helpers.BytesToString(requestCtx.PostArgs().Peek("password")))
	user.SetFullAddress(helpers.BytesToString(requestCtx.PostArgs().Peek("full_address")))
	user.SetPhoneNumber(helpers.BytesToString(requestCtx.PostArgs().Peek("phone_number")))
	user.SetUserCredit(0)
	user.SetsBlocked(false)
	return user
}
