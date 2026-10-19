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
