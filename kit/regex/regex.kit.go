package regexpkg

import (
	"regexp"
)

var (
	phoneRegex    = regexp.MustCompile(`^1[1-9]\d{9}$`)
	emailRegex    = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	idCardRegex   = regexp.MustCompile(`^(\d{17}[\dX]|\d{15})$`)
	postCodeRegex = regexp.MustCompile(`^\d{6}$`)
)

func IsValidPhone(phone string) bool {
	return phoneRegex.MatchString(phone)
}

func IsValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func IsIDCard(id string) bool {
	return idCardRegex.MatchString(id)
}

func IsPostCode(code string) bool {
	return postCodeRegex.MatchString(code)
}
