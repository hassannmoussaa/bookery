package apiCtrls

import (
	"strconv"

	"github.com/hassannmoussaa/bookery/pkg/models"
	"github.com/hassannmoussaa/pill.go/fastmux"
	"github.com/hassannmoussaa/pill.go/helpers"
	"github.com/valyala/fasthttp"
)

type UserConnection struct {
	responsePath string
	toPage       string
}

func (this *UserConnection) ResponsePath() string {
	return this.responsePath
}
func (this *UserConnection) ToPage() string {
	return this.toPage
}
func (this *UserConnection) SetResponsePath(value string) {
	this.responsePath = value
}
func (this *UserConnection) SetToPage(value string) {
	this.toPage = value
}

type UserBotConnection struct {
	*APICtrl //inheritance
}

var UsersConnections = make(map[int]*UserConnection)

func (this *UserBotConnection) SetUserWantToSignUp(requestCtx *fasthttp.RequestCtx) {
	type data map[string]interface{}
	responsePath := helpers.BytesToString(requestCtx.PostArgs().Peek("ResponsePath"))
	userid, _ := strconv.Atoi(helpers.BytesToString(requestCtx.PostArgs().Peek("user_id")))
	UserConnection := &UserConnection{}
	UserConnection.responsePath = responsePath
	UserConnection.toPage = "signup"

	UsersConnections[userid] = UserConnection
	if responsePath == "" || userid == 0 {
		this.Fail(requestCtx, nil, "data_is_invaild", 400)
	} else {
		this.Success(requestCtx, data{"ResponsePath": UsersConnections[userid].responsePath, "UserName": models.GetUserById(int32(userid)).FullName(), "ToPage": "signup"}, "user_set_signup", 200)

	}

}
func (this *UserBotConnection) GetUserToPage(requestCtx *fasthttp.RequestCtx) {
	type data map[string]interface{}
	userid, _ := strconv.Atoi(fastmux.GetParam(requestCtx, "user_id"))
	if UsersConnections[userid] != nil {
		this.Success(requestCtx, data{"ResponsePath": UsersConnections[userid].responsePath, "UserName": models.GetUserById(int32(userid)).FullName(), "ToPage": UsersConnections[userid].toPage}, "user_status", 200)
	} else {
		this.Fail(requestCtx, nil, "no_to_page", 400)

	}

}

