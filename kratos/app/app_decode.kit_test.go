package apputil

import (
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/anypb"
	commonv1 "realibox-global/go-global-services/api/common/v1"
	pingv1 "realibox-global/go-global-services/api/ping-service/v1/resources"
	"testing"
)

// go test -v -count=1 ./business/app -test.run=TestDecodeProto
func TestDecodeProto(t *testing.T) {
	var (
		err      error
		response = &commonv1.Response{
			Code:      400,
			Reason:    "CONTENT_ERROR",
			Message:   "testing error",
			RequestId: "",
			Metadata:  map[string]string{"testdata": "testdata"},
		}
		data = &pingv1.PingResp{
			Message: "ping-pong",
		}
	)

	response.Data, err = anypb.New(data)
	require.Nil(t, err)
	body, err := protojson.Marshal(response)
	require.Nil(t, err)

	tests := []struct {
		name     string
		body     []byte
		wantCode int32
		wantData *commonv1.Response
	}{
		{
			name:     "#test_error",
			body:     body,
			wantCode: response.Code,
			wantData: response,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			actual := &pingv1.PingResp{}
			got, err := DecodeProto(v.body, actual)
			require.Nil(t, err)
			require.Equal(t, v.wantCode, got.Code)
			require.Equal(t, v.wantData.Reason, got.Reason)
			require.Equal(t, data.Message, actual.Message)
		})
	}
}

// go test -v -count=1 ./business/app -test.run=TestDecodeResponse
func TestDecodeResponse(t *testing.T) {
	body := []byte(`{"code":0,"reason":"","message":"","data":{"message":"Received Message : hello"},"request_id":""}`)
	tests := []struct {
		name     string
		body     []byte
		wantCode int
		wantData *pingv1.PingResp
	}{
		{
			name:     "#test_data",
			body:     body,
			wantCode: 0,
			wantData: &pingv1.PingResp{Message: "Received Message : hello"},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			actual := &pingv1.PingResp{}
			got, err := DecodeResponse(v.body, actual)
			require.Nil(t, err)
			require.Equal(t, v.wantCode, int(got.Code))
			require.Equal(t, v.wantData.Message, actual.Message)
		})
	}
}

// go test -v -count=1 ./business/app -test.run=TestDecodeError
func TestDecodeError(t *testing.T) {
	body := []byte(`{"code":400, "reason":"CONTENT_ERROR", "message":"testing error", "requestId":"", "data":null, "metadata":{"testdata":"testdata"}}`)
	tests := []struct {
		name     string
		body     []byte
		wantCode int
		wantData *commonv1.Response
	}{
		{
			name:     "#test_error",
			body:     body,
			wantCode: 400,
			wantData: &commonv1.Response{
				Code:    400,
				Reason:  "CONTENT_ERROR",
				Message: "testing error",
			},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got, err := DecodeError(v.body)
			require.Nil(t, err)
			require.Equal(t, v.wantCode, int(got.Code))
			require.Equal(t, v.wantData.Reason, got.Reason)
			require.Equal(t, v.wantData.Message, got.Message)
		})
	}
}
