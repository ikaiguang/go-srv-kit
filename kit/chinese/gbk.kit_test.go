package chinesepkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// go test -v -count 1 ./chinese -run TestGbkToUtf8AndUtf8ToGbk
func TestGbkToUtf8AndUtf8ToGbk(t *testing.T) {
	src := []byte("中文测试")

	gbkBytes, err := Utf8ToGbk(src)
	require.NoError(t, err)
	assert.False(t, IsUtf8(string(gbkBytes)))
	assert.True(t, IsGBK(string(gbkBytes)))

	got, err := GbkToUtf8(gbkBytes)
	require.NoError(t, err)
	assert.Equal(t, src, got)
	assert.True(t, IsUtf8(string(got)))
}

// go test -v -count 1 ./chinese -run TestIsGBK
func TestIsGBK(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want bool
	}{
		{name: "utf8字符串不是gbk", data: []byte("hello"), want: false},
		{name: "合法gbk双字节", data: []byte{0xd6, 0xd0}, want: true},
		{name: "截断gbk不会panic", data: []byte{0xd6}, want: false},
		{name: "非法第二字节", data: []byte{0xd6, 0x20}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsGBK(string(tt.data)))
		})
	}
}
