package jsonpkg

import (
	"encoding/json"

	bufferpkg "github.com/ikaiguang/go-srv-kit/kit/buffer"
)

// MarshalWithoutEscapeHTML ...
func MarshalWithoutEscapeHTML(data interface{}) ([]byte, error) {
	buffer := bufferpkg.GetBuffer()
	defer bufferpkg.PutBuffer(buffer)

	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(data)
	return buffer.Bytes(), err
}

// MarshalIndentWithoutEscapeHTML ...
func MarshalIndentWithoutEscapeHTML(data interface{}, prefix, indent string) ([]byte, error) {
	buffer := bufferpkg.GetBuffer()
	defer bufferpkg.PutBuffer(buffer)

	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent(prefix, indent)
	err := encoder.Encode(data)
	return buffer.Bytes(), err
}
