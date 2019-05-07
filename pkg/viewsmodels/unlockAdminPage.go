package viewsmodels

import (
	"github.com/hassannmoussaa/bookery/pkg/appCtx"
	"github.com/hassannmoussaa/bookery/pkg/textualContent"
)

type UnlockAdminPage struct {
	*Page
	WasUnlocked bool
}

func GetUnlockAdminPage(ctx *appCtx.Context, success bool) *UnlockAdminPage {
	unlockAdminPage := &UnlockAdminPage{}
	unlockAdminPage.Page = NewPage("unlock-admin-account", ctx)
	unlockAdminPage.WasUnlocked = success
	if success {
		unlockAdminPage.Page.SetMetas(textualContent.OfTitle("account_was_unlocked"), textualContent.OfTitle("account_was_unlocked"))
	} else {
		unlockAdminPage.Page.SetMetas(textualContent.OfTitle("failed_to_unlock_the_account"), textualContent.OfTitle("failed_to_unlock_the_account"))
	}
	return unlockAdminPage
}
