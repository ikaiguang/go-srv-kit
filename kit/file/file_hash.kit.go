package filepkg

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"strconv"
)

// Deprecated: 函数名未明确指示使用的哈希算法（实际使用 MD5）。请使用 MD5 替代。
func Hash(filePath string) (string, int64, error) {
	return MD5(filePath)
}

// Deprecated: 函数名未明确指示使用的哈希算法（实际使用 MD5）。请使用 MD5FromReader 替代。
func HashFromFile(f io.Reader) (string, int64, error) {
	return MD5FromReader(f)
}

// MD5 returns the MD5 digest and size of a file.
func MD5(filePath string) (string, int64, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", 0, err
	}
	defer func() { _ = f.Close() }()

	return MD5FromReader(f)
}

// MD5FromReader returns the MD5 digest and number of bytes read.
func MD5FromReader(f io.Reader) (string, int64, error) {
	h := md5.New()
	size, err := io.Copy(h, f)
	if err != nil {
		return "", size, err
	}
	return hex.EncodeToString(h.Sum(nil)), size, nil
}

// SHA256FromReader returns the SHA-256 digest and number of bytes read.
func SHA256FromReader(f io.Reader) (string, int64, error) {
	h := sha256.New()
	size, err := io.Copy(h, f)
	if err != nil {
		return "", size, err
	}
	return hex.EncodeToString(h.Sum(nil)), size, nil
}

// SHA256 returns the SHA-256 digest and size of a file.
func SHA256(filePath string) (string, int64, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", 0, err
	}
	defer func() { _ = f.Close() }()

	return SHA256FromReader(f)
}

// Deprecated: use MD5 instead.
func Md5(filePath string) (string, int64, error) { return MD5(filePath) }

// Deprecated: use MD5FromReader instead.
func Md5FromFile(f io.Reader) (string, int64, error) { return MD5FromReader(f) }

// Deprecated: use SHA256FromReader instead.
func Sha256FromFile(f io.Reader) (string, int64, error) { return SHA256FromReader(f) }

// Deprecated: use SHA256 instead.
func Sha256(filePath string) (string, int64, error) { return SHA256(filePath) }

// Identifier return hash + "-" + size
// md5(32)+size(19)=51; sha256(64)碰撞概率非常低
// 可结合mime类型来提高可靠性，mime读取可参考： https://github.com/h2non/filetype; pass the file header = first 261 bytes
func Identifier(hash string, size int64) string {
	return hash + "-" + strconv.FormatInt(size, 10)
}

// Deprecated: 请使用 MD5() + Identifier() 组合替代。
func HashIdentifier(filePath string) (string, int64, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", 0, err
	}
	defer func() { _ = f.Close() }()

	return HashIdentifierFromFile(f)
}

// HashIdentifierFromFile ...
// Deprecated: 请使用 MD5FromReader() + Identifier() 组合替代。
func HashIdentifierFromFile(f io.Reader) (string, int64, error) {
	hash, size, err := MD5FromReader(f)
	if err != nil {
		return "", size, err
	}
	return Identifier(hash, size), size, nil
}
