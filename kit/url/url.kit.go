package urlpkg

import (
	"net/url"
	"reflect"
	"strings"
)

func Encode(raw string) string {
	return strings.Replace(raw, "+", "%20", -1)
}
func EncodeValues(values url.Values) string {
	return Encode(values.Encode())
}

// BuildRequestURL joins an endpoint and API path with one separator.
func BuildRequestURL(endpoint, apiPath string) string {
	if endpoint == "" {
		return apiPath
	}
	if apiPath == "" {
		return endpoint
	}
	return strings.TrimRight(endpoint, "/") + "/" + strings.TrimLeft(apiPath, "/")
}

// Deprecated: use BuildRequestURL instead.
func GenRequestURL(endpoint, apiPath string) string {
	return BuildRequestURL(endpoint, apiPath)
}

// QueryParamEncoder ...
type QueryParamEncoder interface {
	Encoder() url.Values
}

// AppendQueryParams merges encoded parameters into an existing request URL.
func AppendQueryParams(requestURL string, req QueryParamEncoder) string {
	if req == nil || isNilEncoder(req) {
		return requestURL
	}
	param := req.Encoder()
	if len(param) == 0 {
		return requestURL
	}
	parsed, err := url.Parse(requestURL)
	if err != nil {
		return requestURL
	}
	query := parsed.Query()
	for key, values := range param {
		for _, value := range values {
			query.Add(key, value)
		}
	}
	parsed.RawQuery = EncodeValues(query)
	return parsed.String()
}

func isNilEncoder(req QueryParamEncoder) bool {
	v := reflect.ValueOf(req)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}

// Deprecated: use AppendQueryParams instead.
func SplicingQueryParam(requestURL string, req QueryParamEncoder) string {
	return AppendQueryParams(requestURL, req)
}
