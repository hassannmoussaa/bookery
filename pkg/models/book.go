package models

import (
	"errors"
	"math"
	"strconv"
	"strings"

	"github.com/hassannmoussaa/bookery/pkg/db"
	"github.com/hassannmoussaa/pill.go/clean"
)

type Book struct {
	id          int32
	book_name   string
	category    *Category
	author_name string
	page_count  int
	quality     string
	front_image string
	back_image  string
	side_image  string
	is_verified bool
	is_recived  bool
	price       int
}

func (this *Book) ToMap(prefix string, excluded bool, fields ...string) map[string]interface{} {
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
	if this.BookName() != "" && check(prefix+"book_name", fields...) {
		result["book_name"] = this.BookName()
	}
	if this.Category().CategoryName() != "" && check(prefix+"category_name", fields...) {
		result["category_name"] = this.Category().CategoryName()
	}
	if this.AuthorName() != "" && check(prefix+"author_name", fields...) {
		result["author_name"] = this.AuthorName()
	}
	if this.PageCount() != 0 && check(prefix+"page_count", fields...) {
		result["page_count"] = this.PageCount()
	}
	if this.Quality() != "" && check(prefix+"quality", fields...) {
		result["quality"] = this.Quality()
	}
	if this.FrontImage() != "" && check(prefix+"front_image", fields...) {
		result["front_image"] = this.FrontImage()
	}
	if this.BackImage() != "" && check(prefix+"back_image", fields...) {
		result["back_image"] = this.BackImage()
	}
	if this.SideImage() != "" && check(prefix+"side_image", fields...) {
		result["side_image"] = this.SideImage()
	}
	if check(prefix+"is_verified", fields...) {
		result["is_verified"] = this.IsVerified()
	}
	if check(prefix+"is_recived", fields...) {
		result["is_recived"] = this.IsRecived()
	}
	if this.Price() != 0 && check(prefix+"price", fields...) {
		result["price"] = this.Price()
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

func (this *Book) ID() int32 {
	return this.id
}
func (this *Book) BookName() string {
	return this.book_name
}
func (this *Book) Category() *Category {
	return this.category
}
func (this *Book) AuthorName() string {
	return this.author_name
}
func (this *Book) PageCount() int {
	return this.page_count
}
func (this *Book) Quality() string {
	return this.quality
}
func (this *Book) FrontImage() string {
	return this.front_image
}
func (this *Book) BackImage() string {
	return this.back_image
}
func (this *Book) SideImage() string {
	return this.side_image
}
func (this *Book) IsVerified() bool {
	return this.is_verified
}
func (this *Book) IsRecived() bool {
	return this.is_recived
}
func (this *Book) Price() int {
	return this.price
}

func (this *Book) SetID(value int32) {
	this.id = value
}
func (this *Book) SetBookName(value string) {
	this.book_name = value
}
func (this *Book) SetCategory(value *Category) {
	this.category = value
}
func (this *Book) SetAuthorName(value string) {
	this.author_name = value
}
func (this *Book) SetPageCount(value int) {
	this.page_count = value
}
func (this *Book) SetQuality(value string) {
	this.quality = value
}
func (this *Book) SetFrontImage(value string) {
	this.front_image = value
}
func (this *Book) SetBackImage(value string) {
	this.back_image = value
}
func (this *Book) SetSideImage(value string) {
	this.side_image = value
}
func (this *Book) SetIsVerified(value bool) {
	this.is_verified = value
}
func (this *Book) SetIsRecived(value bool) {
	this.is_recived = value
}
func (this *Book) SetPrice(value int) {
	this.price = value
}

func AddBook(book *Book) *Book {
	if book != nil {
		sql := "INSERT INTO " + db.BookTable + " (book_name, category_id, author_name, page_count, quality, front_image, back_image, side_image, is_verified, is_recived, price) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id;"
		row := connection.QueryRow(sql, book.book_name, book.category.id, book.author_name, book.page_count, book.quality, book.front_image, book.back_image, book.side_image, book.is_verified, book.is_recived, book.price)
		err := row.Scan(&book.id)
		if err != nil {
			clean.Error(err)
			return nil
		}
		return book
	}
	return nil
}
func GetBookById(id int32) *Book {
	if id != 0 {
		sql := "SELECT coalesce(book_name, ''), category_id, coalesce(author_name, ''), coalesce(page_count, 0), coalesce(quality, ''), coalesce(front_image, ''), coalesce(back_image, ''), coalesce(side_image, ''), coalesce(is_verified, false), coalesce(is_recived, false), coalesce(price, 0) FROM " + db.BookTable + " WHERE id=$1"
		row := connection.QueryRow(sql, id)
		book := &Book{}
		category := &Category{}
		book.id = id
		err := row.Scan(&book.book_name, &category.id, &book.author_name, &book.page_count, &book.quality, &book.front_image, &book.back_image, &book.side_image, &book.is_verified, &book.is_recived, &book.price)
		book.category = GetCategoryById(category.id)
		if err != nil {
			clean.Error(err)
			return nil
		}
		return book
	}
	return nil
}
func GetBookByCatId(catid int32) *Book {
	if catid != 0 {
		sql := "SELECT id, coalesce(book_name, ''), coalesce(author_name, ''), coalesce(page_count, 0), coalesce(quality, ''), coalesce(front_image, ''), coalesce(back_image, ''), coalesce(side_image, ''), coalesce(is_verified, false), coalesce(is_recived, false), coalesce(price, 0) FROM " + db.BookTable + " WHERE category_id=$1"
		row := connection.QueryRow(sql, catid)
		book := &Book{}
		category := &Category{}
		category.id = catid
		err := row.Scan(&book.id, &book.book_name, &book.author_name, &book.page_count, &book.quality, &book.front_image, &book.back_image, &book.side_image, &book.is_verified, &book.is_recived, &book.price)
		book.category = GetCategoryById(category.id)
		if err != nil {
			clean.Error(err)
			return nil
		}
		return book
	}
	return nil
}
func DeleteBook(id int32) bool {
	if id != 0 {
		sql := "DELETE FROM " + db.BookTable + " WHERE id=$1"
		_, err := connection.Exec(sql, id)
		if err != nil {
			clean.Error(err)
			return false
		}
		return true
	}
	return false
}
func DeleteBookByCategory(catid int32) bool {
	if catid != 0 {
		sql := "DELETE FROM " + db.BookTable + " WHERE category_id=$1"
		_, err := connection.Exec(sql, catid)
		if err != nil {
			clean.Error(err)
			return false
		}
		return true
	}
	return false
}
func GetBooks(page int32, count int32, sinceID int32) ([]*Book, bool, int32) {
	sql := "SELECT id, coalesce(book_name, ''), category_id, coalesce(author_name, ''), coalesce(page_count, 0), coalesce(quality, ''), coalesce(front_image, ''), coalesce(back_image, ''), coalesce(side_image, ''), coalesce(is_verified, false), coalesce(is_recived, false), coalesce(price, 0) FROM " + db.BookTable
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
	books := []*Book{}
	for rows.Next() {
		book := &Book{}
		category := &Category{}
		err = rows.Scan(&book.id, &book.book_name, &category.id, &book.author_name, &book.page_count, &book.quality, &book.front_image, &book.back_image, &book.side_image, &book.is_verified, &book.is_recived, &book.price)
		book.category = GetCategoryById(category.id)
		if err != nil {
			clean.Error(err)
			continue
		}
		books = append(books, book)
	}
	var hasMore bool
	var nextPagesCount int32
	booksCount := int32(len(books))
	if (sinceID != 0 && count > 0) || (page <= 0 && count > 0) {
		if booksCount > count {
			hasMore = true
			books = books[:count]
		}
	} else if page > 0 {
		if booksCount > count && count > 0 {
			hasMore = true
			nextPagesCount = int32(math.Ceil(float64(booksCount)/float64(count))) - 1
			books = books[:count]
		}
	}
	return books, hasMore, nextPagesCount
}
func GetBooksCount() int {
	sql := "SELECT COUNT(*) FROM " + db.BookTable
	row := connection.QueryRow(sql)
	var count int64
	err := row.Scan(&count)
	if err != nil {
		clean.Error(err)
		return 0
	}
	return int(count)
}

func PrepareAndValidateBook(book *Book) error {
	if book != nil {
		if book.book_name == "" {
			return errors.New("book_name_is_required")
		}
		if book.category.category_name == "" {
			return errors.New("category_name_is_required")
		}
		if book.author_name == "" {
			return errors.New("author_name_is_required")
		}
		if book.page_count == 0 {
			return errors.New("page_count_is_required")
		}
		if book.quality == "" {
			return errors.New("quality_is_required")
		}
		if book.front_image == "" {
			return errors.New("front_image_is_required")
		}
		if book.back_image == "" {
			return errors.New("back_image_is_required")
		}
		if book.side_image == "" {
			return errors.New("side_image_is_required")
		}

		return nil
	}
	return errors.New("invalid_data")
}
