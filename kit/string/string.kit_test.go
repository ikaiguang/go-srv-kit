package stringpkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToSnake(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"驼峰转蛇形", "XxYy", "xx_yy"},
		{"连续大写", "XxYY", "xx_y_y"},
		{"已有下划线", "X_Y_Z", "x__y__z"},
		{"单个字符", "A", "a"},
		{"空字符串", "", ""},
		{"全小写", "hello", "hello"},
		{"全大写", "ABC", "a_b_c"},
		{"混合", "getUserName", "get_user_name"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToSnake(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestToCamel(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"蛇形转驼峰", "xx_yy", "XxYy"},
		{"双下划线", "a__b__c", "A_B_C"},
		{"单个字符", "a", "A"},
		{"空字符串", "", ""},
		{"已是驼峰", "XxYy", "XxYy"},
		{"多段", "get_user_name", "GetUserName"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToCamel(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestToString(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  string
	}{
		{"nil", nil, ""},
		{"string", "hello", "hello"},
		{"int", 42, "42"},
		{"int8", int8(8), "8"},
		{"int16", int16(16), "16"},
		{"int32", int32(32), "32"},
		{"int64", int64(64), "64"},
		{"uint", uint(10), "10"},
		{"uint8", uint8(8), "8"},
		{"uint16", uint16(16), "16"},
		{"uint32", uint32(32), "32"},
		{"uint64", uint64(64), "64"},
		{"float32", float32(1.5), "1.5"},
		{"float64", float64(3.14), "3.14"},
		{"bytes", []byte("hello"), "hello"},
		{"bool", true, "true"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToString(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
