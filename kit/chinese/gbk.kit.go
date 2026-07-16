package chinesepkg

import (
	"unicode/utf8"

	"golang.org/x/text/encoding/simplifiedchinese"
)

var (
// gbkEncoder 编码
//gbk18030Encoder = simplifiedchinese.GB18030.NewEncoder()
//gbkEncoder      = simplifiedchinese.GBK.NewEncoder()
//gbk2312Encoder  = simplifiedchinese.HZGB2312.NewEncoder()

// gbk18030Decoder 解码
// gbk18030Decoder = simplifiedchinese.GB18030.NewDecoder()
// gbkDecoder      = simplifiedchinese.GBK.NewDecoder()
// gbk2312Decoder  = simplifiedchinese.HZGB2312.NewDecoder()
)

// GBKToUTF8 converts GB18030, GBK, or HZ-GB2312 bytes to UTF-8.
func GBKToUTF8(gbkByte []byte) (res []byte, err error) {
	if res, err = simplifiedchinese.GB18030.NewDecoder().Bytes(gbkByte); err == nil {
		return res, err
	}
	if res, err = simplifiedchinese.GBK.NewDecoder().Bytes(gbkByte); err == nil {
		return res, err
	}
	return simplifiedchinese.HZGB2312.NewDecoder().Bytes(gbkByte)
}

// UTF8ToGBK converts UTF-8 bytes to a compatible Chinese encoding.
func UTF8ToGBK(utf8Byte []byte) (res []byte, err error) {
	if res, err = simplifiedchinese.GB18030.NewEncoder().Bytes(utf8Byte); err == nil {
		return res, err
	}
	if res, err = simplifiedchinese.GBK.NewEncoder().Bytes(utf8Byte); err == nil {
		return res, err
	}
	return simplifiedchinese.HZGB2312.NewEncoder().Bytes(utf8Byte)
}

// IsUTF8 reports whether s contains valid UTF-8.
func IsUTF8(s string) bool {
	return utf8.ValidString(s)
}

// Deprecated: use GBKToUTF8 instead.
func GbkToUtf8(gbkByte []byte) ([]byte, error) { return GBKToUTF8(gbkByte) }

// Deprecated: use UTF8ToGBK instead.
func Utf8ToGbk(utf8Byte []byte) ([]byte, error) { return UTF8ToGBK(utf8Byte) }

// Deprecated: use IsUTF8 instead.
func IsUtf8(s string) bool { return IsUTF8(s) }

// IsGBK 是否gbk
func IsGBK(s string) bool {
	if IsUTF8(s) {
		return false
	}
	data := []byte(s)
	length := len(data)
	var i int = 0
	for i < length {
		//fmt.Printf("for %x\n", data[i])
		if data[i] <= 0x7f {
			i++
			continue
		}
		if i+1 >= length {
			return false
		}
		if data[i] >= 0x81 &&
			data[i] <= 0xfe &&
			data[i+1] >= 0x40 &&
			data[i+1] <= 0xfe &&
			data[i+1] != 0xf7 {
			i += 2
			continue
		}
		return false
	}
	return true
}
