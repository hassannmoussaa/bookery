package hooks

import (
	"github.com/hassannmoussaa/bookery/pkg/models"
	"github.com/nehmeroumani/pill.go/hooks"
	"github.com/nehmeroumani/pill.go/mailer"
)

func userHooks() {
	hooks.Main.AddAction("user_account_is_blocked", func(data ...interface{}) {
		if data != nil {
			if len(data) > 0 {
				if user, ok := data[0].(*models.User); ok {
					data := struct {
						UserName string
					}{UserName: user.FullName()}
					mailer.Send([]string{user.Email()}, "Your account was blocked!", "user-account-is-locked.html", data)
				}
			}
		}
	})
}
