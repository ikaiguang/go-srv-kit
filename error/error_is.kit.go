package errorutil

// IsCode .
func IsCode(err error, code int) bool {
	return Code(err) == code
}

// IsReason .
func IsReason(err error, reason string) bool {
	return Reason(err) == reason
}
