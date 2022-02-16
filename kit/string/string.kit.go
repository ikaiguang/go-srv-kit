package stringutil

import (
	"strings"
)

// ToSnake example : XxYy to xx_yy ; XxYY to xx_yy
func ToSnake(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		if i > 0 && s[i] >= 'A' && s[i] <= 'Z' && j {
			data = append(data, '_')
		}
		if s[i] != '_' {
			j = true
		}
		data = append(data, s[i])
	}
	return strings.ToLower(string(data[:]))
}

// ToCamel example : xx_yy to XxYy
func ToCamel(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}
