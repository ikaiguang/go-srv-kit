package bufferpkg

import (
	"bytes"
	"sync"
)

var (
	bufferPool = sync.Pool{
		New: func() any {
			return &bytes.Buffer{}
		},
	}
)

// GetBuffer .
func GetBuffer() *bytes.Buffer {
	return bufferPool.Get().(*bytes.Buffer)
}

// PutBuffer .
func PutBuffer(buf *bytes.Buffer) {
	if buf == nil {
		return
	}
	buf.Reset()
	bufferPool.Put(buf)
}
