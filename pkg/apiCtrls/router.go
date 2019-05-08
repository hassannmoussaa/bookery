package apiCtrls

import (
	"github.com/hassannmoussaa/bookery/pkg/middlewares"
	"github.com/hassannmoussaa/pill.go/fastmux"
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

	router.Use(middlewares.GetLoggedAdmin, middlewares.GetLoggedUser, middlewares.ParseQueryParams)

	//api ctrls
	adminCtrl := &AdminCtrl{}
	adminAuthCtrl := &AdminAuthCtrl{}
	userCtrl := &UserCtrl{}
	bookCtrl := &BookCtrl{}
	orderCtrl := &OrderCtrl{}
	cardorderCtrl := &CardOrderCtrl{}
	categoryCtrl := &CategoryCtrl{}
	userAuthCtrl := &UserAuthCtrl{}
	uploadCtrl := &UploadCtrl{}
	//AdminAuth
	router.Post("/admins/auth").ThenFunc(adminAuthCtrl.Login)
	router.Delete("/admins/auth").ThenFunc(adminAuthCtrl.Logout)
	//UserAuth
	router.Post("/users/auth").ThenFunc(userAuthCtrl.Login)
	router.Delete("/users/auth").ThenFunc(userAuthCtrl.Logout)
	//AdminCtrl
	router.Post("/admins").Use(middlewares.IsAuthenticatedAdmin).ThenFunc(adminCtrl.Add)
	router.Delete("/admin/:admin_id").Use(middlewares.IsAuthenticatedAdmin).ThenFunc(adminCtrl.Delete)
	router.Get("/admins/me").Use(middlewares.IsAuthenticatedAdmin).ThenFunc(adminCtrl.GetMe)
	router.Post("/admins/auth/refresh").Use(middlewares.IsAuthenticatedAdmin).ThenFunc(adminAuthCtrl.Refresh)
	//UserCtrl
	router.Post("/users").ThenFunc(userCtrl.Add)
	router.Delete("/users/:user_id").Use(middlewares.IsAuthenticatedAdmin).ThenFunc(userCtrl.Delete)
	//BookCtrl
	router.Post("/books").Use(middlewares.IsAuthenticatedUser).ThenFunc(bookCtrl.Add)
	router.Delete("/admin/books/:book_id").Use(middlewares.IsAuthenticatedAdmin).ThenFunc(bookCtrl.Delete)
	router.Delete("/user/books/:book_id").Use(middlewares.IsAuthenticatedUser).ThenFunc(bookCtrl.Delete)
	//OrderCtrl
	router.Post("/orders").Use(middlewares.IsAuthenticatedUser).ThenFunc(orderCtrl.Add)
	router.Delete("/admin/orders/:order_id").Use(middlewares.IsAuthenticatedAdmin).ThenFunc(orderCtrl.Delete)
	router.Delete("/user/orders/:order_id").Use(middlewares.IsAuthenticatedUser).ThenFunc(orderCtrl.Delete)
	router.Post("/upload").ThenFunc(uploadCtrl.Upload)
	//CategoryCtrl
	router.Post("/categories").Use(middlewares.IsAuthenticatedAdmin).ThenFunc(categoryCtrl.Add)
	router.Delete("/admin/categories/:category_id").Use(middlewares.IsAuthenticatedAdmin).ThenFunc(categoryCtrl.Delete)
	//CardOrderCtrl
	router.Post("/cardorders").Use(middlewares.IsAuthenticatedUser).ThenFunc(cardorderCtrl.Add)
	router.Delete("/admin/cardorders/:card_order_id").Use(middlewares.IsAuthenticatedAdmin).ThenFunc(cardorderCtrl.Delete)
	router.Delete("/user/cardorders/:card_order_id").Use(middlewares.IsAuthenticatedUser).ThenFunc(cardorderCtrl.Delete)
	return router
}
