package writerpkg

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// go test -v ./kit/writer/ -count=1 -test.run=TestNewRotateFile
func TestNewRotateFile(t *testing.T) {
	conf := &ConfigRotate{
		Dir:      ".",
		Filename: "test",

		RotateTime:     time.Second,
		StorageCounter: 10,
	}
	//writer, err := NewRotateFile(conf, WithFilenameSuffix(".testdata.log"))
	writer, err := NewRotateFile(conf)
	require.Nil(t, err)

	total := int(conf.StorageCounter + 1)
	for i := 0; i < total; i++ {
		str := fmt.Sprintf("第 %d 行", i+1)
		_, err = writer.Write([]byte(str))
		require.Nil(t, err)

		time.Sleep(time.Second)
	}
}
