package configutil

import (
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	stdlog "log"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	CONFIG_METHOD_LOCAL  = "local"
	CONFIG_METHOD_CONSUL = "consul"
)

func Loading(filePath string, loadingOpts ...Option) (*configpb.Bootstrap, error) {
	bootstrap, err := LoadingFile(filePath, loadingOpts...)
	if err != nil {
		return nil, err
	}
	if bootstrap.GetApp() == nil {
		e := errorpkg.ErrorBadRequest("[CONFIGURATION] config error, key = app")
		return nil, errorpkg.WithStack(e)
	}
	method := strings.ToLower(bootstrap.GetApp().GetConfigMethod())
	switch method {
	default:
		return bootstrap, err
	case CONFIG_METHOD_CONSUL:
		//从consul加载配置
		if bootstrap.GetConsul() == nil {
			e := errorpkg.ErrorBadRequest("[CONFIGURATION] config error, key = consul")
			return nil, errorpkg.WithStack(e)
		}
		consulClient, err := newConsulClient(bootstrap.GetConsul())
		if err != nil {
			return nil, err
		}
		return LoadingConfigFromConsul(consulClient, bootstrap.GetApp(), loadingOpts...)
	}
}

// MergeConfig 合并配置；后面的覆盖前面的
// Merge merges src into dst, which must be a message with the same descriptor.
func MergeConfig(first, second proto.Message) {
	stdlog.Println("|==================== MERGE CONFIGURATION : START ====================|")
	defer stdlog.Println()
	defer stdlog.Println("|==================== MERGE CONFIGURATION : END ====================|")

	// not proto.Merge
	firstMessage := first.ProtoReflect()
	secondMessage := second.ProtoReflect()
	var rangeFn = func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		stdlog.Println("|*** INFO: merge config key: ", fd.Name())
		firstMessage.Set(fd, v)
		return true
	}
	secondMessage.Range(rangeFn)
}

func CurrentPath() string {
	_, file, _, _ := runtime.Caller(0)

	return filepath.Dir(file)
}
