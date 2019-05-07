package models

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/hassannmoussaa/pill.go/auth"
	"github.com/hassannmoussaa/pill.go/clean"
	"github.com/hassannmoussaa/pill.go/hooks"
	"github.com/hassannmoussaa/pill.go/sanitize"
	"github.com/hassannmoussaa/pill.go/validate"
	"github.com/hassannmoussaa/bookery/pkg/db"
)

type Admin struct {
	id          int32
	email       string
	name        string
	lockedOn    *time.Time
	password    string
	hash        string
	accessToken string
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
func (this *Admin) Hash() string {
	return this.hash
}
func (this *Admin) AccessToken() string {
	return this.accessToken
}
func (this *Admin) IsLocked() bool {
	if this.lockedOn != nil {
		now := time.Now().UTC()
		if now.Sub(*this.lockedOn).Minutes() < 5 {
			return true
		}
	}
	return false
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
func (this *Admin) SetHash(value string) {
	this.hash = value
}
func (this *Admin) SetAccessToken(value string) {
	this.accessToken = value
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
		sql := "SELECT coalesce(email, ''), coalesce(name, ''), coalesce(password, ''), coalesce(hash, ''), coalesce(locked_on, NULL) FROM " + db.AdminTable + " WHERE id=$1"
		row := connection.QueryRow(sql, id)
		admin := &Admin{}
		admin.id = id
		err := row.Scan(&admin.email, &admin.name, &admin.password, &admin.hash, &admin.lockedOn)
		if err != nil {
			clean.Error(err)
			return nil
		}
		return admin
	}
	return nil
}

func GenerateHashForAdmin(admin *Admin) string {
	if admin != nil {
		data := []byte(admin.Email() + admin.Password() + strconv.Itoa(int(admin.ID())) + time.Now().String())
		hasher := md5.New()
		hasher.Write([]byte(data))
		hash := hex.EncodeToString(hasher.Sum(nil))
		return hash
	}
	return ""
}

func LockAdminAccount(admin *Admin) bool {
	if admin != nil {
		admin.hash = GenerateHashForAdmin(admin)
		now := time.Now().UTC()
		admin.lockedOn = &now
		sql := "UPDATE " + db.AdminTable + " SET locked_on=$1, hash=$2 WHERE id=$3"
		_, err := connection.Exec(sql, admin.lockedOn, admin.hash, admin.id)
		if err != nil {
			clean.Error(err)
			return false
		}
		hooks.Main.DoAction("admin_account_is_locked", admin)
		return true
	}
	return false
}

func UnlockAdminAccount(adminID int, hash string) bool {
	if adminID != 0 && hash != "" {
		sql := `UPDATE ` + db.AdminTable + ` SET locked_on=null, hash='' WHERE id=$1 AND hash=$2`
		_, err := connection.Exec(sql, adminID, hash)
		if err != nil {
			clean.Error(err)
			return false
		}
		return true
	}
	return false
}

func GetAdminByEmail(email string) *Admin {
	if email != "" {
		sql := "SELECT id, coalesce(name, ''), coalesce(password, ''), coalesce(hash, ''), coalesce(locked_on, null) FROM " + db.AdminTable + " WHERE email=$1"
		row := connection.QueryRow(sql, email)
		admin := &Admin{}
		admin.email = email
		err := row.Scan(&admin.id, &admin.name, &admin.password, &admin.hash, &admin.lockedOn)
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
