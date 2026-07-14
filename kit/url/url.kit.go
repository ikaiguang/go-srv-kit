package urlpkg

import (
	"net/url"
	"strings"

	"github.com/valyala/bytebufferpool"
)

func Encode(raw string) string {
	return strings.Replace(raw, "+", "%20", -1)
}
func EncodeValues(values url.Values) string {
	return Encode(values.Encode())
}

// GenRequestURL 拼接请求 URL
func GenRequestURL(endpoint, apiPath string) string {
	return endpoint + apiPath
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

	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)

	buf.WriteString(requestURL)
	buf.WriteString("?")
	buf.WriteString(paramString)
	return buf.String()
}
