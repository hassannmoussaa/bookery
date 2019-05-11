package models

import (
	"errors"
	"math"
	"strconv"
	"strings"

	"github.com/hassannmoussaa/bookery/pkg/db"
	"github.com/hassannmoussaa/pill.go/auth"
	"github.com/hassannmoussaa/pill.go/clean"
	"github.com/hassannmoussaa/pill.go/hooks"
	"github.com/hassannmoussaa/pill.go/validate"
)

type User struct {
	id           int32
	email        string
	full_name    string
	password     string
	full_address string
	phone_number string
	user_credit  int
	is_blocked   bool
}

func (this *User) ToMap(prefix string, excluded bool, fields ...string) map[string]interface{} {
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
	if this.Email() != "" && check(prefix+"email", fields...) {
		result["email"] = this.Email()
	}
	if this.FullName() != "" && check(prefix+"full_name", fields...) {
		result["full_name"] = this.FullName()
	}
	if this.FullAddress() != "" && check(prefix+"full_address", fields...) {
		result["full_address"] = this.FullAddress()
	}
	if this.PhoneNumber() != "" && check(prefix+"phone_number", fields...) {
		result["phone_number"] = this.PhoneNumber()
	}
	if this.UserCredit() != 0 && check(prefix+"user_credit", fields...) {
		result["user_credit"] = this.UserCredit()
	}
	if check(prefix+"is_blocked", fields...) {
		result["is_blocked"] = this.IsBlocked()
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

func (this *User) ID() int32 {
	return this.id
}
func (this *User) Email() string {
	return this.email
}
func (this *User) FullName() string {
	return this.full_name
}
func (this *User) Password() string {
	return this.password
}
func (this *User) FullAddress() string {
	return this.full_address
}
func (this *User) PhoneNumber() string {
	return this.phone_number
}
func (this *User) UserCredit() int {
	return this.user_credit
}
func (this *User) IsBlocked() bool {
	return this.is_blocked
}

func (this *User) SetID(value int32) {
	this.id = value
}
func (this *User) SetEmail(value string) {
	this.email = value
}
func (this *User) SetFullName(value string) {
	this.full_name = value
}
func (this *User) SetPassword(value string) {
	this.password = value
}
func (this *User) SetFullAddress(value string) {
	this.full_address = value
}
func (this *User) SetPhoneNumber(value string) {
	this.phone_number = value
}
func (this *User) SetUserCredit(value int) {
	this.user_credit = value
}
func (this *User) SetsBlocked(value bool) {
	this.is_blocked = value
}

func AddUser(user *User) *User {
	if user != nil {
		sql := "INSERT INTO " + db.UserTable + " (email, full_name, password, full_address, phone_number, user_credit, is_blocked) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;"
		row := connection.QueryRow(sql, user.email, user.full_name, user.password, user.full_address, user.phone_number, user.user_credit, user.is_blocked)
		err := row.Scan(&user.id)
		if err != nil {
			clean.Error(err)
			return nil
		}
		return user
	}
	return nil
}
func GetUserById(id int32) *User {
	if id != 0 {
		sql := "SELECT coalesce(email, ''), coalesce(full_name, ''), coalesce(password, ''), coalesce(full_address, ''), coalesce(phone_number, ''), coalesce(user_credit, 0), coalesce(is_blocked, false) FROM " + db.UserTable + " WHERE id=$1"
		row := connection.QueryRow(sql, id)
		user := &User{}
		user.id = id
		err := row.Scan(&user.email, &user.full_name, &user.password, &user.full_address, &user.phone_number, &user.user_credit, &user.is_blocked)
		if err != nil {
			clean.Error(err)
			return nil
		}
		return user
	}
	return nil
}
func GetUserByEmail(email string) *User {
	if email != "" {
		sql := "SELECT id, coalesce(full_name, ''), coalesce(password, ''), coalesce(full_address, ''), coalesce(phone_number, ''), coalesce(user_credit, 0), coalesce(is_blocked, false) FROM " + db.UserTable + " WHERE email=$1"
		row := connection.QueryRow(sql, email)
		user := &User{}
		user.email = email
		err := row.Scan(&user.id, &user.full_name, &user.password, &user.full_address, &user.phone_number, &user.user_credit, &user.is_blocked)
		if err != nil {
			clean.Error(err)
			return nil
		}
		return user
	}
	return nil
}
func IsUserEmailExist(email string, opts ...int32) bool {
	if email != "" {
		sql := "SELECT id FROM " + db.UserTable + " WHERE email=$1"
		var UserId int32
		if opts != nil {
			if len(opts) > 0 {
				UserId = opts[0]
			}
		}
		if UserId != 0 {
			sql += " AND id!=" + strconv.Itoa(int(UserId))
		}
		row := connection.QueryRow(sql, email)
		UserId = 0
		err := row.Scan(&UserId)
		if err != nil {
			clean.Error(err)
			return false
		}
		if UserId != 0 {
			return true
		}
		return false
	}
	return true
}
func DeleteUser(id int32) bool {
	if id != 0 {
		sql := "DELETE FROM " + db.UserTable + " WHERE id=$1"
		_, err := connection.Exec(sql, id)
		if err != nil {
			clean.Error(err)
			return false
		}
		return true
	}
	return false
}

func CheckIFUserBlocked(id int32) bool {
	var isBlocked bool
	if id != 0 {
		sql := "SELECT coalesce(is_blocked, false) FROM " + db.UserTable + " WHERE id=$1"
		row := connection.QueryRow(sql, id)
		err := row.Scan(isBlocked)
		if err != nil {
			clean.Error(err)
			return false
		}
		return isBlocked
	}
	return false
}
func BlockUser(user *User) bool {
	if user != nil {
		user.is_blocked = true
		sql := "UPDATE " + db.UserTable + " SET is_blocked=$1 WHERE id=$2"
		_, err := connection.Exec(sql, user.is_blocked, user.id)
		if err != nil {
			clean.Error(err)
			return false
		}
		hooks.Main.DoAction("user_account_is_blocked", user)
		return true
	}
	return false
}
func UnBlockUser(user *User) bool {
	if user != nil {
		user.is_blocked = false
		sql := "UPDATE " + db.UserTable + " SET is_blocked=$1 WHERE id=$2"
		_, err := connection.Exec(sql, user.is_blocked, user.id)
		if err != nil {
			clean.Error(err)
			return false
		}
		return true
	}
	return false
}
func GetUsers(page int32, count int32, sinceID int32) ([]*User, bool, int32) {
	sql := "SELECT id, coalesce(full_name, '') , coalesce(email, ''), coalesce(password, ''), coalesce(full_address, ''), coalesce(phone_number, ''), coalesce(user_credit, 0), coalesce(is_blocked, false) FROM " + db.UserTable
	values := make([]interface{}, 3)
	j := 0
	if sinceID > 0 {
		sql += " WHERE "
	}
	if sinceID > 0 {
		sql += "id<$" + strconv.Itoa(j+1)
		values[j] = sinceID
		j++
	}
	sql += ` ORDER BY id DESC`
	if sinceID == 0 {
		if page > 0 {
			offset := (page - 1) * count
			sql += " OFFSET $" + strconv.Itoa(j+1)
			values[j] = offset
			j++
		}
	}
	if count > 0 {
		sql += " LIMIT $" + strconv.Itoa(j+1)
		if (sinceID != 0 && count > 0) || (page <= 0 && count > 0) {
			values[j] = count + 1
		} else {
			values[j] = count * 4
		}
		j++
	}
	values = values[:j]
	rows, err := connection.Query(sql, values...)
	defer rows.Close()
	if err != nil {
		clean.Error(err)
		return nil, false, 0
	}
	users := []*User{}
	for rows.Next() {
		user := &User{}
		err = rows.Scan(&user.id, &user.full_name, &user.email, &user.password, &user.full_address, &user.phone_number, &user.user_credit, &user.is_blocked)
		if err != nil {
			clean.Error(err)
			continue
		}
		users = append(users, user)
	}
	var hasMore bool
	var nextPagesCount int32
	usersCount := int32(len(users))
	if (sinceID != 0 && count > 0) || (page <= 0 && count > 0) {
		if usersCount > count {
			hasMore = true
			users = users[:count]
		}
	} else if page > 0 {
		if usersCount > count && count > 0 {
			hasMore = true
			nextPagesCount = int32(math.Ceil(float64(usersCount)/float64(count))) - 1
			users = users[:count]
		}
	}
	return users, hasMore, nextPagesCount
}
func GetSimilariUsers(page int32, count int32, sinceID int32, search string) ([]*User, bool, int32) {
	search = strings.ToLower(search)
	sql := "SELECT id, coalesce(full_name, '') , coalesce(email, ''), coalesce(password, ''), coalesce(full_address, ''), coalesce(phone_number, ''), coalesce(user_credit, 0), coalesce(is_blocked, false)  FROM " + db.UserTable + " WHERE full_name LIKE $1"
	values := make([]interface{}, 3)
	j := 1
	if sinceID > 0 {
		sql += " WHERE "
	}
	if sinceID > 0 {
		sql += "id<$" + strconv.Itoa(j+1)
		values[j] = sinceID
		j++
	}
	sql += ` ORDER BY id DESC`
	if sinceID == 0 {
		if page > 0 {
			offset := (page - 1) * count
			sql += " OFFSET $" + strconv.Itoa(j+1)
			values[j] = offset
			j++
		}
	}
	if count > 0 {
		sql += " LIMIT $" + strconv.Itoa(j+1)
		if (sinceID != 0 && count > 0) || (page <= 0 && count > 0) {
			values[j] = count + 1
		} else {
			values[j] = count * 4
		}
		j++
	}
	if len(search) >= 2 {
		values[0] = search[0:2] + "%"
	} else {
		values[0] = search + "%"
	}
	values = values[:j]
	rows, err := connection.Query(sql, values...)
	defer rows.Close()
	if err != nil {
		clean.Error(err)
		return nil, false, 0
	}
	users := []*User{}
	for rows.Next() {
		user := &User{}
		err = rows.Scan(&user.id, &user.full_name, &user.email, &user.password, &user.full_address, &user.phone_number, &user.user_credit, &user.is_blocked)
		if err != nil {
			clean.Error(err)
			continue
		}
		users = append(users, user)
	}
	var hasMore bool
	var nextPagesCount int32
	usersCount := int32(len(users))
	if (sinceID != 0 && count > 0) || (page <= 0 && count > 0) {
		if usersCount > count {
			hasMore = true
			users = users[:count]
		}
	} else if page > 0 {
		if usersCount > count && count > 0 {
			hasMore = true
			nextPagesCount = int32(math.Ceil(float64(usersCount)/float64(count))) - 1
			users = users[:count]
		}
	}
	return users, hasMore, nextPagesCount
}
func GetUsersCount() int {
	sql := "SELECT COUNT(*) FROM " + db.UserTable
	row := connection.QueryRow(sql)
	var count int64
	err := row.Scan(&count)
	if err != nil {
		clean.Error(err)
		return 0
	}
	return int(count)
}
func PrepareAndValidateUser(user *User) error {
	if user != nil {
		if user.email == "" {
			return errors.New("email_is_required")
		}
		if user.phone_number == "" {
			return errors.New("phone_number_is_required")
		}
		if !validate.Email(user.email) {
			return errors.New("invalid_email")
		}

		if IsUserEmailExist(user.email, user.id) {
			return errors.New("email_exist")
		}

		if user.full_name == "" {
			return errors.New("name_is_required")
		}
		if !validate.StringLength(user.full_name, 3, 30) {
			return errors.New("invalid_name_length")
		}
		if user.password == "" {
			return errors.New("password_is_required")
		}
		if !validate.StringLength(user.password, 6, 16) {
			return errors.New("invalid_password_length")
		}
		user.password = auth.HashPassword(user.password)
		return nil
	}
	return errors.New("invalid_data")
}
