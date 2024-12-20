package regexpkg

import (
	"regexp"
)

var (
	phoneRegex = regexp.MustCompile(`^1[1-9]\d{9}$`)
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

func IsValidPhone(phone string) bool {
	return phoneRegex.MatchString(phone)
}

func IsValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}
