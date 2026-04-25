package ptrpkg

import (
	"testing"
)

// === 泛型函数测试 ===

func TestPtr(t *testing.T) {
	// 测试 string
	s := "hello"
	p := Ptr(s)
	if *p != s {
		t.Errorf("Ptr(string): got %v, want %v", *p, s)
	}

	// 测试 int
	i := 42
	pi := Ptr(i)
	if *pi != i {
		t.Errorf("Ptr(int): got %v, want %v", *pi, i)
	}

	// 测试 float64
	f := 3.14
	pf := Ptr(f)
	if *pf != f {
		t.Errorf("Ptr(float64): got %v, want %v", *pf, f)
	}

	// 测试 bool
	b := true
	pb := Ptr(b)
	if *pb != b {
		t.Errorf("Ptr(bool): got %v, want %v", *pb, b)
	}
}

func TestValue(t *testing.T) {
	// 测试正常值
	s := "hello"
	if got := Value(&s); got != s {
		t.Errorf("Value(&string): got %v, want %v", got, s)
	}

	// 测试 nil 指针返回零值
	var nilStr *string
	if got := Value(nilStr); got != "" {
		t.Errorf("Value(nil *string): got %q, want %q", got, "")
	}

	var nilInt *int
	if got := Value(nilInt); got != 0 {
		t.Errorf("Value(nil *int): got %v, want 0", got)
	}

	var nilBool *bool
	if got := Value(nilBool); got != false {
		t.Errorf("Value(nil *bool): got %v, want false", got)
	}
}

func TestSlice(t *testing.T) {
	// 测试正常切片
	a := []int{1, 2, 3}
	result := Slice(a)
	if len(result) != len(a) {
		t.Fatalf("Slice: len got %d, want %d", len(result), len(a))
	}
	for i, v := range result {
		if *v != a[i] {
			t.Errorf("Slice[%d]: got %v, want %v", i, *v, a[i])
		}
	}

	// 测试 nil 切片返回 nil
	var nilSlice []int
	if got := Slice(nilSlice); got != nil {
		t.Errorf("Slice(nil): got %v, want nil", got)
	}

	// 测试空切片返回空切片（非 nil）
	emptySlice := []int{}
	got := Slice(emptySlice)
	if got == nil {
		t.Errorf("Slice([]int{}): got nil, want non-nil empty slice")
	}
	if len(got) != 0 {
		t.Errorf("Slice([]int{}): len got %d, want 0", len(got))
	}
}

func TestValueSlice(t *testing.T) {
	// 测试正常切片
	v1, v2, v3 := 1, 2, 3
	a := []*int{&v1, &v2, &v3}
	result := ValueSlice(a)
	if len(result) != len(a) {
		t.Fatalf("ValueSlice: len got %d, want %d", len(result), len(a))
	}
	for i, v := range result {
		if v != *a[i] {
			t.Errorf("ValueSlice[%d]: got %v, want %v", i, v, *a[i])
		}
	}

	// 测试 nil 切片返回 nil
	var nilSlice []*int
	if got := ValueSlice(nilSlice); got != nil {
		t.Errorf("ValueSlice(nil): got %v, want nil", got)
	}

	// 测试包含 nil 元素的切片
	v4 := 10
	mixed := []*int{&v4, nil, &v4}
	mixedResult := ValueSlice(mixed)
	if mixedResult[0] != 10 || mixedResult[1] != 0 || mixedResult[2] != 10 {
		t.Errorf("ValueSlice(mixed): got %v, want [10 0 10]", mixedResult)
	}
}

// === Deprecated 函数向后兼容性测试 ===

func TestDeprecatedStringFunctions(t *testing.T) {
	s := "test"
	p := String(s)
	if *p != s {
		t.Errorf("String: got %v, want %v", *p, s)
	}
	if got := StringValue(p); got != s {
		t.Errorf("StringValue: got %v, want %v", got, s)
	}
	if got := StringValue(nil); got != "" {
		t.Errorf("StringValue(nil): got %q, want %q", got, "")
	}
}

func TestDeprecatedIntFunctions(t *testing.T) {
	i := 42
	p := Int(i)
	if *p != i {
		t.Errorf("Int: got %v, want %v", *p, i)
	}
	if got := IntValue(p); got != i {
		t.Errorf("IntValue: got %v, want %v", got, i)
	}
	if got := IntValue(nil); got != 0 {
		t.Errorf("IntValue(nil): got %v, want 0", got)
	}
}

func TestDeprecatedInt32Functions(t *testing.T) {
	var i int32 = 32
	p := Int32(i)
	if *p != i {
		t.Errorf("Int32: got %v, want %v", *p, i)
	}
	if got := Int32Value(p); got != i {
		t.Errorf("Int32Value: got %v, want %v", got, i)
	}
	if got := Int32Value(nil); got != 0 {
		t.Errorf("Int32Value(nil): got %v, want 0", got)
	}
}

func TestDeprecatedInt64Functions(t *testing.T) {
	var i int64 = 64
	p := Int64(i)
	if *p != i {
		t.Errorf("Int64: got %v, want %v", *p, i)
	}
	if got := Int64Value(p); got != i {
		t.Errorf("Int64Value: got %v, want %v", got, i)
	}
	if got := Int64Value(nil); got != 0 {
		t.Errorf("Int64Value(nil): got %v, want 0", got)
	}
}

func TestDeprecatedBoolFunctions(t *testing.T) {
	b := true
	p := Bool(b)
	if *p != b {
		t.Errorf("Bool: got %v, want %v", *p, b)
	}
	if got := BoolValue(p); got != b {
		t.Errorf("BoolValue: got %v, want %v", got, b)
	}
	if got := BoolValue(nil); got != false {
		t.Errorf("BoolValue(nil): got %v, want false", got)
	}
}

func TestDeprecatedFloat64Functions(t *testing.T) {
	f := 3.14
	p := Float64(f)
	if *p != f {
		t.Errorf("Float64: got %v, want %v", *p, f)
	}
	if got := Float64Value(p); got != f {
		t.Errorf("Float64Value: got %v, want %v", got, f)
	}
	if got := Float64Value(nil); got != 0 {
		t.Errorf("Float64Value(nil): got %v, want 0", got)
	}
}

func TestDeprecatedSliceFunctions(t *testing.T) {
	// IntSlice / IntValueSlice
	ints := []int{1, 2, 3}
	intPtrs := IntSlice(ints)
	if len(intPtrs) != 3 || *intPtrs[0] != 1 {
		t.Errorf("IntSlice: unexpected result")
	}
	intVals := IntValueSlice(intPtrs)
	if len(intVals) != 3 || intVals[0] != 1 {
		t.Errorf("IntValueSlice: unexpected result")
	}

	// StringSlice / StringSliceValue
	strs := []string{"a", "b"}
	strPtrs := StringSlice(strs)
	if len(strPtrs) != 2 || *strPtrs[0] != "a" {
		t.Errorf("StringSlice: unexpected result")
	}
	strVals := StringSliceValue(strPtrs)
	if len(strVals) != 2 || strVals[0] != "a" {
		t.Errorf("StringSliceValue: unexpected result")
	}

	// nil 切片
	if got := IntSlice(nil); got != nil {
		t.Errorf("IntSlice(nil): got %v, want nil", got)
	}
	if got := IntValueSlice(nil); got != nil {
		t.Errorf("IntValueSlice(nil): got %v, want nil", got)
	}
}
