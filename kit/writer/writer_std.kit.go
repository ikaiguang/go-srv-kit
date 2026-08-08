package writerpkg

import (
	"io"
	"os"
)

type stdoutWriter struct {
	io.Writer
}

func NewStdoutWriter() (io.Writer, error) {
	return &stdoutWriter{
		Writer: os.Stdout,
	}, nil
}
