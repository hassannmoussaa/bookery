package models

import (
	"errors"
	"math"
	"strconv"
	"strings"

	"github.com/hassannmoussaa/bookery/pkg/db"
	"github.com/hassannmoussaa/pill.go/auth"
	"github.com/hassannmoussaa/pill.go/clean"
	"github.com/hassannmoussaa/pill.go/sanitize"
	"github.com/hassannmoussaa/pill.go/validate"
)

type Admin struct {
	id       int32
	email    string
	name     string
	password string
}

func (this *Admin) ToMap(prefix string, excluded bool, fields ...string) map[string]interface{} {
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
	if this.Name() != "" && check(prefix+"name", fields...) {
		result["name"] = this.Name()
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

func (this *Admin) ID() int32 {
	return this.id
}
func (this *Admin) Email() string {
	return this.email
}
func (this *Admin) Name() string {
	return this.name
}
func (this *Admin) Password() string {
	return this.password
}

func (this *Admin) SetID(value int32) {
	this.id = value
}
func (this *Admin) SetEmail(value string) {
	this.email = strings.ToLower(strings.TrimSpace(value))
}
func (this *Admin) SetName(value string) {
	this.name = strings.TrimSpace(sanitize.StripTags(value))
}
func (this *Admin) SetPassword(value string) {
	this.password = value
}

func AddAdmin(admin *Admin) *Admin {
	if admin != nil {
		sql := "INSERT INTO " + db.AdminTable + " (email, name, password) VALUES ($1, $2, $3) RETURNING id;"
		row := connection.QueryRow(sql, admin.email, admin.name, admin.password)
		err := row.Scan(&admin.id)
		if err != nil {
			clean.Error(err)
			return nil
		}
		return admin
	}
	return nil
}
func GetAdminById(id int32) *Admin {
	if id != 0 {
		sql := "SELECT coalesce(email, ''), coalesce(name, ''), coalesce(password, '') FROM " + db.AdminTable + " WHERE id=$1"
		row := connection.QueryRow(sql, id)
		admin := &Admin{}
		admin.id = id
		err := row.Scan(&admin.email, &admin.name, &admin.password)
		if err != nil {
			clean.Error(err)
			return nil
		}
		return admin
	}
	return nil
}
func GetAdminByEmail(email string) *Admin {
	if email != "" {
		sql := "SELECT id, coalesce(name, ''), coalesce(password, '') FROM " + db.AdminTable + " WHERE email=$1"
		row := connection.QueryRow(sql, email)
		admin := &Admin{}
		admin.email = email
		err := row.Scan(&admin.id, &admin.name, &admin.password)
		if err != nil {
			clean.Error(err)
			return nil
		}
		return admin
	}
	return nil
}
func IsAdminEmailExist(email string, opts ...int32) bool {
	if email != "" {
		sql := "SELECT id FROM " + db.AdminTable + " WHERE email=$1"
		var adminId int32
		if opts != nil {
			if len(opts) > 0 {
				adminId = opts[0]
			}
		}
		if adminId != 0 {
			sql += " AND id!=" + strconv.Itoa(int(adminId))
		}
		row := connection.QueryRow(sql, email)
		adminId = 0
		err := row.Scan(&adminId)
		if err != nil {
			clean.Error(err)
			return false
		}
		if adminId != 0 {
			return true
		}
		return false
	}
	return true
}

func DeleteAdmin(id int32) bool {
	if id != 0 {
		sql := "DELETE FROM " + db.AdminTable + " WHERE id=$1"
		_, err := connection.Exec(sql, id)
		if err != nil {
			clean.Error(err)
			return false
		}
		return true
	}
	return false
}
func GetAdmins(page int32, count int32, sinceID int32) ([]*Admin, bool, int32) {
	sql := "SELECT id, coalesce(name, ''), coalesce(email, ''), coalesce(password, '') FROM " + db.AdminTable
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
	admins := []*Admin{}
	for rows.Next() {
		admin := &Admin{}
		err = rows.Scan(&admin.id, &admin.name, &admin.email, &admin.password)
		if err != nil {
			clean.Error(err)
			continue
		}
		admins = append(admins, admin)
	}
	var hasMore bool
	var nextPagesCount int32
	adminsCount := int32(len(admins))
	if (sinceID != 0 && count > 0) || (page <= 0 && count > 0) {
		if adminsCount > count {
			hasMore = true
			admins = admins[:count]
		}
	} else if page > 0 {
		if adminsCount > count && count > 0 {
			hasMore = true
			nextPagesCount = int32(math.Ceil(float64(adminsCount)/float64(count))) - 1
			admins = admins[:count]
		}
	}
	return admins, hasMore, nextPagesCount
}
func GetAdminsCount() int {
	sql := "SELECT COUNT(*) FROM " + db.AdminTable
	row := connection.QueryRow(sql)
	var count int64
	err := row.Scan(&count)
	if err != nil {
		clean.Error(err)
		return 0
	}
	return int(count)
}

func PrepareAndValidateAdmin(admin *Admin) error {
	if admin != nil {
		if admin.email == "" {
			return errors.New("email_is_required")
		}
		if !validate.Email(admin.email) {
			return errors.New("invalid_email")
		}

		if IsAdminEmailExist(admin.email, admin.id) {
			return errors.New("email_exist")
		}

		if admin.name == "" {
			return errors.New("name_is_required")
		}
		if !validate.StringLength(admin.name, 3, 30) {
			return errors.New("invalid_name_length")
		}
		if admin.password == "" {
			return errors.New("password_is_required")
		}
		if !validate.StringLength(admin.password, 6, 16) {
			return errors.New("invalid_password_length")
		}
		admin.password = auth.HashPassword(admin.password)
		return nil
	}
	return errors.New("invalid_data")
}
