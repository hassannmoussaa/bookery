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
	router.Get("/").ThenFunc(cpCtrl.Index)

	router.Get("/cp/login").ThenFunc(cpCtrl.Login)
	router.Get("/signup").Use(middlewares.IsAuthenticatedUserBot).ThenFunc(cpCtrl.SignUp)
	router.Get("/cp").Use(middlewares.IsAuthenticatedAdmin).ThenFunc(cpCtrl.Dashboard)
	router.Get("/cp/users/:search").Use(middlewares.IsAuthenticatedAdmin).ThenFunc(cpCtrl.Users)
	router.Get("/cp/users").Use(middlewares.IsAuthenticatedAdmin).ThenFunc(cpCtrl.Users)
	router.Get("/cp/orders").Use(middlewares.IsAuthenticatedAdmin).ThenFunc(cpCtrl.Orders)
	router.Get("/cp/categories").Use(middlewares.IsAuthenticatedAdmin).ThenFunc(cpCtrl.Categories)
	router.Get("/cp/category/add").Use(middlewares.IsAuthenticatedAdmin).ThenFunc(cpCtrl.AddCategory)
	router.Get("/cp/admins").Use(middlewares.IsAuthenticatedAdmin).ThenFunc(cpCtrl.Admins)
	router.Get("/cp/admin/add").Use(middlewares.IsAuthenticatedAdmin).ThenFunc(cpCtrl.AddAdmin)
	router.Get("/cp/books/:search").Use(middlewares.IsAuthenticatedAdmin).ThenFunc(cpCtrl.Books)
	router.Get("/cp/books").Use(middlewares.IsAuthenticatedAdmin).ThenFunc(cpCtrl.Books)
	router.Get("/cp/cardorders").Use(middlewares.IsAuthenticatedAdmin).ThenFunc(cpCtrl.CardOrders)

	return router
}
