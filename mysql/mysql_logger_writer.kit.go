package mysqlutil

import (
	"encoding/json"
	"fmt"
	"io"
	stdlog "log"
	"os"
	"time"

	"gorm.io/gorm/logger"

	timeutil "github.com/ikaiguang/go-srv-kit/time"
)

// NewStdWriter
func NewStdWriter() logger.Writer {
	return stdlog.New(os.Stderr, "\r\n", stdlog.LstdFlags)
}

// NewJSONWriter .
func NewJSONWriter(w io.Writer) logger.Writer {
	return &jsonWriter{w: w}
}

// NewWriter .
func NewWriter(w io.Writer) logger.Writer {
	return &writer{w: w}
	//return stdlog.New(w, "\r\n", stdlog.LstdFlags)
}

// NewDummyWriter .
func NewDummyWriter() logger.Writer {
	return &dummy{}
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

// jsonWriter 实现 logger.Writer
type jsonWriter struct {
	w io.Writer
}

// jsonStruct json结构
type jsonStruct struct {
	Name string `json:"name"`
	Time string `json:"time"`
	Msg  string `json:"msg"`
}

// Printf 输出
func (w *jsonWriter) Printf(format string, args ...interface{}) {
	bodyBytes, _ := json.Marshal(&jsonStruct{
		Name: "mysql",
		Time: time.Now().Format(timeutil.YmdHmsMLogger),
		Msg:  fmt.Sprintf(format, args...),
	})
	bodyBytes = append(bodyBytes, '\n')
	_, _ = w.w.Write(bodyBytes)
}

// dummy 实现 logger.Writer
type dummy struct{}

// Printf 输出
func (w *dummy) Printf(string, ...interface{}) {}
