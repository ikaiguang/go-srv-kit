package md5pkg

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// Md5 .
func Md5(bodyBytes []byte) (res string, err error) {
	handler := md5.New()
	_, err = handler.Write(bodyBytes)
	if err != nil {
		return res, err
	}
	res = hex.EncodeToString(handler.Sum(nil))
	return res, err
}

// FileMd5 获取文件的MD5
func FileMd5(path string) (string, error) {
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
