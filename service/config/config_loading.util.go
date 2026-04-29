package configutil

import (
	"path/filepath"
	"runtime"
	"strings"

	stdlog "log"

	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	ConfigMethodLocal  = "local"
	ConfigMethodConsul = "consul"
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

	loadOpts := &options{}
	for _, opt := range loadingOpts {
		opt(loadOpts)
	}

	method := strings.ToLower(bootstrap.GetApp().GetConfigMethod())
	switch method {
	default:
		return bootstrap, err
	case ConfigMethodConsul:
		// 从 consul 加载配置
		if loadOpts.consulConfigLoader == nil {
			e := errorpkg.ErrorBadRequest("[CONFIGURATION] consul config loader not provided; use WithConsulConfigLoader option")
			return nil, errorpkg.WithStack(e)
		}
		if bootstrap.GetConsul() == nil {
			e := errorpkg.ErrorBadRequest("[CONFIGURATION] config error, key = consul")
			return nil, errorpkg.WithStack(e)
		}
		return loadOpts.consulConfigLoader(bootstrap.GetConsul(), bootstrap.GetApp(), loadingOpts...)
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
