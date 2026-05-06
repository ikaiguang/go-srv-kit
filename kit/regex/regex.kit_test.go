package regexpkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count 1 ./regex -run TestIsPhone
func TestIsPhone(t *testing.T) {
	tests := []struct {
		name  string
		phone string
		want  bool
	}{
		{"有效手机号", "13800138000", true},
		{"有效手机号_19开头", "19912345678", true},
		{"有效手机号_15开头", "15012345678", true},
		{"10开头无效", "10012345678", false},
		{"位数不足", "1380013800", false},
		{"位数过多", "138001380001", false},
		{"空字符串", "", false},
		{"非数字", "abcdefghijk", false},
		{"含空格", "138 0013 8000", false},
		{"座机号", "02112345678", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsPhone(tt.phone))
		})
	}
}

// go test -v -count 1 ./regex -run TestIsEmail
func TestIsEmail(t *testing.T) {
	tests := []struct {
		name  string
		email string
		want  bool
	}{
		{"有效邮箱", "test@example.com", true},
		{"有效邮箱_含点", "user.name@example.com", true},
		{"有效邮箱_含加号", "user+tag@example.com", true},
		{"有效邮箱_子域名", "test@sub.example.com", true},
		{"缺少@", "testexample.com", false},
		{"缺少域名", "test@", false},
		{"缺少用户名", "@example.com", false},
		{"空字符串", "", false},
		{"含空格", "test @example.com", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsEmail(tt.email))
		})
	}
}

// go test -v -count 1 ./regex -run TestIsIDCard
func TestIsIDCard(t *testing.T) {
	tests := []struct {
		name string
		id   string
		want bool
	}{
		{"18位有效身份证", "110101199001011234", true},
		{"18位末尾X", "11010119900101123X", true},
		{"15位有效身份证", "110101900101123", true},
		{"位数不足", "1101011990010112", false},
		{"位数过多", "1101011990010112345", false},
		{"空字符串", "", false},
		{"含字母", "11010119900101ABCD", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsIDCard(tt.id))
		})
	}
}

// go test -v -count 1 ./regex -run TestIsPostCode
func TestIsPostCode(t *testing.T) {
	tests := []struct {
		name string
		code string
		want bool
	}{
		{"有效邮编", "100000", true},
		{"有效邮编_上海", "200000", true},
		{"位数不足", "10000", false},
		{"位数过多", "1000001", false},
		{"空字符串", "", false},
		{"含字母", "10000a", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsPostCode(tt.code))
		})
	}
}
