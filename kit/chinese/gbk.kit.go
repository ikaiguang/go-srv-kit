package chineseutil

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"unicode/utf8"
)

var (
// gbkEncoder 编码
//gbk18030Encoder = simplifiedchinese.GB18030.NewEncoder()
//gbkEncoder      = simplifiedchinese.GBK.NewEncoder()
//gbk2312Encoder  = simplifiedchinese.HZGB2312.NewEncoder()

// gbk18030Decoder 解码
//gbk18030Decoder = simplifiedchinese.GB18030.NewDecoder()
//gbkDecoder      = simplifiedchinese.GBK.NewDecoder()
//gbk2312Decoder  = simplifiedchinese.HZGB2312.NewDecoder()
)

// GbkToUtf8 ...
func GbkToUtf8(gbkByte []byte) (res []byte, err error) {
	if res, err = simplifiedchinese.GB18030.NewDecoder().Bytes(gbkByte); err == nil {
		return res, err
	}
	if res, err = simplifiedchinese.GBK.NewDecoder().Bytes(gbkByte); err == nil {
		return res, err
	}
	return simplifiedchinese.HZGB2312.NewDecoder().Bytes(gbkByte)
}

// Utf8ToGbk ...
func Utf8ToGbk(utf8Byte []byte) (res []byte, err error) {
	if res, err = simplifiedchinese.GB18030.NewEncoder().Bytes(utf8Byte); err == nil {
		return res, err
	}
	if res, err = simplifiedchinese.GBK.NewEncoder().Bytes(utf8Byte); err == nil {
		return res, err
	}
	return simplifiedchinese.HZGB2312.NewEncoder().Bytes(utf8Byte)
}

// IsUtf8 是否utf8
func IsUtf8(s string) bool {
	return utf8.ValidString(s)
}

// IsGBK 是否gbk
func IsGBK(s string) bool {
	if IsUtf8(s) {
		return false
	}
	data := []byte(s)
	length := len(data)
	var i int = 0
	for i < length {
		//fmt.Printf("for %x\n", data[i])
		if data[i] <= 0xff {
			i++
			continue
		} else {
			if data[i] >= 0x81 &&
				data[i] <= 0xfe &&
				data[i+1] >= 0x40 &&
				data[i+1] <= 0xfe &&
				data[i+1] != 0xf7 {
				i += 2
				continue
			} else {
				return false
			}
		}
	}
	return true
}
