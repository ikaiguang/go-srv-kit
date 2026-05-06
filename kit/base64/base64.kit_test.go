package base64pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count 1 ./base64 -run TestBase64
func TestBase64(t *testing.T) {
	src := []byte("hello world")

	encoded := Encode(src)
	decoded, err := Decode(encoded)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(decoded))

	assert.Equal(t, src, decoded)
}
