package threadpkg

import (
	"github.com/go-kratos/kratos/v2/log"
	logpkg "github.com/ikaiguang/go-srv-kit/kratos/log"
	"testing"
	"time"
)

// go test -v -count=1 ./kratos/thread -test.run=TestGoSafe
func TestGoSafe(t *testing.T) {
	logpkg.Setup(log.DefaultLogger)
	type args struct {
		fn func()
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "#TestGoSafe",
			args: args{fn: func() { panic(1) }},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GoSafe(tt.args.fn)
		})
	}
	time.Sleep(time.Second)
}
