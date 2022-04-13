package gormutil

import (
	"regexp"
)

var (
	// regColumn 正则表达式:列
	regColumn = regexp.MustCompile("^[A-Za-z-_]+$")
)

// IsValidField 判断是否为有效的字段名
func IsValidField(field string) bool {
	return regColumn.MatchString(field)
}
