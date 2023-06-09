package operatorpkg

// Ternary 三元表达式 = cond ? v1 : v2
func Ternary[T any](cond bool, v1 T, v2 T) T {
	if cond {
		return v1
	}
	return v2
}
