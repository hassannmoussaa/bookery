package models

import (
	"errors"
	"math"
	"strconv"
	"strings"

	"github.com/hassannmoussaa/bookery/pkg/db"
	"github.com/hassannmoussaa/pill.go/clean"
)

type Transaction struct {
	id              int32
	user            *User
	user_old_credit int
	user_new_credit int
	date            string
}

func (this *Transaction) ToMap(prefix string, excluded bool, fields ...string) map[string]interface{} {
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
	if this.User().FullName() != "" && check(prefix+"user_name", fields...) {
		result["user_name"] = this.User().FullName()
	}
	if this.UserOldCredit() != 0 && check(prefix+"user_old_credit", fields...) {
		result["user_old_credit"] = this.UserOldCredit()
	}
	if this.UserNewCredit() != 0 && check(prefix+"user_new_credit", fields...) {
		result["user_new_credit"] = this.UserNewCredit()
	}
	if this.Date() != "" && check(prefix+"date", fields...) {
		result["date"] = this.Date()
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

func (this *Transaction) ID() int32 {
	return this.id
}
func (this *Transaction) User() *User {
	return this.user
}
func (this *Transaction) UserOldCredit() int {
	return this.user_old_credit
}
func (this *Transaction) UserNewCredit() int {
	return this.user_new_credit
}
func (this *Transaction) Date() string {
	return this.date
}
func (this *Transaction) SetID(value int32) {
	this.id = value
}
func (this *Transaction) SetUser(value *User) {
	this.user = value
}
func (this *Transaction) SetUserOldCredit(value int) {
	this.user_old_credit = value
}
func (this *Transaction) SetUserNewCredit(value int) {
	this.user_new_credit = value
}
func (this *Transaction) SetDate(value string) {
	this.date = value
}

func AddTransaction(transaction *Transaction) *Transaction {
	if transaction != nil {
		sql := "INSERT INTO " + db.TransactionTable + " (user_id, user_old_credit, user_new_credit, date) VALUES ($1, $2, $3, $4) RETURNING id;"
		row := connection.QueryRow(sql, transaction.user.id, transaction.user_old_credit, transaction.user_new_credit, transaction.date)
		err := row.Scan(&transaction.id)
		if err != nil {
			clean.Error(err)
			return nil
		}
		return transaction
	}
	return nil
}
func GetTransactionById(id int32) *Transaction {
	if id != 0 {
		sql := "SELECT user_id, coalesce(user_old_credit, 0), coalesce(user_new_credit, 0), coalesce(date, '') FROM " + db.TransactionTable + " WHERE id=$1"
		row := connection.QueryRow(sql, id)
		transaction := &Transaction{}
		transaction.id = id
		user := &User{}
		err := row.Scan(&user.id, &transaction.user_old_credit, &transaction.user_new_credit, &transaction.date)
		transaction.user = GetUserById(user.id)
		if err != nil {
			clean.Error(err)
			return nil
		}
		return transaction
	}
	return nil
}
func GetTransactionByUserId(userid int32) *Transaction {
	if userid != 0 {
		sql := "SELECT id, coalesce(user_old_credit, 0), coalesce(user_new_credit, 0), coalesce(date, '') FROM " + db.TransactionTable + " WHERE user_id=$1"
		row := connection.QueryRow(sql, userid)
		transaction := &Transaction{}
		user := &User{}
		user.id = userid
		err := row.Scan(&transaction.id, &transaction.user_old_credit, &transaction.user_new_credit, &transaction.date)
		transaction.user = GetUserById(user.id)
		if err != nil {
			clean.Error(err)
			return nil
		}
		return transaction
	}
	return nil
}
func DeleteTransaction(id int32) bool {
	if id != 0 {
		sql := "DELETE FROM " + db.TransactionTable + " WHERE id=$1"
		_, err := connection.Exec(sql, id)
		if err != nil {
			clean.Error(err)
			return false
		}
		return true
	}
	return false
}
func GetTransactions(page int32, count int32, sinceID int32) ([]*Transaction, bool, int32) {
	sql := "SELECT id,  user_id, coalesce(user_old_credit, 0), coalesce(user_new_credit, 0), coalesce(date, '')FROM " + db.TransactionTable
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
	transactions := []*Transaction{}
	for rows.Next() {
		transaction := &Transaction{}
		user := &User{}
		err = rows.Scan(&transaction.id, &user.id, &transaction.user_old_credit, &transaction.user_new_credit, &transaction.date)
		if err != nil {
			clean.Error(err)
			continue
		}
		transactions = append(transactions, transaction)
	}
	var hasMore bool
	var nextPagesCount int32
	transactionsCount := int32(len(transactions))
	if (sinceID != 0 && count > 0) || (page <= 0 && count > 0) {
		if transactionsCount > count {
			hasMore = true
			transactions = transactions[:count]
		}
	} else if page > 0 {
		if transactionsCount > count && count > 0 {
			hasMore = true
			nextPagesCount = int32(math.Ceil(float64(transactionsCount)/float64(count))) - 1
			transactions = transactions[:count]
		}
	}
	return transactions, hasMore, nextPagesCount
}
func GetTransactionsCount() int {
	sql := "SELECT COUNT(*) FROM " + db.TransactionTable
	row := connection.QueryRow(sql)
	var count int64
	err := row.Scan(&count)
	if err != nil {
		clean.Error(err)
		return 0
	}
	return int(count)
}
func PrepareAndValidateTransaction(transaction *Transaction) error {
	if transaction != nil {
		if transaction.date == "" {
			return errors.New("date_is_required")
		}

		return nil
	}
	return errors.New("invalid_data")
}
