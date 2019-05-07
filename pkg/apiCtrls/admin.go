package apiCtrls

import (
	"github.com/hassannmoussaa/pill.go/helpers"
	"github.com/hassannmoussaa/bookery/pkg/appCtx"
	"github.com/hassannmoussaa/bookery/pkg/models"
	"github.com/valyala/fasthttp"
)

type AdminCtrl struct {
	*APICtrl //inheritance
}

func (this *AdminCtrl) Add(requestCtx *fasthttp.RequestCtx) {
	if admin := ParseAdminFromRequest(requestCtx); admin != nil {
		if err := models.PrepareAndValidateAdmin(admin); err == nil {
			if admin = models.AddAdmin(admin); admin != nil {
				this.Success(requestCtx, admin, "new_admin_was_added", 201)
			} else {
				this.Fail(requestCtx, nil, "admin_cannot_be_added", 400)
			}
		} else {
			this.ValidationError(requestCtx, nil, err.Error())
		}
	} else {
		this.ServerError(requestCtx, "server_error")
	}
}

func (this *AdminCtrl) GetMe(requestCtx *fasthttp.RequestCtx) {
	context := appCtx.Get(requestCtx)
	fields, excluded := this.SetFields(requestCtx)
	data := context.LoggedAdmin.ToMap("", excluded, fields...)
	this.Success(requestCtx, data, "")
}

func ParseAdminFromRequest(requestCtx *fasthttp.RequestCtx) *models.Admin {
	admin := &models.Admin{}
	admin.SetName(helpers.BytesToString(requestCtx.PostArgs().Peek("name")))
	admin.SetEmail(helpers.BytesToString(requestCtx.PostArgs().Peek("email")))
	admin.SetPassword(helpers.BytesToString(requestCtx.PostArgs().Peek("password")))
	return admin
}
