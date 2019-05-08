package models

import (
	"errors"
	"strings"

	"github.com/hassannmoussaa/bookery/pkg/db"
	"github.com/hassannmoussaa/pill.go/clean"
)

type UserBook struct {
	id   int32
	user *User
	book *Book
}

func (this *UserBook) ToMap(prefix string, excluded bool, fields ...string) map[string]interface{} {
	result := map[string]interface{}{}
	if prefix != "" && !strings.HasSuffix(prefix, ".") {
		prefix += "."
	}
	check := Included
	if excluded {
		check = Excluded
	}
	if this.ID() != 0 && check(prefix+"id", fields...) {
		result["id"] = this.ID()
	}
	if this.User().FullName() != "" && check(prefix+"full_name", fields...) {
		result["full_name"] = this.User().FullName()
	}
	if this.Book().BookName() != "" && check(prefix+"book_name", fields...) {
		result["book_name"] = this.Book().BookName()
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

func (this *UserBook) ID() int32 {
	return this.id
}
func (this *UserBook) User() *User {
	return this.user
}
func (this *UserBook) Book() *Book {
	return this.book
}

func (this *UserBook) SetID(value int32) {
	this.id = value
}
func (this *UserBook) SetUser(value *User) {
	this.user = value
}
func (this *UserBook) SetBook(value *Book) {
	this.book = value
}

func AddUserBook(userbook *UserBook) *UserBook {
	if userbook != nil {
		sql := "INSERT INTO " + db.UserBookTable + " (user_id , book_id) VALUES ($1, $2) RETURNING id;"
		row := connection.QueryRow(sql, userbook.user.id, userbook.book.id)
		err := row.Scan(&category.id)
		if err != nil {
			clean.Error(err)
			return nil
		}
		return userbook
	}
	return nil
}
func GetUserBookById(id int32) *UserBook {
	if id != 0 {
		sql := "SELECT user_id , book_id FROM " + db.UserBookTable + " WHERE id=$1"
		row := connection.QueryRow(sql, id)
		userbook := &UserBook{}
		userbook.id = id
		user := &User{}
		book := &Book{}
		err := row.Scan(&user.id, &book.id)
		userbook.user = GetUserById(user.id)
		userbook.book = GetBookById(book.id)
		if err != nil {
			clean.Error(err)
			return nil
		}
		return category
	}
	return nil
}
func GetUserBookByUserId(userid int32) *UserBook {
	if id != 0 {
		sql := "SELECT id , book_id FROM " + db.UserBookTable + " WHERE user_id=$1"
		row := connection.QueryRow(sql, userid)
		userbook := &UserBook{}
		user := &User{}
		user.id = userid
		book := &Book{}
		err := row.Scan(&userbook.id, &book.id)
		userbook.user = GetUserById(user.id)
		userbook.book = GetBookById(book.id)
		if err != nil {
			clean.Error(err)
			return nil
		}
		return category
	}
	return nil
}
func GetUserBookByBookId(bookid int32) *UserBook {
	if id != 0 {
		sql := "SELECT id , user_id FROM " + db.UserBookTable + " WHERE bookid=$1"
		row := connection.QueryRow(sql, userid)
		userbook := &UserBook{}
		user := &User{}
		book := &Book{}
		book.id = bookid
		err := row.Scan(&userbook.id, &user.id)
		userbook.user = GetUserById(user.id)
		userbook.book = GetBookById(book.id)
		if err != nil {
			clean.Error(err)
			return nil
		}
		return category
	}
	return nil
}
func DeleteUserBook(id int32) bool {
	if id != 0 {
		sql := "DELETE FROM " + db.UserBookTable + " WHERE id=$1"
		_, err := connection.Exec(sql, id)
		if err != nil {
			clean.Error(err)
			return false
		}
		return true
	}
	return false
}

func PrepareAndValidateUserBook(category *UserBook) error {
	if category != nil {
		if category.category_name == "" {
			return errors.New("category_name_is_required")
		}

		return nil
	}
	return errors.New("invalid_data")
}
