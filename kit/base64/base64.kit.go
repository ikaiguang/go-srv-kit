package base64pkg

import "encoding/base64"

// Encryptor ...
type Encryptor interface {
	EncryptToString(plaintext string) (string, error)
	DecryptToString(ciphertext string) (string, error)
}

// B64 ...
type B64 struct{}

func (s *B64) EncryptToString(plaintext string) (string, error) {
	res := Encode([]byte(plaintext))
	return string(res), nil
}

func (s *B64) DecryptToString(ciphertext string) (string, error) {
	res, err := Decode([]byte(ciphertext))
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// Encode 编码
func Encode(src []byte) []byte {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(dst, src)
	return dst
}

// Decode 解码
func Decode(src []byte) ([]byte, error) {
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(src)))
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
