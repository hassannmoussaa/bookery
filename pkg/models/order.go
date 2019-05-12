package models

import (
	"errors"
	"math"
	"strconv"
	"strings"

	"github.com/hassannmoussaa/bookery/pkg/db"
	"github.com/hassannmoussaa/pill.go/clean"
)

type Order struct {
	id              int32
	user            *User
	book            *Book
	transaction     *Transaction
	date            string
	order_status    string
	delivery_method string
}

func (this *Order) ToMap(prefix string, excluded bool, fields ...string) map[string]interface{} {
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
	if this.Book().BookName() != "" && check(prefix+"book_name", fields...) {
		result["book_name"] = this.Book().BookName()
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
	if this.OrderStatus() != "" && check(prefix+"order_status", fields...) {
		result["order_status"] = this.OrderStatus()
	}
	if this.DeliveryMethod() != "" && check(prefix+"delivery_method", fields...) {
		result["delivery_method"] = this.DeliveryMethod()
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

func (this *Order) ID() int32 {
	return this.id
}
func (this *Order) User() *User {
	return this.user
}
func (this *Order) Book() *Book {
	return this.book
}
func (this *Order) Transaction() *Transaction {
	return this.transaction
}
func (this *Order) Date() string {
	return this.date
}
func (this *Order) OrderStatus() string {
	return this.order_status
}
func (this *Order) DeliveryMethod() string {
	return this.delivery_method
}
func (this *Order) SetID(value int32) {
	this.id = value
}
func (this *Order) SetUser(value *User) {
	this.user = value
}
func (this *Order) SetBook(value *Book) {
	this.book = value
}
func (this *Order) SetTransaction(value *Transaction) {
	this.transaction = value
}
func (this *Order) SetDate(value string) {
	this.date = value
}
func (this *Order) SetOrderStatus(value string) {
	this.order_status = value
}
func (this *Order) SetDeliveryMethod(value string) {
	this.delivery_method = value
}

func AddOrder(order *Order) *Order {
	if order != nil {
		sql := "INSERT INTO " + db.OrderTable + " (user_id, book_id, transaction_id, date , order_status , delivery_method) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;"
		row := connection.QueryRow(sql, order.user.id, order.book.id, order.transaction.id, order.date, order.order_status, order.delivery_method)
		err := row.Scan(&order.id)
		if err != nil {
			clean.Error(err)
			return nil
		}
		return order
	}
	return nil
}
func GetOrderById(id int32) *Order {
	if id != 0 {
		sql := "SELECT user_id, book_id, transaction_id, coalesce(date, ''), coalesce(order_status, ''), coalesce(delivery_method, '')  FROM " + db.OrderTable + " WHERE id=$1"
		row := connection.QueryRow(sql, id)
		order := &Order{}
		order.id = id
		user := &User{}
		book := &Book{}
		transaction := &Transaction{}
		err := row.Scan(&user.id, &book.id, &transaction.id, &order.date, &order.order_status, &order.delivery_method)
		order.user = GetUserById(user.id)
		order.book = GetBookById(book.id)
		order.transaction = GetTransactionById(transaction.id)
		if err != nil {
			clean.Error(err)
			return nil
		}
		return order
	}
	return nil
}
func GetOrderByUserId(userid int32) *Order {
	if userid != 0 {
		sql := "SELECT id,  book_id, transaction_id, coalesce(date, ''), coalesce(order_status, ''), coalesce(delivery_method, '') FROM " + db.OrderTable + " WHERE user_id=$1"
		row := connection.QueryRow(sql, userid)
		order := &Order{}
		user := &User{}
		user.id = userid
		book := &Book{}
		transaction := &Transaction{}
		err := row.Scan(&order.id, &book.id, &transaction.id, &order.date, &order.order_status, &order.delivery_method)
		order.user = GetUserById(user.id)
		order.book = GetBookById(book.id)
		order.transaction = GetTransactionById(transaction.id)
		if err != nil {
			clean.Error(err)
			return nil
		}
		return order
	}
	return nil
}
func DeleteOrder(id int32) bool {
	if id != 0 {
		sql := "DELETE FROM " + db.OrderTable + " WHERE id=$1"
		_, err := connection.Exec(sql, id)
		if err != nil {
			clean.Error(err)
			return false
		}
		return true
	}
	return false
}
func DeleteOrderByBookId(bookid int32) bool {
	if bookid != 0 {
		sql := "DELETE FROM " + db.OrderTable + " WHERE book_id=$1"
		_, err := connection.Exec(sql, bookid)
		if err != nil {
			clean.Error(err)
			return false
		}
		return true
	}
	return false
}
func DeleteOrderByUserId(userid int32) bool {
	if userid != 0 {
		sql := "DELETE FROM " + db.OrderTable + " WHERE user_id=$1"
		_, err := connection.Exec(sql, userid)
		if err != nil {
			clean.Error(err)
			return false
		}
		return true
	}
	return false
}
func CompleteOrder(order *Order) bool {
	if order != nil {
		order.order_status = "completed"
		sql := "UPDATE " + db.OrderTable + " SET order_status=$1 WHERE id=$2"
		_, err := connection.Exec(sql, order.order_status, order.id)
		if err != nil {
			clean.Error(err)
			return false
		}
		return true
	}
	return false
}
func GetOrders(page int32, count int32, sinceID int32) ([]*Order, bool, int32) {
	sql := "SELECT id,  user_id, book_id, transaction_id, coalesce(date, ''), coalesce(order_status, ''), coalesce(delivery_method, '') FROM " + db.OrderTable
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
	orders := []*Order{}
	for rows.Next() {
		order := &Order{}
		user := &User{}
		book := &Book{}
		transaction := &Transaction{}
		err = rows.Scan(&order.id, &user.id, &book.id, &transaction.id, &order.date, &order.order_status, &order.delivery_method)
		order.user = GetUserById(user.id)
		order.book = GetBookById(book.id)
		order.transaction = GetTransactionById(transaction.id)
		if err != nil {
			clean.Error(err)
			continue
		}
		orders = append(orders, order)
	}
	var hasMore bool
	var nextPagesCount int32
	ordersCount := int32(len(orders))
	if (sinceID != 0 && count > 0) || (page <= 0 && count > 0) {
		if ordersCount > count {
			hasMore = true
			orders = orders[:count]
		}
	} else if page > 0 {
		if ordersCount > count && count > 0 {
			hasMore = true
			nextPagesCount = int32(math.Ceil(float64(ordersCount)/float64(count))) - 1
			orders = orders[:count]
		}
	}
	return orders, hasMore, nextPagesCount
}

func GetOrdersCount() int {
	sql := "SELECT COUNT(*) FROM " + db.OrderTable
	row := connection.QueryRow(sql)
	var count int64
	err := row.Scan(&count)
	if err != nil {
		clean.Error(err)
		return 0
	}
	return int(count)
}
func PrepareAndValidateOrder(order *Order) error {
	if order != nil {
		if order.date == "" {
			return errors.New("date_is_required")
		}
		if order.delivery_method == "" {
			return errors.New("delivery_method_is_required")
		}

		return nil
	}
	return errors.New("invalid_data")
}
