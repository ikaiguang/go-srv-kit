package aespkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAesCipher(t *testing.T) {
	key := []byte("1234567890ABCDEF")
	cipher, err := NewAesCipher(key)
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

func BenchmarkAes(b *testing.B) {
	aesCipher, err := NewAesCipher([]byte("ur38ifsewn8b49i9"))
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
