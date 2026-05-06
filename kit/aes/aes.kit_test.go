package aespkg

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncryptCBCAndDecryptCBC(t *testing.T) {
	key := []byte("1234567890ABCDEF")
	plaintext := []byte("hello cbc")

	encrypted, err := EncryptCBC(plaintext, key)
	require.NoError(t, err)
	assert.NotEmpty(t, encrypted)

	got, err := DecryptCBC(encrypted, key)
	require.NoError(t, err)
	assert.Equal(t, string(plaintext), got)
}

func TestDecryptCBCInvalidPadding(t *testing.T) {
	key := []byte("1234567890ABCDEF")
	raw := make([]byte, 32)
	ciphertext := base64.URLEncoding.EncodeToString(raw)

	got, err := DecryptCBC(ciphertext, key)
	require.Error(t, err)
	assert.Empty(t, got)
}
