package rsapkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 私钥生成
// openssl genrsa -out rsa_private_key.pem 1024
var privateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQCmgQXSIeJuR+Y0TJb9TYBMRIjeyPZF98/ZJHwqUZ/nDct1nQV/
YYmfe4jtVKJkORj+wBjg/QEfgxpa1YxMlDqBWOt987yVHzK4Wnc9bWUcb34YKAnU
WS70GyoPBCZt7WjVKaGpU7aOZDw8zEojD7vhi4q/A5JwkWJ4ZfSs6F2tfwIDAQAB
AoGACEO9OzntWFX/SjdHA1m2dZKtTImjF8P+MCQMeblFe52GrNbXcAQyZZUnLciW
quzenb6BPaGxTZQfWcThyudMpAx0cFCDWh/uETmCz/1RJwaZmgXP/lq2BspaF00v
7xp5MLIBJePaU6ulww0ofbkjoZpSwXThD0CwKSzRzEZ5vXECQQDZxzRja4BQRx02
yqcGuJbEvJEE+EiYsxpF8lgs1RA93FZtVxM3Xe+HzBJrRRgLkGkTZTlSbwSPDCxo
o1EEmGBJAkEAw7oQq35IGKTYeiinExZDbIlKHVKJQcCXnohnzszRapZlxfu9ncj1
RZ8Xp83Or2K8Y1erVnB0m16SHBcdsqmvhwJAd72B3pBDEuCm/XNbduSTcUTE79ic
AemoLoFbXfsgXQMDOkdAN5cclqvsDLMGz4TtYU6sv9hux0BIQphZeY9WkQJAOAnx
8+f4JHYuNOumymQ5cb3tJnAXNGg8APv1HNSvsODWytTE+YQsFX7zeuwGHVkbryXO
vLT97e4pzzkfG6RRyQJAfB9VxUEk/skZpkZfb+LFj908RgiltYLBuqmT6fQkVQqG
D4yTBgvSK5dvgE2HocnNjXW1Eg7dmfQhUE/F490Xsg==
-----END RSA PRIVATE KEY-----`)

// 公钥: 根据私钥生成
// openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
var publicKey = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCmgQXSIeJuR+Y0TJb9TYBMRIje
yPZF98/ZJHwqUZ/nDct1nQV/YYmfe4jtVKJkORj+wBjg/QEfgxpa1YxMlDqBWOt9
87yVHzK4Wnc9bWUcb34YKAnUWS70GyoPBCZt7WjVKaGpU7aOZDw8zEojD7vhi4q/
A5JwkWJ4ZfSs6F2tfwIDAQAB
-----END PUBLIC KEY-----`)

func TestRsa_Encrypt(t *testing.T) {
	r, err := NewRsaCipher(publicKey, privateKey)
	if err != nil {
		t.Fatal(err)
	}

	msg := "hello"
	encrypt, err := r.EncryptToString(msg)
	if err != nil {
		t.Fatal(err)
	}

	decrypt, err := r.DecryptToString(encrypt)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, msg, decrypt)
}

func TestRsaCipher_Sign(t *testing.T) {
	r, err := NewRsaCipher(publicKey, privateKey)
	if err != nil {
		t.Fatal(err)
	}

	msg := []byte("hello")
	sign, err := r.Sign(msg)
	if err != nil {
		t.Fatal(err)
	}

	valid, err := r.VerifySign(msg, sign)
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, valid)
}

func TestRsa(t *testing.T) {
	// rsa 密钥文件产生
	priKey, pubKey, err := GenRsaKey()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(priKey))
	t.Log(string(pubKey))

	r, err := NewRsaCipher(pubKey, priKey)
	if err != nil {
		t.Fatal(err)
	}

	msg := "hello"
	encrypt, err := r.EncryptToString(msg)
	if err != nil {
		t.Fatal(err)
	}

	decrypt, err := r.DecryptToString(encrypt)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, msg, decrypt)

	msgBytes := []byte(msg)
	sign, err := r.Sign(msgBytes)
	if err != nil {
		t.Fatal(err)
	}

	valid, err := r.VerifySign(msgBytes, sign)
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, valid)
}

func BenchmarkRasEncrypt(b *testing.B) {
	cipher, err := NewRsaCipher(publicKey, privateKey)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encrypt, err := cipher.EncryptToString("hello")
		if err != nil {
			b.Fatal(err)
		}
		_, err = cipher.DecryptToString(encrypt)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRasSign(b *testing.B) {
	cipher, err := NewRsaCipher(publicKey, privateKey)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sign, err := cipher.Sign([]byte("hello"))
		if err != nil {
			b.Fatal(err)
		}
		valid, err := cipher.VerifySign([]byte("hello"), sign)
		if err != nil {
			b.Fatal(err)
		}
		if !valid {
			b.Fatal(err)
		}
	}
}
