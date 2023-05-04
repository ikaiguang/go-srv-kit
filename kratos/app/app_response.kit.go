package apputil

// ResponseInterface .
type ResponseInterface interface {
	GetCode() int32
	GetReason() string
	GetMessage() string
	GetRequestId() string
	GetMetadata() map[string]string
}

// Response 响应
// 关联更新 responsev1.Response
type Response struct {
	Code     int32             `json:"code"`
	Reason   string            `json:"reason"`
	Message  string            `json:"message"`
	Metadata map[string]string `json:"metadata,omitempty"`

	Data      interface{} `json:"data"`
	RequestId string      `json:"request_id"`
}

func (x *Response) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Response) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *Response) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}
