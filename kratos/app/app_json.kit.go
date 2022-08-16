// Package apputil
// from github.com/go-kratos/kratos/v2/encoding/json
package apputil

import (
	stdjson "encoding/json"
	"reflect"

	"github.com/go-kratos/kratos/v2/encoding/json"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var (
	// MarshalOptions is a configurable JSON format marshaller.
	MarshalOptions = protojson.MarshalOptions{
		//UseProtoNames:   true,
		//UseEnumNumbers:  true,
		EmitUnpopulated: true,
	}
	// UnmarshalOptions is a configurable JSON format parser.
	UnmarshalOptions = protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
)

// SetJSONMarshalOptions 设置json编码选项
func SetJSONMarshalOptions(opt *protojson.MarshalOptions) {
	json.MarshalOptions = *opt
}

// SetJSONUnmarshalOptions 设置json编码选项
func SetJSONUnmarshalOptions(opt *protojson.UnmarshalOptions) {
	json.UnmarshalOptions = *opt
}

// MarshalJSON 编码json
func MarshalJSON(v interface{}) ([]byte, error) {
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
func UnmarshalJSON(data []byte, v interface{}) error {
	switch m := v.(type) {
	case stdjson.Unmarshaler:
		return m.UnmarshalJSON(data)
	case proto.Message:
		return UnmarshalOptions.Unmarshal(data, m)
	default:
		rv := reflect.ValueOf(v)
		for rv := rv; rv.Kind() == reflect.Ptr; {
			if rv.IsNil() {
				rv.Set(reflect.New(rv.Type().Elem()))
			}
			rv = rv.Elem()
		}
		if m, ok := reflect.Indirect(rv).Interface().(proto.Message); ok {
			return UnmarshalOptions.Unmarshal(data, m)
		}
		return stdjson.Unmarshal(data, m)
	}
}
