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

// IsChineseMainlandMobile validates the format of a mainland China mobile number.
func IsChineseMainlandMobile(phone string) bool {
	return phoneRegex.MatchString(phone)
}

func IsEmail(email string) bool {
	return emailRegex.MatchString(email)
}

// IsChineseCitizenIDFormat validates only the length and character format.
func IsChineseCitizenIDFormat(id string) bool {
	return idCardRegex.MatchString(id)
}

// IsChinesePostalCode validates a six-digit mainland China postal code.
func IsChinesePostalCode(code string) bool {
	return postCodeRegex.MatchString(code)
}

// Deprecated: use IsChineseMainlandMobile instead.
func IsPhone(phone string) bool { return IsChineseMainlandMobile(phone) }

// Deprecated: use IsChineseCitizenIDFormat instead.
func IsIDCard(id string) bool { return IsChineseCitizenIDFormat(id) }

// Deprecated: use IsChinesePostalCode instead.
func IsPostCode(code string) bool { return IsChinesePostalCode(code) }
