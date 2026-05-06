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

// errInvalidCiphertext 密文格式错误
var errInvalidCiphertext = errors.New("invalid ciphertext: length is not a multiple of block size")

// errInvalidPadding 填充格式错误
var errInvalidPadding = errors.New("invalid pkcs7 padding")

// aesCipher aes加解密
// WARNING: 此实现使用固定 IV（密钥前缀），属于确定性加密。
// 相同明文+相同密钥会产生相同密文，不适用于需要语义安全的场景。
// 新代码建议使用 aes.kit.go 中的 EncryptCBC/DecryptCBC（随机 IV）。
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
	if len(encryptedByte) == 0 || len(encryptedByte)%a.blockSize != 0 {
		return "", errInvalidCiphertext
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
		return nil, errInvalidPadding
	}
	unPadding := int(data[length-1])
	if unPadding == 0 || unPadding > length {
		return nil, errInvalidPadding
	}
	for _, b := range data[length-unPadding:] {
		if int(b) != unPadding {
			return nil, errInvalidPadding
		}
	}
	return data[:(length - unPadding)], nil
}
