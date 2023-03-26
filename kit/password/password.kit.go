package passwordutil

import (
	"golang.org/x/crypto/bcrypt"
)

// Encrypt 加密密码
func Encrypt(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Compare 比较密码
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Verify 验证
func Verify(hashed, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}

// PasswordHash 密码验证
type PasswordHash struct{}

// NewPasswordHash 密码验证
func NewPasswordHash() *PasswordHash {
	return &PasswordHash{}
}

func (e *PasswordHash) Encrypt(src string) (des string, err error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(src), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (e *PasswordHash) Verify(hashed, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}
