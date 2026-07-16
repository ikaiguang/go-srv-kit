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

// SecureVerifyCode generates a numeric verification code using crypto/rand.
func SecureVerifyCode(length int) (string, error) {
	return SecureString(length, CharsetNumeral)
}

// SecurePassword generates a password containing uppercase, lowercase,
// numeric, and special characters using crypto/rand.
func SecurePassword(length int) (string, error) {
	if length < 8 {
		length = 8
	}
	password := make([]byte, length)
	requiredCharsets := []string{CharsetUppercase, CharsetLowercase, CharsetNumeral, CharsetSpecial}
	for i, charset := range requiredCharsets {
		value, err := SecureString(1, charset)
		if err != nil {
			return "", err
		}
		password[i] = value[0]
	}
	rest, err := SecureString(length-len(requiredCharsets), CharsetPassword)
	if err != nil {
		return "", err
	}
	copy(password[len(requiredCharsets):], rest)
	if err := secureShuffle(password); err != nil {
		return "", err
	}
	return string(password), nil
}

func secureShuffle(data []byte) error {
	for i := len(data) - 1; i > 0; i-- {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return err
		}
		j := int(index.Int64())
		data[i], data[j] = data[j], data[i]
	}
	return nil
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
