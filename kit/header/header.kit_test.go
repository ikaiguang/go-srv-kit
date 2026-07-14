package headerpkg

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsValue(t *testing.T) {
	tests := []struct {
		name   string
		header http.Header
		key    string
		value  string
		want   bool
	}{
		{
			name:   "精确匹配",
			header: http.Header{"Upgrade": {"websocket"}},
			key:    "Upgrade",
			value:  "websocket",
			want:   true,
		},
		{
			name:   "大小写不敏感",
			header: http.Header{"Upgrade": {"WebSocket"}},
			key:    "Upgrade",
			value:  "websocket",
			want:   true,
		},
		{
			name:   "逗号分隔列表",
			header: http.Header{"Connection": {"keep-alive, Upgrade"}},
			key:    "Connection",
			value:  "Upgrade",
			want:   true,
		},
		{
			name:   "不存在的值",
			header: http.Header{"Upgrade": {"websocket"}},
			key:    "Upgrade",
			value:  "http2",
			want:   false,
		},
		{
			name:   "不存在的key",
			header: http.Header{},
			key:    "Upgrade",
			value:  "websocket",
			want:   false,
		},
		{
			name:   "空header",
			header: http.Header{},
			key:    "Any",
			value:  "value",
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsValue(tt.header, tt.key, tt.value)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestConstants(t *testing.T) {
	// 验证常量值不为空
	assert.NotEmpty(t, RequestID)
	assert.NotEmpty(t, TraceID)
	assert.NotEmpty(t, InternalToken)
	assert.NotEmpty(t, AuthorizationKey)
	assert.NotEmpty(t, ContentType)
	assert.NotEmpty(t, ContentTypeJSON)
	assert.NotEmpty(t, ContentTypeJSONUtf8)
	assert.NotEmpty(t, ContentTypeProto)
	assert.NotEmpty(t, ContentTypeProtobuf)
	assert.NotEmpty(t, ContentTypeFormURLEncoded)
	assert.NotEmpty(t, ContentTypeMultipartForm)
}

func TestCommonHeaderHelpers(t *testing.T) {
	header := http.Header{}

	SetContentType(header, ContentTypeJSON)
	assert.Equal(t, ContentTypeJSON, GetContentType(header))

	SetAuthorization(header, "Bearer token")
	assert.Equal(t, "Bearer token", GetAuthorization(header))

	SetTraceID(header, "trace-id")
	assert.Equal(t, "trace-id", GetTraceID(header))

	SetRequestID(header, "request-id")
	assert.Equal(t, "request-id", GetRequestID(header))

	SetIsWebsocket(header)
	assert.True(t, GetIsWebsocket(header))
}
