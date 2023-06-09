package md5pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v ./pkg/crypto -count=1 -test.run=TestMd5
func TestMd5(t *testing.T) {
	tests := []struct {
		name      string
		plaintext string
		//ciphertext string
		want string
	}{
		{name: "加密：123456", plaintext: "123456", want: "e10adc3949ba59abbe56e057f20f883e"},
		{name: "加密：abcdef", plaintext: "abcdef", want: "e80b5017098950fc58aad83c8c14978e"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Md5([]byte(tt.plaintext))
			assert.Nil(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
