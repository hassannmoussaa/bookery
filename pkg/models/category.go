package models

import (
	"errors"
	"math"
	"strconv"
	"strings"

	"github.com/hassannmoussaa/bookery/pkg/db"
	"github.com/hassannmoussaa/pill.go/clean"
)

type Category struct {
	id            int32
	category_name string
}

func (this *Category) ToMap(prefix string, excluded bool, fields ...string) map[string]interface{} {
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
	if this.CategoryName() != "" && check(prefix+"category_name", fields...) {
		result["category_name"] = this.CategoryName()
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

func (this *Category) ID() int32 {
	return this.id
}
func (this *Category) CategoryName() string {
	return this.category_name
}

func (this *Category) SetID(value int32) {
	this.id = value
}
func (this *Category) SetCategoryName(value string) {
	this.category_name = value
}

func AddCategory(category *Category) *Category {
	if category != nil {
		sql := "INSERT INTO " + db.CategoryTable + " (category_name) VALUES ($1) RETURNING id;"
		row := connection.QueryRow(sql, category.category_name)
		err := row.Scan(&category.id)
		if err != nil {
			clean.Error(err)
			return nil
		}
		return category
	}
	return nil
}
func GetCategoryById(id int32) *Category {
	if id != 0 {
		sql := "SELECT coalesce(category_name, '') FROM " + db.CategoryTable + " WHERE id=$1"
		row := connection.QueryRow(sql, id)
		category := &Category{}
		category.id = id
		err := row.Scan(&category.category_name)
		if err != nil {
			clean.Error(err)
			return nil
		}
		return category
	}
	return nil
}
func DeleteCategory(id int32) bool {
	if id != 0 {
		sql := "DELETE FROM " + db.CategoryTable + " WHERE id=$1"
		_, err := connection.Exec(sql, id)
		if err != nil {
			clean.Error(err)
			return false
		}
		return true
	}
	return false
}
func GetCategories(page int32, count int32, sinceID int32) ([]*Category, bool, int32) {
	sql := "SELECT id, coalesce(category_name, '') FROM " + db.CategoryTable
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
	categories := []*Category{}
	for rows.Next() {
		category := &Category{}

		err = rows.Scan(&category.id, &category.category_name)

		if err != nil {
			clean.Error(err)
			continue
		}
		categories = append(categories, category)
	}
	var hasMore bool
	var nextPagesCount int32
	categoriesCount := int32(len(categories))
	if (sinceID != 0 && count > 0) || (page <= 0 && count > 0) {
		if categoriesCount > count {
			hasMore = true
			categories = categories[:count]
		}
	} else if page > 0 {
		if categoriesCount > count && count > 0 {
			hasMore = true
			nextPagesCount = int32(math.Ceil(float64(categoriesCount)/float64(count))) - 1
			categories = categories[:count]
		}
	}
	return categories, hasMore, nextPagesCount
}
func GetCategoriesCount() int {
	sql := "SELECT COUNT(*) FROM " + db.CategoryTable
	row := connection.QueryRow(sql)
	var count int64
	err := row.Scan(&count)
	if err != nil {
		clean.Error(err)
		return 0
	}
	return int(count)
}
func PrepareAndValidateCategory(category *Category) error {
	if category != nil {
		if category.category_name == "" {
			return errors.New("category_name_is_required")
		}

		return nil
	}
	return errors.New("invalid_data")
}
