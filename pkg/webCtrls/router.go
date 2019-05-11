package webCtrls

import (
	"github.com/hassannmoussaa/bookery/pkg/middlewares"
	"github.com/hassannmoussaa/pill.go/fastmux"
	"github.com/hassannmoussaa/pill.go/fastmux/util"
	"github.com/valyala/fasthttp"
)

func Register() *fastmux.Mux {
	//New Router
	router := fastmux.New()

	router.Get("/health").ThenFunc(func(requestCtx *fasthttp.RequestCtx) {
		requestCtx.Write([]byte("I am alive!"))
	})

	router.Use(middlewares.OnlyHttps)

	router.Get("/static/scripts/app/cp-pages/*file").Use(middlewares.GetLoggedAdmin, middlewares.IsAuthenticatedAdmin).ThenFunc(util.StaticFilesServe)
	router.Get("/static").ThenFunc(util.StaticFilesServe)
	router.Get("/uploads").ThenFunc(util.UploadsServe)
	router.Get("/static/*file").ThenFunc(util.StaticFilesServe)
	router.Get("/uploads/*file").ThenFunc(util.UploadsServe)

	router.Use(middlewares.GetLoggedAdmin, middlewares.CSRFToken, middlewares.ParseQueryParams)
	//Custom "not found" page
	router.NotFoundHandler(func(requestCtx *fasthttp.RequestCtx) {
		PageNotFound(requestCtx)
	})

	notFoundCtrl := &NotFoundCtrl{}
	cpCtrl := &CPCtrl{}

	router.Get("/404").ThenFunc(notFoundCtrl.Get)
	router.Get("/cp/login").ThenFunc(cpCtrl.Login)
	router.Get("/cp").Use(middlewares.IsAuthenticatedAdmin).ThenFunc(cpCtrl.Dashboard)
	router.Get("/cp/users").Use(middlewares.IsAuthenticatedAdmin).ThenFunc(cpCtrl.Users)
	return router
}
