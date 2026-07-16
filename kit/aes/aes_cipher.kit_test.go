package aespkg

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAesCipher(t *testing.T) {
	key := []byte("1234567890ABCDEF")
	cipher, err := NewAESCipher(key)
	if err != nil {
		t.Fatal(err)
	}

	msg := "hello world"
	encrypted, err := cipher.EncryptToString(msg)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(encrypted)

	decrypted, err := cipher.DecryptToString(encrypted)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(decrypted)

	assert.Equal(t, msg, decrypted)
}

func TestAesCipherCopiesKeyMaterial(t *testing.T) {
	key := []byte("1234567890ABCDEF")
	cipher, err := NewAESCipher(key)
	if err != nil {
		t.Fatal(err)
	}

	before, err := cipher.EncryptToString("message")
	if err != nil {
		t.Fatal(err)
	}
	for i := range key {
		key[i] = 'X'
	}
	after, err := cipher.EncryptToString("message")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, before, after)
}

func TestAesCipherDecryptInvalidCiphertext(t *testing.T) {
	cipher, err := NewAESCipher([]byte("1234567890ABCDEF"))
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name       string
		ciphertext string
	}{
		{name: "空密文", ciphertext: ""},
		{name: "非block倍数", ciphertext: base64.StdEncoding.EncodeToString([]byte("short"))},
		{name: "非法padding", ciphertext: base64.StdEncoding.EncodeToString(make([]byte, 16))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cipher.DecryptToString(tt.ciphertext)
			assert.Error(t, err)
			assert.Empty(t, got)
		})
	}
}

func BenchmarkAes(b *testing.B) {
	aesCipher, err := NewAESCipher([]byte("ur38ifsewn8b49i9"))
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encrypt, err := aesCipher.EncryptToString("hello")
		if err != nil {
			b.Fatal(err)
		}
		_, err = aesCipher.DecryptToString(encrypt)
		if err != nil {
			b.Fatal(err)
		}
	}
}
