package mysqlutil

import (
	"fmt"
	"io"
	stdlog "log"
	"os"

	"gorm.io/gorm/logger"
)

// NewWriter .
func NewWriter(w io.Writer) logger.Writer {
	return &writer{w: w}
}

// NewStdWriter
func NewStdWriter() logger.Writer {
	return stdlog.New(os.Stderr, "\r\n", stdlog.LstdFlags)
}

// writer 实现 logger.Writer
type writer struct {
	w io.Writer
}

// Printf 输出
func (w *writer) Printf(format string, args ...interface{}) {
	_, _ = w.w.Write([]byte(fmt.Sprintf(format, args...) + "\n"))
}

// multiWriter 实现 logger.Writer
type multiWriter struct {
	writers []logger.Writer
}

// Printf 输出
func (w *multiWriter) Printf(format string, args ...interface{}) {
	for _, ww := range w.writers {
		ww.Printf(format, args...)
	}
}
