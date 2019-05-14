package middlewares

import (
	"github.com/hassannmoussaa/bookery/pkg/appCtx"
	"github.com/hassannmoussaa/bookery/pkg/models"
	"github.com/hassannmoussaa/bookery/pkg/textualContent"
	"github.com/hassannmoussaa/pill.go/auth"
	"github.com/hassannmoussaa/pill.go/fastmux/util"
	"github.com/hassannmoussaa/pill.go/helpers"
	"github.com/valyala/fasthttp"
)

func GetLoggedAdmin(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	fn := func(requestCtx *fasthttp.RequestCtx) {
		ctx := appCtx.Get(requestCtx)
		//if ctx.LoggedUser == nil {
		token := auth.GetTokenFromFastHttpRequest(requestCtx)
		isAuthenticated, adminId, role := auth.IsAuthenticated(token)
		if isAuthenticated && role == "admin" {
			ctx.LoggedAdmin = models.GetAdminById(int32(adminId))
			if ctx.LoggedAdmin != nil {
				auth.RefreshAccessTokenFastHttpCookie(requestCtx, adminId, role)
			}
		}
		//}
		if next != nil {
			next(requestCtx)
		}
	}
	return fn
}
func GetLoggedUser(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	fn := func(requestCtx *fasthttp.RequestCtx) {
		ctx := appCtx.Get(requestCtx)
		//if ctx.LoggedUser == nil {
		token := auth.GetTokenFromFastHttpRequest(requestCtx)
		isAuthenticated, userId, role := auth.IsAuthenticated(token)
		if isAuthenticated && role == "user" {
			ctx.LoggedUser = models.GetUserById(int32(userId))
			if ctx.LoggedAdmin != nil {
				auth.RefreshAccessTokenFastHttpCookie(requestCtx, userId, role)
			}
		}
		//}
		if next != nil {
			next(requestCtx)
		}
	}
	return fn
}
func IsAuthenticatedAdmin(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	fn := func(requestCtx *fasthttp.RequestCtx) {
		ctx := appCtx.Get(requestCtx)
		if ctx.LoggedAdmin == nil {
			if helpers.BytesToString(requestCtx.Response.Header.Peek("Content-Type")) == "application/json" {
				requestCtx.SetStatusCode(401)
				requestCtx.Write([]byte(`{"status": "fail", "message": "` + textualContent.OfErrorMsg("unauthorized") + `"}`))
			} else {
				requestCtx.Redirect("/cp/login", 307)
			}
			return
		} else {
			if next != nil {
				next(requestCtx)
			}
		}
	}
	return fn
}
func IsAuthenticatedUserBot(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	fn := func(requestCtx *fasthttp.RequestCtx) {
		if requestCtx.Request.Header.Cookie("UserResponsePath") == nil {
			requestCtx.Redirect("/", 307)
			return
		} else {
			if next != nil {
				next(requestCtx)
			}
		}
	}
	return fn
}
func IsAuthenticatedUser(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	fn := func(requestCtx *fasthttp.RequestCtx) {
		ctx := appCtx.Get(requestCtx)
		if ctx.LoggedUser == nil {
			if helpers.BytesToString(requestCtx.Response.Header.Peek("Content-Type")) == "application/json" {
				requestCtx.SetStatusCode(401)
				requestCtx.Write([]byte(`{"status": "fail", "message": "` + textualContent.OfErrorMsg("unauthorized") + `"}`))
			} else {
				requestCtx.Redirect("/cp/login", 307)
			}
			return
		} else {
			if next != nil {
				next(requestCtx)
			}
		}
	}
	return fn
}

// func GetLoggedUser(next fasthttp.RequestHandler) fasthttp.RequestHandler {
// 	f := func(requestCtx *fasthttp.RequestCtx) {
// 		ctx := appCtx.Get(requestCtx)
// 		if ctx.LoggedAdmin == nil {
// 			token := auth.GetTokenFromFastHttpRequest(requestCtx)
// 			isAuthenticated, userId, role := auth.IsAuthenticated(token)
// 			if isAuthenticated && role == "user" {
// 				ctx.LoggedUser = models.GetUserByID(int32(userId))
// 				if ctx.LoggedUser != nil {
// 					auth.RefreshAccessTokenFastHttpCookie(requestCtx, userId, role)
// 				}
// 			}
// 		}
// 		if next != nil {
// 			next(requestCtx)
// 		}
// 	}
// 	return fasthttp.RequestHandler(f)
// }

// func IsAuthenticatedUser(next fasthttp.RequestHandler) fasthttp.RequestHandler {
// 	f := func(requestCtx *fasthttp.RequestCtx) {
// 		ctx := appCtx.Get(requestCtx)
// 		if ctx.LoggedUser == nil {
// 			unauthorizedUserResponse(requestCtx, ctx)
// 		} else {
// 			if next != nil {
// 				next(requestCtx)
// 			}
// 		}
// 	}
// 	return fasthttp.RequestHandler(f)
// }

func noPermissionResponse(requestCtx *fasthttp.RequestCtx, ctx *appCtx.Context) {
	if helpers.BytesToString(requestCtx.Request.Header.Peek("Content-Type")) == "application/json" {
		requestCtx.SetStatusCode(401)
		requestCtx.Write([]byte(`{"status": "fail", "message": "` + textualContent.OfErrorMsg("unauthorized") + `"}`))
	} else {
		requestCtx.Redirect("/404", 307)
	}
}

func unauthorizedUserResponse(requestCtx *fasthttp.RequestCtx, ctx *appCtx.Context) {
	if helpers.BytesToString(requestCtx.Request.Header.Peek("Content-Type")) == "application/json" {
		requestCtx.SetStatusCode(401)
		requestCtx.Write([]byte(`{"status": "fail", "message": "` + textualContent.OfErrorMsg("unauthorized") + `"}`))
	} else {
		requestCtx.Redirect("/u/login", 307)
	}
}

//API SERVER AUTHENTICATION
func APIServerBasicAuth(requestCtx *fasthttp.RequestCtx) bool {
	username, password := util.GetBasicAuthData(requestCtx, "X-Api-Auth")
	if username != "" && password != "" {
		if !ValidateAPIServerAuth(username, password) {
			requestCtx.SetStatusCode(401)
			requestCtx.Write([]byte(`{"status": "fail", "message": "API: Unauthorized"}`))
			return false
		}
	}
	return true
}

func ValidateAPIServerAuth(username, password string) bool {
	if username == apiServerUsername && password == apiServerPassword {
		return true
	}
	return false
}
