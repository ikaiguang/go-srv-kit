package urlpkg

import (
	"net/url"
	"strings"

	bufferpkg "github.com/ikaiguang/go-srv-kit/kit/buffer"
)

func Encode(raw string) string {
	return strings.Replace(raw, "+", "%20", -1)
}
func EncodeValues(values url.Values) string {
	return Encode(values.Encode())
}

// GenRequestURL ...
func GenRequestURL(endpoint, apiPath string) string {
	buf := bufferpkg.GetBuffer()
	defer bufferpkg.PutBuffer(buf)

	buf.WriteString(endpoint)
	buf.WriteString(apiPath)
	return buf.String()
}

// QueryParamEncoder ...
type QueryParamEncoder interface {
	Encoder() url.Values
}

// SplicingQueryParam 输出例子：a=1&b=xxx
func SplicingQueryParam(requestURL string, req QueryParamEncoder) string {
	param := req.Encoder()
	if len(param) == 0 {
		return requestURL
	}
	paramString := param.Encode()
	paramString = strings.Replace(paramString, "+", "%20", -1)

	buf := bufferpkg.GetBuffer()
	defer bufferpkg.PutBuffer(buf)

	buf.WriteString(requestURL)
	buf.WriteString("?")
	buf.WriteString(paramString)
	return buf.String()
}
