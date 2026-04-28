package bufferpkg

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBuffer(t *testing.T) {
	buf := GetBuffer()
	assert.NotNil(t, buf)
	assert.Equal(t, 0, buf.Len(), "新获取的 buffer 应为空")
	PutBuffer(buf)
}

func TestPutBuffer_Reset(t *testing.T) {
	buf := GetBuffer()
	buf.WriteString("hello world")
	assert.Greater(t, buf.Len(), 0)

	PutBuffer(buf)

	// 重新获取的 buffer 应已被 Reset
	buf2 := GetBuffer()
	assert.Equal(t, 0, buf2.Len(), "归还后重新获取的 buffer 应为空")
	PutBuffer(buf2)
}

func TestGetBuffer_WriteAndRead(t *testing.T) {
	buf := GetBuffer()
	defer PutBuffer(buf)

	buf.WriteString("hello")
	buf.WriteString(" ")
	buf.WriteString("world")

	assert.Equal(t, "hello world", buf.String())
}

func TestGetBuffer_ConcurrentSafety(t *testing.T) {
	const goroutines = 100
	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(idx int) {
			defer wg.Done()
			buf := GetBuffer()
			buf.WriteString("test")
			assert.Equal(t, "test", buf.String())
			PutBuffer(buf)
		}(i)
	}
	wg.Wait()
}
