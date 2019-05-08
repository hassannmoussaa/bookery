package apiCtrls

import (
	"errors"
	"net/http"

	"github.com/hassannmoussaa/bookery/pkg/appCtx"
	"github.com/hassannmoussaa/bookery/pkg/models"
	"github.com/hassannmoussaa/pill.go/auth"
	"github.com/hassannmoussaa/pill.go/helpers"
	"github.com/valyala/fasthttp"
)

type UserAuthCtrl struct {
	*APICtrl
}

func loginUser(email string, password string) (int, string, *models.User, error) {
	requestedUser := models.GetUserByEmail(email)
	if requestedUser != nil {
		jwtAuth := auth.GetJWTAuth()
		if auth.CompareHashAndPassword(password, requestedUser.Password()) {
			auth.ResetUserAttempts(int(requestedUser.ID()))
			token, err := jwtAuth.GenerateToken(int(requestedUser.ID()), "user")
			if err != nil {
				return http.StatusInternalServerError, "", nil, nil
			}
			return http.StatusOK, token, requestedUser, nil
		}
		return http.StatusUnauthorized, "", nil, errors.New("authentication_credentials_invalid")
	}
	return http.StatusUnauthorized, "", nil, errors.New("authentication_credentials_invalid")
}

func (this *UserAuthCtrl) Login(requestCtx *fasthttp.RequestCtx) {
	var err error
	statusCode := 200
	var accessToken string
	var user *models.User
	email := helpers.BytesToString(requestCtx.PostArgs().Peek("email"))
	password := helpers.BytesToString(requestCtx.PostArgs().Peek("password"))

	if email == "" || password == "" {
		err = errors.New("authentication_credentials_incomplete")
		statusCode = 422
	} else {
		statusCode, accessToken, user, err = loginUser(email, password)
	}

	if statusCode != 200 {
		if statusCode == 422 {
			this.ValidationError(requestCtx, nil, err.Error())
		} else {
			this.Fail(requestCtx, nil, err.Error(), 401)
		}
	} else {
		rememberMe := requestCtx.PostArgs().GetBool("remember_me")
		auth.SetAccessTokenFastHttpCookie(requestCtx, accessToken, rememberMe)
		fields, excluded := this.SetFields(requestCtx)
		data := user.ToMap("", excluded, fields...)
		this.Success(requestCtx, data, "login_successfully")
	}
}

func (this *UserAuthCtrl) Logout(requestCtx *fasthttp.RequestCtx) {
	auth.RemoveAccessTokenFastHttpCookie(requestCtx)
	this.Success(requestCtx, nil, "logout_successfully")
}

func (this *UserAuthCtrl) Refresh(requestCtx *fasthttp.RequestCtx) {
	ctx := appCtx.Get(requestCtx)
	jwtAuth := auth.GetJWTAuth()
	token, err := jwtAuth.GenerateToken(int(ctx.LoggedUser.ID()), "user")
	if err != nil {
		this.Fail(requestCtx, nil, err.Error(), 401)
	} else {
		this.Success(requestCtx, struct {
			AccessToken string `json:"access_token,omitempty"`
		}{AccessToken: token}, "")
	}
}
