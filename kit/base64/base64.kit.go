package base64util

import "encoding/base64"

// ExampleDecodeString ...
func ExampleDecodeString(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

// ExampleEncodeToString ...
func ExampleEncodeToString(s []byte) string {
	return base64.StdEncoding.EncodeToString(s)
}
