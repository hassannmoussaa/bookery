package hooks

import (
	"github.com/hassannmoussaa/bookery/pkg/models"
	"github.com/nehmeroumani/pill.go/hooks"
	"github.com/nehmeroumani/pill.go/mailer"
)

func bookHooks() {
	hooks.Main.AddAction("book_is_recived", func(data ...interface{}) {
		if data != nil {
			if len(data) > 0 {
				if book, ok := data[0].(*models.Book); ok {
					data := struct {
						UserName string
						BookName string
					}{UserName: book.User().FullName()}
					mailer.Send([]string{book.User().Email()}, "Your Book was recived!", "bookrecived.html", data)
				}
			}
		}
	})
	hooks.Main.AddAction("book_is_verified", func(data ...interface{}) {
		if data != nil {
			if len(data) > 0 {
				if book, ok := data[0].(*models.Book); ok {
					data := struct {
						UserName string
						BookName string
					}{UserName: book.User().FullName()}
					mailer.Send([]string{book.User().Email()}, "Your Book was verified!", "bookverified.html", data)
				}
			}
		}
	})
}
