package models

import (
	"net/url"
	"strconv"
	"strings"
)

type Pagination struct {
	Page           int32
	NextPagesCount int32
}

func NewPagination(currentPage int32, nextPagesCount int32) *Pagination {
	if currentPage <= 0 {
		currentPage = 1
	}
	return &Pagination{Page: currentPage, NextPagesCount: nextPagesCount}
}

func (this *Pagination) Nav(URL string, urlQueryParams map[string]string) string {
	if URL != "" {
		urlQueryParamsStr := ""
		if urlQueryParams != nil && len(urlQueryParams) > 0 {
			for k, v := range urlQueryParams {
				k = strings.TrimSpace(k)
				v = strings.TrimSpace(v)
				if k != "" && v != "" {
					urlQueryParamsStr += "&" + k + "=" + url.QueryEscape(v)
				}
			}
		}
		var pageItemsCount int32 = 3
		out := `<nav id="pagination" aria-label="Page navigation">
                <ul class="pagination">`
		out += `<li class="page-item page-arrow`
		if this.Page-1 <= 0 {
			out += " disabled"
		}
		out += `">`
		if this.Page-1 <= 0 {
			out += `<span>&laquo;</span>`
		} else {
			out += `<a class="page-link" href="`
			out += URL + "?page=" + strconv.Itoa(int(this.Page-1)) + urlQueryParamsStr
			out += `" >&laquo;</a>`
		}
		out += `</li>`

		for i := pageItemsCount; i > 0; i-- {
			if this.Page-i > 0 {
				out += `<li class="page-item"><a class="page-link" href="`
				out += URL + "?page=" + strconv.Itoa(int(this.Page-i)) + urlQueryParamsStr
				out += `">`
				out += strconv.Itoa(int(this.Page - i))
				out += `</a></li>`
			}
		}
		out += `<li class="page-item active"><a class="page-link" href="`
		out += URL + "?page=" + strconv.Itoa(int(this.Page)) + urlQueryParamsStr
		out += `">`
		out += strconv.Itoa(int(this.Page))
		out += `</a></li>`

		var j int32
		for j = 1; j <= pageItemsCount; j++ {
			if this.Page+j <= this.Page+this.NextPagesCount {
				out += `<li class="page-item"><a class="page-link" href="`
				out += URL + "?page=" + strconv.Itoa(int(this.Page+j)) + urlQueryParamsStr
				out += `">`
				out += strconv.Itoa(int(this.Page + j))
				out += `</a></li>`
			}
		}
		out += `<li class="page-item page-arrow`
		if this.Page+1 > this.Page+this.NextPagesCount {
			out += " disabled"
		}
		out += `">`
		if this.Page+1 > this.Page+this.NextPagesCount {
			out += `<span>&raquo;</span>`
		} else {
			out += `<a class="page-link" href="`
			out += URL + "?page=" + strconv.Itoa(int(this.Page+1)) + urlQueryParamsStr
			out += `" >&raquo;</a>`
		}
		out += `</li>`
		out += `</ul></nav>`
		return out
	}
	return ""
}
