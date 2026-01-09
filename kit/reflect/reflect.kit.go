package reflectpkg

import "reflect"

// IsDefaultValue golang default value
func IsDefaultValue(i interface{}) bool {
	value := reflect.ValueOf(i)

	switch value.Kind() {
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Complex64, reflect.Complex128:
		return value.Complex() == 0
	case reflect.String:
		return value.Len() == 0
	case reflect.Map, reflect.Ptr, reflect.Slice:
		return value.IsNil()
	default:
		return false
	}
}

// IsEmpty gets whether the specified object is considered empty or not.
func IsEmpty(object interface{}) bool {
	// get nil case out of the way
	if object == nil {
		return true
	}

	objValue := reflect.ValueOf(object)

	switch objValue.Kind() {
	// collection types are empty when they have no element
	case reflect.Chan, reflect.Map, reflect.Slice:
		return objValue.Len() == 0
	// pointers are empty if nil or if the value they point to is empty
	case reflect.Ptr:
		if objValue.IsNil() {
			return true
		}
		deref := objValue.Elem().Interface()
		return IsEmpty(deref)
	// for all other types, compare against the zero value
	// array types are empty when they match their zero-initialized state
	default:
		zero := reflect.Zero(objValue.Type())
		return reflect.DeepEqual(object, zero.Interface())
	}
}

func SwapObject(dst, src interface{}) bool {
	dstType := reflect.TypeOf(dst)
	if dstType.Kind() != reflect.Ptr {
		return false
	}
	srcType := reflect.TypeOf(src)
	if srcType.Kind() != reflect.Ptr {
		srcType = reflect.PointerTo(srcType)
	}
	if dstType.String() != srcType.String() {
		return false
	}
	dstValue := reflect.ValueOf(dst)
	if dstValue.Kind() == reflect.Ptr {
		dstValue = dstValue.Elem()
	}
	if !dstValue.IsValid() {
		return false
	}
	srcValue := reflect.ValueOf(src)
	if srcValue.Kind() == reflect.Ptr {
		srcValue = srcValue.Elem()
	}
	if !srcValue.IsValid() {
		return false
	}
	dstValue.Set(srcValue)
	return true
}

func NewObject(dst interface{}) interface{} {
	typ := reflect.TypeOf(dst)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return reflect.New(typ).Interface()
}
