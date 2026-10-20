package apppkg

import (
	"io"
	stdhttp "net/http"
	"testing"

	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/v3/error"
	"github.com/stretchr/testify/require"
)

type countingReader struct {
	remaining int64
	read      int64
}

func (r *countingReader) Read(p []byte) (int, error) {
	if r.remaining == 0 {
		return 0, io.EOF
	}
	n := int64(len(p))
	if n > r.remaining {
		n = r.remaining
	}
	r.remaining -= n
	r.read += n
	return int(n), nil
}

func TestRequestDecoderLimitsBodyRead(t *testing.T) {
	RegisterCodec()
	body := &countingReader{remaining: DefaultUploadMaxSize * 2}
	req, err := stdhttp.NewRequest(stdhttp.MethodPost, "/", io.NopCloser(body))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	err = RequestDecoder(req, &Response{})
	require.True(t, errorpkg.IsCode(err, stdhttp.StatusRequestEntityTooLarge))
	require.LessOrEqual(t, body.read, int64(DefaultUploadMaxSize+1))
}
