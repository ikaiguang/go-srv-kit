package setup

import (
	"os"

	"github.com/go-kratos/kratos/v2/log"

	iputil "github.com/ikaiguang/go-srv-kit/kit/ip"
)

// LoggerPrefixField .
func (s *engines) LoggerPrefixField() *LoggerPrefixField {
	s.loggerPrefixFieldMutex.Do(func() {
		s.loggerPrefixField = s.assemblyLoggerPrefixField()
	})
	return s.loggerPrefixField
}

// assemblyLoggerPrefixField 组装日志前缀
func (s *engines) assemblyLoggerPrefixField() *LoggerPrefixField {
	appConfig := s.AppConfig()

	fields := &LoggerPrefixField{
		AppName:    appConfig.Name,
		AppVersion: appConfig.Version,
		AppEnv:     appConfig.Env,
		ServerIP:   iputil.LocalIP(),
	}
	fields.Hostname, _ = os.Hostname()
	return fields
}

// withLoggerPrefix ...
func (s *engines) withLoggerPrefix(logger log.Logger) log.Logger {
	var (
		prefixKey   = "app"
		prefixField = s.LoggerPrefixField()
	)
	return log.With(logger, prefixKey, prefixField.String())
}
