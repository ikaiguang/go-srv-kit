package apppkg

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/anypb"
)

// go test -v -count=1 ./kratos/app -run TestCustomDecodeProtobufResponse
func TestCustomDecodeProtobufResponse(t *testing.T) {
	var (
		err      error
		response = &Response{
			Code:     400,
			Reason:   "CONTENT_ERROR",
			Message:  "testing error",
			Metadata: map[string]string{"testdata": "testdata"},
		}
		data = &Response{
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
		wantData *Response
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
			actual := &Response{}
			got, err := CustomDecodeProtobufResponse(v.body, actual)
			require.Nil(t, err)
			require.Equal(t, v.wantCode, got.Code)
			require.Equal(t, v.wantData.Reason, got.Reason)
			require.Equal(t, data.Message, actual.Message)
		})
	}
}

// go test -v -count=1 ./business/app -run TestCustomDecodeHTTPResponse
func TestCustomDecodeHTTPResponse(t *testing.T) {
	body := []byte(`{"code":0,"reason":"","message":"","data":{"message":"Received Message : hello"},"request_id":""}`)
	tests := []struct {
		name     string
		body     []byte
		wantCode int
		wantData *Response
	}{
		{
			name:     "#test_data",
			body:     body,
			wantCode: 0,
			wantData: &Response{Message: "Received Message : hello"},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			actual := &Response{}
			got, err := CustomDecodeHTTPResponse(v.body, actual)
			require.Nil(t, err)
			require.Equal(t, v.wantCode, int(got.Code))
			require.Equal(t, v.wantData.Message, actual.Message)
		})
	}
}

// go test -v -count=1 ./business/app -run TestCustomDecodeResponseError
func TestCustomDecodeResponseError(t *testing.T) {
	body := []byte(`{"code":400, "reason":"CONTENT_ERROR", "message":"testing error", "requestId":"", "data":null, "metadata":{"testdata":"testdata"}}`)
	tests := []struct {
		name     string
		body     []byte
		wantCode int
		wantData *Response
	}{
		{
			name:     "#test_error",
			body:     body,
			wantCode: 400,
			wantData: &Response{
				Code:    400,
				Reason:  "CONTENT_ERROR",
				Message: "testing error",
			},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got, err := CustomDecodeResponseError(v.body)
			require.Nil(t, err)
			require.Equal(t, v.wantCode, int(got.Code))
			require.Equal(t, v.wantData.Reason, got.Reason)
			require.Equal(t, v.wantData.Message, got.Message)
		})
	}
}
