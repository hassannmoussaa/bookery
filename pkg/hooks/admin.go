package hooks

import (
	"strconv"

	"github.com/hassannmoussaa/pill.go/hooks"
	"github.com/hassannmoussaa/pill.go/mailer"
	"github.com/hassannmoussaa/bookery/pkg/models"
)

func adminHooks() {
	hooks.Main.AddAction("admin_account_is_locked", func(data ...interface{}) {
		if data != nil {
			if len(data) > 0 {
				if admin, ok := data[0].(*models.Admin); ok {
					data := struct {
						UnlockAccountURL string
						AdminName        string
					}{UnlockAccountURL: webHost + "/unlock-admin-account?hash=" + admin.Hash() + "&admin_id=" + strconv.Itoa(int(admin.ID())), AdminName: admin.Name()}
					mailer.Send([]string{admin.Email()}, "Your account was locked!", "admin-account-is-locked.html", data)
				}
			}
		}
	})
}
