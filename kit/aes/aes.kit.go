package aespkg

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	stderrors "errors"
	"io"
)

// AESEncryptor ...
type AESEncryptor interface {
	EncryptToString(plaintext, key string) (string, error)
	DecryptToString(ciphertext, key string) (string, error)
}

// cbcEncryptor ...
type cbcEncryptor struct{}

func NewCBCCipher() AESEncryptor {
	return &cbcEncryptor{}
}

func (s *cbcEncryptor) EncryptToString(plaintext, key string) (string, error) {
	return EncryptCBC([]byte(plaintext), []byte(key))
}

func (s *cbcEncryptor) DecryptToString(ciphertext, key string) (string, error) {
	return DecryptCBC(ciphertext, []byte(key))
}

// EncryptCBC CBC模式加密
func EncryptCBC(rawData, key []byte) (res string, err error) {
	// block
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	// 填充原文
	blockSize := block.BlockSize()
	padding := blockSize - len(rawData)%blockSize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	rawData = append(rawData, paddingText...)

	// 初始向量IV必须是唯一，但不需要保密
	cipherText := make([]byte, blockSize+len(rawData))
	iv := cipherText[:blockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	// block大小和初始向量大小一定要一致
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[blockSize:], rawData)

	res = base64.URLEncoding.EncodeToString(cipherText)
	return
}

// DecryptCBC CBC模式解密
func DecryptCBC(rawData string, key []byte) (res string, err error) {
	// base64
	encryptData, err := base64.URLEncoding.DecodeString(rawData)
	if err != nil {
		return res, err
	}

	// block
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	blockSize := block.BlockSize()

	if len(encryptData) < blockSize {
		err = stderrors.New("cipherText too short")
		return
	}
	iv := encryptData[:blockSize]
	encryptData = encryptData[blockSize:]

	// CBC mode always works in whole blocks.
	if len(encryptData)%blockSize != 0 {
		err = stderrors.New("cipherText is not a multiple of the block size")
		return
	}

	// CryptBlocks can work in-place if the two arguments are the same.
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(encryptData, encryptData)

	// 解填充
	unPadding := int(encryptData[len(encryptData)-1])
	res = string(encryptData[:(len(encryptData) - unPadding)])
	return
}
