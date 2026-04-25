package ptrpkg

// Ptr 将任意类型的值转换为指针
func Ptr[T any](v T) *T {
	return &v
}

// Value 将指针转换为值，nil 指针返回零值
func Value[T comparable](p *T) T {
	if p == nil {
		var zero T
		return zero
	}
	return *p
}

// Slice 将值切片转换为指针切片
func Slice[T any](a []T) []*T {
	if a == nil {
		return nil
	}
	res := make([]*T, len(a))
	for i := range a {
		res[i] = &a[i]
	}
	return res
}

// ValueSlice 将指针切片转换为值切片
func ValueSlice[T any](a []*T) []T {
	if a == nil {
		return nil
	}
	res := make([]T, len(a))
	for i := range a {
		if a[i] != nil {
			res[i] = *a[i]
		}
	}
	return res
}
