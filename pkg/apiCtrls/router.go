package apiCtrls

import (
	"github.com/hassannmoussaa/pill.go/fastmux"
	"github.com/hassannmoussaa/bookery/pkg/middlewares"
	"github.com/valyala/fasthttp"
)

func Register(basePath string) *fastmux.Mux {
	router := fastmux.New(basePath)

	router.Get("/health").ThenFunc(func(requestCtx *fasthttp.RequestCtx) {
		requestCtx.Write([]byte("I am alive!"))
	})
	
	router.Use(middlewares.SetResponseContentTypeAsJSON, middlewares.OnlyHttps, middlewares.AllowCORSRequest)

	//Allow cors request
	router.Options("/*").ThenFunc(func(requestCtx *fasthttp.RequestCtx) {
		requestCtx.Write([]byte{})
	})

	router.Use(middlewares.GetLoggedAdmin, middlewares.ParseQueryParams)

	//api ctrls
	adminCtrl := &AdminCtrl{}
	adminAuthCtrl := &AdminAuthCtrl{}
	uploadCtrl := &UploadCtrl{}

	router.Post("/admins/auth").ThenFunc(adminAuthCtrl.Login)
	router.Delete("/admins/auth").ThenFunc(adminAuthCtrl.Logout)

	router.Use(middlewares.IsAuthenticatedAdmin)

	router.Post("/admins").ThenFunc(adminCtrl.Add)
	router.Get("/admins/me").ThenFunc(adminCtrl.GetMe)
	router.Post("/admins/auth/refresh").ThenFunc(adminAuthCtrl.Refresh)

	router.Post("/upload").ThenFunc(uploadCtrl.Upload)

	return router
}
