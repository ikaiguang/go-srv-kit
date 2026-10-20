package apppkg

import (
	"testing"

	"github.com/go-kratos/kratos/v3/encoding"
	"github.com/stretchr/testify/require"
)

func TestRegisterCodecPreservesProtoJSONBehavior(t *testing.T) {
	RegisterCodec()
	codec := encoding.GetCodec("json")
	require.NotNil(t, codec)

	data, err := codec.Marshal(&Response{RequestId: "request-1"})
	require.NoError(t, err)
	require.Contains(t, string(data), `"request_id":"request-1"`)

	var response Response
	require.NoError(t, codec.Unmarshal(data, &response))
	require.Equal(t, "request-1", response.RequestId)
}

func TestUnmarshalJSONRejectsNilTargetsWithoutPanicking(t *testing.T) {
	require.NotPanics(t, func() {
		require.Error(t, UnmarshalJSON([]byte(`{}`), nil))
	})

	var response *Response
	require.NotPanics(t, func() {
		require.Error(t, UnmarshalJSON([]byte(`{}`), response))
	})

	var decoded map[string]any
	require.NoError(t, UnmarshalJSON([]byte(`{"ok":true}`), &decoded))
	require.Equal(t, true, decoded["ok"])

	var responsePointer *Response
	require.NoError(t, UnmarshalJSON([]byte(`{"request_id":"request-2"}`), &responsePointer))
	require.Equal(t, "request-2", responsePointer.RequestId)

	require.NotPanics(t, func() {
		SetJSONMarshalOptions(nil)
		SetJSONUnmarshalOptions(nil)
	})
}
