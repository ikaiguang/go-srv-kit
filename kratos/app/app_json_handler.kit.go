package apppkg

import (
	stdjson "encoding/json"
	"reflect"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var (
	// MarshalOptions is a configurable JSON format marshaller.
	MarshalOptions = protojson.MarshalOptions{
		UseProtoNames:   true,
		UseEnumNumbers:  true,
		EmitUnpopulated: true,
	}
	// UnmarshalOptions is a configurable JSON format parser.
	UnmarshalOptions = protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
)

// SetJSONMarshalOptions 设置json编码选项
func SetJSONMarshalOptions(opt *protojson.MarshalOptions) {
	if opt == nil {
		return
	}
	MarshalOptions = *opt
}

// SetJSONUnmarshalOptions 设置json编码选项
func SetJSONUnmarshalOptions(opt *protojson.UnmarshalOptions) {
	if opt == nil {
		return
	}
	UnmarshalOptions = *opt
}

// MarshalJSON 编码json
func MarshalJSON(v any) ([]byte, error) {
	switch m := v.(type) {
	case stdjson.Marshaler:
		return m.MarshalJSON()
	case proto.Message:
		return MarshalOptions.Marshal(m)
	default:
		return stdjson.Marshal(m)
	}
}

// UnmarshalJSON 解码json
func UnmarshalJSON(data []byte, v any) error {
	rv := reflect.ValueOf(v)
	if !rv.IsValid() || (rv.Kind() == reflect.Pointer && rv.IsNil()) {
		return stdjson.Unmarshal(data, v)
	}
	switch m := v.(type) {
	case stdjson.Unmarshaler:
		return m.UnmarshalJSON(data)
	case proto.Message:
		return UnmarshalOptions.Unmarshal(data, m)
	default:
		for rv := rv; rv.Kind() == reflect.Pointer; {
			if rv.IsNil() {
				if !rv.CanSet() {
					return stdjson.Unmarshal(data, v)
				}
				rv.Set(reflect.New(rv.Type().Elem()))
			}
			if rv.CanInterface() {
				if message, ok := rv.Interface().(proto.Message); ok {
					return UnmarshalOptions.Unmarshal(data, message)
				}
			}
			rv = rv.Elem()
		}
		return stdjson.Unmarshal(data, m)
	}
}
