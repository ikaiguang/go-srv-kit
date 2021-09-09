package logutil

import (
	"github.com/go-kratos/kratos/v2/log"
)

// dummy .
type dummy struct{}

// NewDummyLogger 假啊日志手柄
func NewDummyLogger() (log.Logger, error) {
	return &dummy{}, nil
}

// Log .
func (s *dummy) Log(level log.Level, keyvals ...interface{}) (err error) {
	return err
}
