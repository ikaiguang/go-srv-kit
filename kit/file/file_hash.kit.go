package filepkg

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"strconv"
)

func Hash(filePath string) (string, int64, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", 0, err
	}
	defer func() { _ = f.Close() }()

	return HashFromFile(f)
}

func HashFromFile(f io.Reader) (string, int64, error) {
	hash := md5.New()
	buf := make([]byte, 1<<20)
	var size int64 = 0
	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			return "", 0, err
		}
		if n == 0 {
			break
		}
		hash.Write(buf[:n])
		size += int64(n)
	}
	return hex.EncodeToString(hash.Sum(nil)), size, nil
}

// Md5 return md5, size, err
func Md5(filePath string) (string, int64, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", 0, err
	}
	defer func() { _ = f.Close() }()

	return Md5FromFile(f)
}

// Md5FromFile return md5, size, err
func Md5FromFile(f io.Reader) (string, int64, error) {
	h := md5.New()
	size, err := io.Copy(h, f)
	if err != nil {
		return "", size, err
	}
	return hex.EncodeToString(h.Sum(nil)), size, nil
}

// Sha256FromFile return hash, size, err
func Sha256FromFile(f io.Reader) (string, int64, error) {
	h := sha256.New()
	size, err := io.Copy(h, f)
	if err != nil {
		return "", size, err
	}
	return hex.EncodeToString(h.Sum(nil)), size, nil
}

// Sha256 return hash, size, err
func Sha256(filePath string) (string, int64, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", 0, err
	}
	defer func() { _ = f.Close() }()

	return Sha256FromFile(f)
}

// Identifier return hash + "-" + size
// md5(32)+size(19)=51; sha256(64)碰撞概率非常低
// 可结合mime类型来提高可靠性，mime读取可参考： https://github.com/h2non/filetype; pass the file header = first 261 bytes
func Identifier(hash string, size int64) string {
	return hash + "-" + strconv.FormatInt(size, 10)
}

// HashIdentifier 文件标识符; return identifier, size, err
// identifier = hash + "-" + size;
func HashIdentifier(filePath string) (string, int64, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", 0, err
	}
	defer func() { _ = f.Close() }()

	return HashIdentifierFromFile(f)
}

// HashIdentifierFromFile 文件标识符; return identifier, size, err
// identifier = hash + "-" + size
func HashIdentifierFromFile(f io.Reader) (string, int64, error) {
	hash, size, err := HashFromFile(f)
	if err != nil {
		return "", size, err
	}
	return Identifier(hash, size), size, nil
}
