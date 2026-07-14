package jsonpkg

import (
	"encoding/json"

	"github.com/valyala/bytebufferpool"
)

// MarshalWithoutEscapeHTML ...
func MarshalWithoutEscapeHTML(data any) ([]byte, error) {
	buffer := bytebufferpool.Get()
	defer bytebufferpool.Put(buffer)

	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(data)
	// 在归还 buffer 前复制数据，避免数据竞争
	result := make([]byte, buffer.Len())
	copy(result, buffer.Bytes())
	return result, err
}

// MarshalIndentWithoutEscapeHTML ...
func MarshalIndentWithoutEscapeHTML(data any, prefix, indent string) ([]byte, error) {
	buffer := bytebufferpool.Get()
	defer bytebufferpool.Put(buffer)

	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent(prefix, indent)
	err := encoder.Encode(data)
	// 在归还 buffer 前复制数据，避免数据竞争
	result := make([]byte, buffer.Len())
	copy(result, buffer.Bytes())
	return result, err
}
