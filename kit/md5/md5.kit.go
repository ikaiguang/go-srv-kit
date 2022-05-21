package md5util

import (
	"crypto/md5"
	"encoding/hex"
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
