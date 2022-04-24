package setup

import (
	"github.com/go-kratos/kratos/v2/log"
)

// LoggerPrefixField .
func (s *modules) LoggerPrefixField() *LoggerPrefixField {
	s.loggerPrefixFieldMutex.Do(func() {
		s.loggerPrefixField = s.assemblyLoggerPrefixField()
	})
	return s.loggerPrefixField
}

// assemblyLoggerPrefixField 组装日志前缀
func (s *modules) assemblyLoggerPrefixField() *LoggerPrefixField {
	appConfig := s.AppConfig()

	return &LoggerPrefixField{
		AppName:    appConfig.Name,
		AppVersion: appConfig.Version,
		AppEnv:     appConfig.Env,
	}
}

// withLoggerPrefix ...
func (s *modules) withLoggerPrefix(logger log.Logger) log.Logger {
	var (
		prefixKey   = "app"
		prefixField = s.LoggerPrefixField()
	)
	return log.With(logger, prefixKey, prefixField.String())
}
