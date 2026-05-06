package randompkg

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"math/big"
)

// SecureBytes 使用 crypto/rand 生成指定长度的安全随机字节。
func SecureBytes(length int) ([]byte, error) {
	if length <= 0 {
		return []byte{}, nil
	}
	buf := make([]byte, length)
	if _, err := rand.Read(buf); err != nil {
		return nil, err
	}
	return buf, nil
}

// SecureString 使用 crypto/rand 从指定字符集中生成安全随机字符串。
func SecureString(length int, charset string) (string, error) {
	if length <= 0 {
		return "", nil
	}
	if charset == "" {
		return "", errors.New("charset is empty")
	}

	res := make([]byte, length)
	max := big.NewInt(int64(len(charset)))
	for i := range res {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		res[i] = charset[n.Int64()]
	}
	return string(res), nil
}

// SecureToken 生成 URL 安全的随机 token。
func SecureToken(length int) (string, error) {
	return SecureString(length, CharsetAlphanumeric)
}

// SecureHex 生成十六进制安全随机字符串。
func SecureHex(length int) (string, error) {
	return SecureString(length, CharsetHex)
}

// SecureBase64URL 生成 URL 安全的 base64 token。
// length 表示随机字节长度，不是最终字符串长度。
func SecureBase64URL(length int) (string, error) {
	data, err := SecureBytes(length)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(data), nil
}
