package aespkg

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

// Encryptor ...
type Encryptor interface {
	EncryptToString(plaintext string) (string, error)
	DecryptToString(ciphertext string) (string, error)
}

// aesCipher aes加解密
type aesCipher struct {
	key       []byte
	block     cipher.Block
	blockSize int
}

// NewAesCipher aes加解密
func NewAesCipher(key []byte) (Encryptor, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	a := &aesCipher{
		key:       key,
		block:     block,
		blockSize: block.BlockSize(),
	}
	return a, nil
}

// EncryptToString 加密
func (a *aesCipher) EncryptToString(msg string) (string, error) {
	// 转成字节数组
	origData := []byte(msg)
	// 补全码
	origData = pkcs7Padding(origData, a.block.BlockSize())
	// 创建数组
	encrypted := make([]byte, len(origData))
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(a.block, a.key[:a.blockSize])
	// 加密
	blockMode.CryptBlocks(encrypted, origData)
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// DecryptToString 解密
func (a *aesCipher) DecryptToString(encrypted string) (string, error) {
	encryptedByte, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}
	if len(encryptedByte)%a.blockSize != 0 {
		return "", errors.New("encrypted wrong")
	}
	// 创建数组
	decrypted := make([]byte, len(encryptedByte))
	// 解密
	blockMode := cipher.NewCBCDecrypter(a.block, a.key[:a.blockSize])
	blockMode.CryptBlocks(decrypted, encryptedByte)
	// 去补全码
	decrypted, err = pkcs7UnPadding(decrypted)
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}

// 补码
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// 去码
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("invalid UnPadding string")
	}
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}
