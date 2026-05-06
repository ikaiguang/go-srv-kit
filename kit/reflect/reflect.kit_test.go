package reflectpkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count 1 ./reflect -run TestIsDefaultValue
func TestIsDefaultValue(t *testing.T) {
	tests := []struct {
		name  string
		value any
		want  bool
	}{
		// 布尔
		{"bool_zero", false, true},
		{"bool_nonzero", true, false},
		// 整型
		{"int_zero", 0, true},
		{"int_nonzero", 1, false},
		{"int8_zero", int8(0), true},
		{"int16_zero", int16(0), true},
		{"int32_zero", int32(0), true},
		{"int64_zero", int64(0), true},
		{"int64_nonzero", int64(42), false},
		// 无符号整型
		{"uint_zero", uint(0), true},
		{"uint_nonzero", uint(1), false},
		{"uint8_zero", uint8(0), true},
		{"uint16_zero", uint16(0), true},
		{"uint32_zero", uint32(0), true},
		{"uint64_zero", uint64(0), true},
		// 浮点
		{"float32_zero", float32(0), true},
		{"float32_nonzero", float32(3.14), false},
		{"float64_zero", float64(0), true},
		{"float64_nonzero", float64(3.14), false},
		// 复数
		{"complex64_zero", complex64(0), true},
		{"complex64_nonzero", complex64(1 + 2i), false},
		{"complex128_zero", complex128(0), true},
		{"complex128_nonzero", complex128(1 + 2i), false},
		// 字符串
		{"string_zero", "", true},
		{"string_nonzero", "hello", false},
		// 切片
		{"slice_nil", ([]int)(nil), true},
		{"slice_nonnil", []int{}, false},
		// Map
		{"map_nil", (map[string]int)(nil), true},
		{"map_nonnil", map[string]int{}, false},
		// 指针
		{"nil", nil, true},
		{"ptr_nil", (*int)(nil), true},
		{"ptr_nonnil", new(int), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsDefaultValue(tt.value)
			assert.Equal(t, tt.want, got)
		})
	}
}

// go test -v -count 1 ./reflect -run TestIsEmpty
func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name  string
		value any
		want  bool
	}{
		{"nil", nil, true},
		{"empty_string", "", true},
		{"nonempty_string", "hello", false},
		{"zero_int", 0, true},
		{"nonzero_int", 42, false},
		{"nil_slice", ([]int)(nil), true},
		{"empty_slice", []int{}, true},
		{"nonempty_slice", []int{1}, false},
		{"nil_map", (map[string]int)(nil), true},
		{"empty_map", map[string]int{}, true},
		{"nonempty_map", map[string]int{"a": 1}, false},
		{"nil_ptr", (*int)(nil), true},
		{"nonnil_ptr_zero", new(int), true},
		{"false_bool", false, true},
		{"true_bool", true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsEmpty(tt.value)
			assert.Equal(t, tt.want, got)
		})
	}
}

// go test -v -count 1 ./reflect -run TestSwapObject
func TestSwapObject(t *testing.T) {
	t.Run("正常交换", func(t *testing.T) {
		dst := 10
		src := 20
		ok := SwapObject(&dst, &src)
		assert.True(t, ok)
		assert.Equal(t, 20, dst)
	})

	t.Run("dst非指针返回false", func(t *testing.T) {
		dst := 10
		src := 20
		ok := SwapObject(dst, &src)
		assert.False(t, ok)
	})

	t.Run("类型不匹配返回false", func(t *testing.T) {
		dst := 10
		src := "hello"
		ok := SwapObject(&dst, &src)
		assert.False(t, ok)
	})

	t.Run("nil返回false", func(t *testing.T) {
		assert.False(t, SwapObject(nil, nil))
	})

	t.Run("结构体交换", func(t *testing.T) {
		type Foo struct {
			Name string
			Age  int
		}
		dst := Foo{Name: "old", Age: 1}
		src := Foo{Name: "new", Age: 2}
		ok := SwapObject(&dst, &src)
		assert.True(t, ok)
		assert.Equal(t, "new", dst.Name)
		assert.Equal(t, 2, dst.Age)
	})
}

// go test -v -count 1 ./reflect -run TestNewObject
func TestNewObject(t *testing.T) {
	t.Run("从指针创建", func(t *testing.T) {
		src := &struct{ Name string }{Name: "test"}
		obj := NewObject(src)
		assert.NotNil(t, obj)
	})

	t.Run("从值创建", func(t *testing.T) {
		src := struct{ Name string }{Name: "test"}
		obj := NewObject(src)
		assert.NotNil(t, obj)
	})

	t.Run("nil返回nil", func(t *testing.T) {
		assert.Nil(t, NewObject(nil))
	})
}
