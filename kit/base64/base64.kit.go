package base64util

import "encoding/base64"

// Encode 编码
func Encode(src []byte) (dst []byte) {
	dst = make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(dst, src)
	return dst
}

// Decode 解码
func Decode(src []byte) (dst []byte, err error) {
	dst = make([]byte, base64.StdEncoding.DecodedLen(len(src)))
	n, err := base64.StdEncoding.Decode(dst, src)
	if err != nil {
		return nil, err
	}
	return dst[:n], nil
}

// ExampleDecodeString ...
func ExampleDecodeString(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

// ExampleEncodeToString ...
func ExampleEncodeToString(s []byte) string {
	return base64.StdEncoding.EncodeToString(s)
}
