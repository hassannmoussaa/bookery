package models

import (
	"errors"
	"math"
	"strconv"
	"strings"

	"github.com/hassannmoussaa/bookery/pkg/db"
	"github.com/hassannmoussaa/pill.go/clean"
)

type CardOrder struct {
	id          int32
	user        *User
	card_number int
	credit      int
	transaction *Transaction
	date        string
}

func (this *CardOrder) ToMap(prefix string, excluded bool, fields ...string) map[string]interface{} {
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
	if this.User().FullName() != "" && check(prefix+"username", fields...) {
		result["username"] = this.User().FullName()
	}
	if this.Transaction().UserOldCredit() != 0 && check(prefix+"user_old_credit", fields...) {
		result["user_old_credit"] = this.Transaction().UserOldCredit()
	}
	if this.Transaction().UserNewCredit() != 0 && check(prefix+"user_new_credit", fields...) {
		result["user_new_credit"] = this.Transaction().UserNewCredit()
	}
	if this.Date() != "" && check(prefix+"date", fields...) {
		result["date"] = this.Date()
	}
	if this.CardNumber() != 0 && check(prefix+"card_number", fields...) {
		result["card_number"] = this.CardNumber()
	}
	if this.Credit() != 0 && check(prefix+"credit", fields...) {
		result["credit"] = this.Credit()
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

func (this *CardOrder) ID() int32 {
	return this.id
}
func (this *CardOrder) User() *User {
	return this.user
}
func (this *CardOrder) CardNumber() int {
	return this.card_number
}
func (this *CardOrder) Transaction() *Transaction {
	return this.transaction
}
func (this *CardOrder) Credit() int {
	return this.credit
}
func (this *CardOrder) Date() string {
	return this.date
}
func (this *CardOrder) SetID(value int32) {
	this.id = value
}
func (this *CardOrder) SetUser(value *User) {
	this.user = value
}
func (this *CardOrder) SetCardNumber(value int) {
	this.card_number = value
}
func (this *CardOrder) SetTransaction(value *Transaction) {
	this.transaction = value
}
func (this *CardOrder) SetCredit(value int) {
	this.credit = value
}
func (this *CardOrder) SetDate(value string) {
	this.date = value
}

func AddCardOrder(cardorder *CardOrder) *CardOrder {
	if cardorder != nil {
		sql := "INSERT INTO " + db.CardOrderTable + " (user_id ,card_number ,credit ,transaction_id ,date) VALUES ($1, $2, $3, $4, $5) RETURNING id;"
		row := connection.QueryRow(sql, cardorder.user.id, cardorder.card_number, cardorder.credit, cardorder.transaction.id, cardorder.date)
		err := row.Scan(&cardorder.id)
		if err != nil {
			clean.Error(err)
			return nil
		}
		return cardorder
	}
	return nil
}
func GetCardOrderById(id int32) *CardOrder {
	if id != 0 {
		sql := "SELECT user_id , card_number, credit, transaction_id, date FROM " + db.CardOrderTable + " WHERE id=$1"
		row := connection.QueryRow(sql, id)
		cardorder := &CardOrder{}
		cardorder.id = id
		user := &User{}
		transaction := &Transaction{}
		err := row.Scan(&user.id, &cardorder.card_number, &cardorder.credit, &transaction.id, &cardorder.date)
		cardorder.user = GetUserById(user.id)
		cardorder.transaction = GetTransactionById(transaction.id)
		if err != nil {
			clean.Error(err)
			return nil
		}
		return cardorder
	}
	return nil
}

func DeleteCardOrder(id int32) bool {
	if id != 0 {
		sql := "DELETE FROM " + db.CardOrderTable + " WHERE id=$1"
		_, err := connection.Exec(sql, id)
		if err != nil {
			clean.Error(err)
			return false
		}
		return true
	}
	return false
}
func DeleteCardOrderByUserId(userid int32) bool {
	if userid != 0 {
		sql := "DELETE FROM " + db.CardOrderTable + " WHERE user_id=$1"
		_, err := connection.Exec(sql, userid)
		if err != nil {
			clean.Error(err)
			return false
		}
		return true
	}
	return false
}
func GetCardOrders(page int32, count int32, sinceID int32) ([]*CardOrder, bool, int32) {
	sql := "SELECT id, user_id , card_number, credit, transaction_id, date FROM " + db.CardOrderTable
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
	cardorders := []*CardOrder{}
	for rows.Next() {
		cardorder := &CardOrder{}
		user := &User{}
		transaction := &Transaction{}
		err = rows.Scan(&cardorder.id, &user.id, &cardorder.card_number, &cardorder.credit, &transaction.id, &cardorder.date)
		cardorder.user = GetUserById(user.id)
		cardorder.transaction = GetTransactionById(transaction.id)
		if err != nil {
			clean.Error(err)
			continue
		}
		cardorders = append(cardorders, cardorder)
	}
	var hasMore bool
	var nextPagesCount int32
	cardordersCount := int32(len(cardorders))
	if (sinceID != 0 && count > 0) || (page <= 0 && count > 0) {
		if cardordersCount > count {
			hasMore = true
			cardorders = cardorders[:count]
		}
	} else if page > 0 {
		if cardordersCount > count && count > 0 {
			hasMore = true
			nextPagesCount = int32(math.Ceil(float64(cardordersCount)/float64(count))) - 1
			cardorders = cardorders[:count]
		}
	}
	return cardorders, hasMore, nextPagesCount
}

func PrepareAndValidateCardOrder(cardorder *CardOrder) error {
	if cardorder != nil {
		if cardorder.date == "" {
			return errors.New("date_is_required")
		}

		return nil
	}
	return errors.New("invalid_data")
}
