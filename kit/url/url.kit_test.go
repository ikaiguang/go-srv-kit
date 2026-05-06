package urlpkg

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count 1 ./url -run TestEncode
func TestEncode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"替换加号", "hello+world", "hello%20world"},
		{"无加号", "hello", "hello"},
		{"多个加号", "a+b+c", "a%20b%20c"},
		{"空字符串", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Encode(tt.input))
		})
	}
}

// go test -v -count 1 ./url -run TestEncodeValues
func TestEncodeValues(t *testing.T) {
	values := url.Values{}
	values.Set("name", "hello world")
	values.Set("age", "20")

	result := EncodeValues(values)
	assert.NotEmpty(t, result)
	assert.NotContains(t, result, "+", "EncodeValues 应将 + 替换为 %%20")
	assert.Contains(t, result, "name=hello%20world")
}

// go test -v -count 1 ./url -run TestGenRequestURL
func TestGenRequestURL(t *testing.T) {
	tests := []struct {
		name     string
		endpoint string
		apiPath  string
		want     string
	}{
		{"正常拼接", "https://api.example.com", "/v1/users", "https://api.example.com/v1/users"},
		{"空路径", "https://api.example.com", "", "https://api.example.com"},
		{"空端点", "", "/v1/users", "/v1/users"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, GenRequestURL(tt.endpoint, tt.apiPath))
		})
	}
}

// mockQueryParam 实现 QueryParamEncoder 接口
type mockQueryParam struct {
	values url.Values
}

func (m *mockQueryParam) Encoder() url.Values {
	return m.values
}

// go test -v -count 1 ./url -run TestSplicingQueryParam
func TestSplicingQueryParam(t *testing.T) {
	t.Run("有参数", func(t *testing.T) {
		values := url.Values{}
		values.Set("page", "1")
		values.Set("size", "10")
		req := &mockQueryParam{values: values}

		result := SplicingQueryParam("https://api.example.com/users", req)
		assert.Contains(t, result, "https://api.example.com/users?")
		assert.Contains(t, result, "page=1")
		assert.Contains(t, result, "size=10")
	})

	t.Run("无参数", func(t *testing.T) {
		req := &mockQueryParam{values: url.Values{}}
		result := SplicingQueryParam("https://api.example.com/users", req)
		assert.Equal(t, "https://api.example.com/users", result)
	})
}
