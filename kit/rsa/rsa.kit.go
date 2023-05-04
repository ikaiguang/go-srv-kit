package rsautil

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"

	base64util "github.com/ikaiguang/go-srv-kit/kit/base64"
)

// Encryptor ...
type Encryptor interface {
	EncryptToString(plaintext string) (string, error)
	DecryptToString(ciphertext string) (string, error)
}

// GenRsaKey RSA公钥私钥产生
func GenRsaKey() ([]byte, []byte, error) {
	// 私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return nil, nil, err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	priKeyBytes := pem.EncodeToMemory(block)

	// 公钥
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, nil, err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}

	pubKeyBytes := pem.EncodeToMemory(block)
	return priKeyBytes, pubKeyBytes, nil
}

// ParserPublicKey 解析公钥
func ParserPublicKey(key []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(key)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("failed to decode PEM block containing public key")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("asset public key fail")
	}
	return pub, nil
}

// ParserPrivateKey 解析私钥
func ParserPrivateKey(key []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(key)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("failed to decode PEM block containing private key")
	}
	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pri, nil
}

// RsaCipher rsa加解密
type RsaCipher struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

// NewRsaCipher rsa加解密
func NewRsaCipher(pubKey, priKey []byte) (*RsaCipher, error) {
	r := &RsaCipher{}
	err := r.parsePriKey(priKey)
	if err != nil {
		return nil, err
	}
	err = r.parsePubKey(pubKey)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// NewRsaCipherBase64 rsa加解密
func NewRsaCipherBase64(pubKeyBase64, priKeyBase64 []byte) (*RsaCipher, error) {
	pubKey, err := base64util.Decode(pubKeyBase64)
	if err != nil {
		return nil, err
	}
	priKey, err := base64util.Decode(priKeyBase64)
	if err != nil {
		return nil, err
	}
	return NewRsaCipher(pubKey, priKey)
}

func (r *RsaCipher) parsePubKey(pubKey []byte) error {
	publicKey, err := ParserPublicKey(pubKey)
	if err != nil {
		return err
	}
	r.publicKey = publicKey
	return nil
}

func (r *RsaCipher) parsePriKey(priKey []byte) error {
	privateKey, err := ParserPrivateKey(priKey)
	if err != nil {
		return err
	}
	r.privateKey = privateKey
	return nil
}

func (r *RsaCipher) Encrypt(plainText []byte) ([]byte, error) {
	encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, r.publicKey, plainText)
	if err != nil {
		return nil, err
	}
	return base64util.Encode(encrypted), nil
}

func (r *RsaCipher) Decrypt(cipherText []byte) ([]byte, error) {
	s, err := base64util.Decode(cipherText)
	if err != nil {
		return nil, err
	}
	decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, r.privateKey, s)
	if err != nil {
		return nil, err
	}
	return decrypted, nil
}

func (r *RsaCipher) EncryptToString(plainText string) (string, error) {
	encrypted, err := r.Encrypt([]byte(plainText))
	if err != nil {
		return "", err
	}
	return string(encrypted), nil
}

func (r *RsaCipher) DecryptToString(cipherText string) (string, error) {
	decrypted, err := r.Decrypt([]byte(cipherText))
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}

func (r *RsaCipher) Sign(text []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(text)
	signature, err := rsa.SignPKCS1v15(rand.Reader, r.privateKey, crypto.SHA256, h.Sum(nil))
	if err != nil {
		return nil, err
	}
	return signature, nil
}

func (r *RsaCipher) VerifySign(text, signature []byte) (bool, error) {
	hashed := sha256.Sum256(text)
	err := rsa.VerifyPKCS1v15(r.publicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		return false, err
	}
	return true, nil
}
