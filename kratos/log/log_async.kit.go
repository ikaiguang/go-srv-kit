package logpkg

import (
	"io"
	"os"
	"runtime"
	"sync"

	"github.com/valyala/bytebufferpool"
)

const defaultCachePoolSize = 4096

var _ io.Writer = (*AsyncWriter)(nil)

// AsyncWriter ...
type AsyncWriter struct {
	w          io.Writer
	c          chan *bytebufferpool.ByteBuffer
	mu         sync.RWMutex
	closed     bool
	done       sync.WaitGroup
	bufferPool bytebufferpool.Pool
}

// NewAsyncWriter ...
func NewAsyncWriter(writer io.Writer, size int) *AsyncWriter {
	if size <= 0 {
		size = defaultCachePoolSize
	}
	if writer == nil {
		writer = io.Discard
	}
	w := &AsyncWriter{
		w: writer,
		c: make(chan *bytebufferpool.ByteBuffer, size),
	}
	w.done.Add(1)
	go w.loop()
	return w
}

// Write ...
func (w *AsyncWriter) Write(p []byte) (n int, err error) {
	w.mu.RLock()
	defer w.mu.RUnlock()
	if w.closed {
		return 0, os.ErrClosed
	}

	// 异步情况下，不能直接使用传入的切片，需要拷贝，防止被篡改
	buf := w.bufferPool.Get()
	n, err = buf.Write(p)
	if err != nil {
		w.bufferPool.Put(buf)
		return 0, err
	}

	select {
	case w.c <- buf:
	default:
		w.bufferPool.Put(buf)
	}
	return n, nil
}

func (w *AsyncWriter) loop() {
	defer w.done.Done()
	defer Recover()
	for buf := range w.c {
		_, _ = w.w.Write(buf.Bytes())
		w.bufferPool.Put(buf)
	}
}

func (w *AsyncWriter) Close() {
	w.mu.Lock()
	if !w.closed {
		w.closed = true
		close(w.c)
	}
	w.mu.Unlock()
	w.done.Wait()
}

// Recover is used with defer to do cleanup on panics.
// Use it like: defer Recover(func() {})
func Recover() {
	if rerr := recover(); rerr != nil {
		buf := make([]byte, 64<<10) //nolint:mnd
		n := runtime.Stack(buf, false)
		buf = buf[:n]
		Errorf("threadpkg.Recover: %+v\n%s\n", rerr, buf)
	}
}
