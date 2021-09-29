package reflectutil

import "reflect"

// IsDefaultValue golang default value
func IsDefaultValue(i interface{}) bool {
	value := reflect.ValueOf(i)

	switch value.Kind() {
	case reflect.Bool:
		if !value.Bool() {
			return true
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if value.Int() == 0 {
			return true
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if value.Uint() == 0 {
			return true
		}
	case reflect.Float32, reflect.Float64:
		if value.Float() == 0 {
			return true
		}
	case reflect.Complex64, reflect.Complex128:
		if value.Complex() == 0 {
			return true
		}
	case reflect.String:
		if value.Len() == 0 {
			return true
		}
	case reflect.Map, reflect.Ptr, reflect.Slice:
		if value.IsNil() {
			return true
		}
	}
	return false
}
