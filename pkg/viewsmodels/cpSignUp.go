package viewsmodels

import (
	"github.com/hassannmoussaa/bookery/pkg/appCtx"
	"github.com/hassannmoussaa/bookery/pkg/textualContent"
)

type CPSignUp struct {
	*Page
}

func GetCPSignUp(ctx *appCtx.Context) *CPSignUp {
	cpSignUp := &CPSignUp{}
	cpSignUp.Page = NewPage("cp-signup", ctx)
	cpSignUp.Page.SetMetas(textualContent.OfTitle("signup"), textualContent.OfTitle("signup"))
	return cpSignUp
}
