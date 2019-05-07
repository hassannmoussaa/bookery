package models

import (
	"strings"

	"github.com/jackc/pgx"
)

var (
	connection         *pgx.ConnPool
	staticFilesURLPath string
	uploadsURLPath     string
)

func Included(item string, list ...string) bool {
	if item != "" && list != nil && len(list) > 0 {
		for _, v := range list {
			if v == item || strings.HasPrefix(v, item) {
				return true
			}
		}
		return false
	}
	return true
}

func Excluded(item string, list ...string) bool {
	if item != "" && list != nil && len(list) > 0 {
		for _, v := range list {
			if v == item {
				return false
			}
		}
		return true
	}
	return true
}

func Init(conn *pgx.ConnPool, StaticFilesURLPath string, UploadsURLPath string) {
	connection = conn
	staticFilesURLPath = StaticFilesURLPath
	uploadsURLPath = UploadsURLPath
}

func EscapeQuotes(text string) string {
	return strings.Replace(strings.Replace(text, "'", "''", -1), `""`, "", -1)
}
