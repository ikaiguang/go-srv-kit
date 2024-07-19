package logpkg

import (
	"github.com/go-kratos/kratos/v2/log"
)

// dummy .
type dummy struct{}

// NewDummyLogger ...
func NewDummyLogger() (log.Logger, error) {
	return &dummy{}, nil
}

// NewNopLogger ...
func NewNopLogger() log.Logger {
	return &dummy{}
}

// Log .
func (s *dummy) Log(level log.Level, keyvals ...interface{}) error {
	return nil
}
