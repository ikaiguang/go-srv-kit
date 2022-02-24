package apputil

import (
	"github.com/go-kratos/kratos/v2/transport/http"
)

const (
	baseContentType = "application"
)

var (
	_ = http.DefaultRequestDecoder
	_ = http.DefaultResponseEncoder
	_ = http.DefaultErrorEncoder
)

// Response 响应
// 关联更新 v1.Response
type Response struct {
	Code     int32             `json:"code"`
	Reason   string            `json:"reason"`
	Message  string            `json:"message"`
	Metadata map[string]string `json:"metadata,omitempty"`

	Data      interface{} `json:"data"`
	RequestId string      `json:"request_id"`
}
