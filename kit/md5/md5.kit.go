package md5pkg

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// MD5 returns the MD5 digest of bodyBytes as lowercase hexadecimal.
func MD5(bodyBytes []byte) (res string, err error) {
	handler := md5.New()
	_, err = handler.Write(bodyBytes)
	if err != nil {
		return res, err
	}
	res = hex.EncodeToString(handler.Sum(nil))
	return res, err
}

// FileMD5 returns the MD5 digest of a file as lowercase hexadecimal.
func FileMD5(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func() { _ = f.Close() }()

	h := md5.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

// Deprecated: use MD5 instead.
func Md5(bodyBytes []byte) (string, error) { return MD5(bodyBytes) }

// Deprecated: use FileMD5 instead.
func FileMd5(path string) (string, error) { return FileMD5(path) }
