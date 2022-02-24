package apputil

import (
	"testing"

	"github.com/stretchr/testify/require"

	v1 "github.com/ikaiguang/go-srv-kit/api/ping/v1"
	respv1 "github.com/ikaiguang/go-srv-kit/api/response/v1"
)

// go test -v -count=1 ./kratos/app -test.run=TestDecodeSuccessResponse
func TestDecodeSuccessResponse(t *testing.T) {
	body := []byte(`{"code":0,"reason":"","message":"","data":{"message":"Received Message : hello"},"request_id":""}`)
	tests := []struct {
		name     string
		body     []byte
		wantCode int
		wantData *v1.PingResp
	}{
		{
			name:     "#test_data",
			body:     body,
			wantCode: 0,
			wantData: &v1.PingResp{Message: "Received Message : hello"},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			actual := &v1.PingResp{}
			got, err := DecodeSuccessResponse(v.body, actual)
			require.Nil(t, err)
			require.Equal(t, v.wantCode, int(got.Code))
			require.Equal(t, v.wantData.Message, actual.Message)
		})
	}
}

// go test -v -count=1 ./kratos/app -test.run=TestDecodeErrorResponse
func TestDecodeErrorResponse(t *testing.T) {
	body := []byte(`{"code":400, "reason":"CONTENT_ERROR", "message":"testing error", "requestId":"", "data":null, "metadata":{"testdata":"testdata"}}`)
	tests := []struct {
		name     string
		body     []byte
		wantCode int
		wantData *respv1.Response
	}{
		{
			name:     "#test_error",
			body:     body,
			wantCode: 400,
			wantData: &respv1.Response{
				Code:    400,
				Reason:  "CONTENT_ERROR",
				Message: "testing error",
			},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got, err := DecodeErrorResponse(v.body)
			require.Nil(t, err)
			require.Equal(t, v.wantCode, int(got.Code))
			require.Equal(t, v.wantData.Reason, got.Reason)
			require.Equal(t, v.wantData.Message, got.Message)
		})
	}
}
