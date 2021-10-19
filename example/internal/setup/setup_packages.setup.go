package setup

import (
	"io"
	stdlog "log"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	pkgerrors "github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	debugutil "github.com/ikaiguang/go-srv-kit/debug"
	logutil "github.com/ikaiguang/go-srv-kit/log"
	loghelper "github.com/ikaiguang/go-srv-kit/log/helper"
	mysqlutil "github.com/ikaiguang/go-srv-kit/mysql"
	redisutil "github.com/ikaiguang/go-srv-kit/redis"
	writerutil "github.com/ikaiguang/go-srv-kit/writer"
)

// up 启动手柄
type up struct {
	Config

	// loggerFileWriterMutex 日志文件写手柄
	loggerFileWriterMutex sync.Once
	loggerFileWriter      io.Writer

	// loggerMutex 日志
	loggerMutex sync.Once
	logger      log.Logger

	// mysqlGormMutex mysql gorm
	mysqlGormMutex sync.Once
	mysqlGormDB    *gorm.DB

	// redisClientMutex redis客户端
	redisClientMutex sync.Once
	redisClient      *redis.Client
}

// NewUpPackages .
func NewUpPackages(conf Config) Packages {
	return newUpHandler(conf)
}

// newUpHandler .
func newUpHandler(conf Config) *up {
	return &up{
		Config: conf,
	}
}

// LoggerFileWriter 文件日志写手柄
func (s *up) LoggerFileWriter() (io.Writer, error) {
	var err error
	s.loggerFileWriterMutex.Do(func() {
		s.loggerFileWriter, err = s.setupLoggerFileWriter()
	})
	if err != nil {
		return nil, err
	}
	if s.loggerFileWriter != nil {
		return s.loggerFileWriter, err
	}

	s.loggerFileWriter, err = s.setupLoggerFileWriter()
	if err != nil {
		return nil, err
	}
	return s.loggerFileWriter, err
}

// Logger 日志处理示例
func (s *up) Logger() (log.Logger, error) {
	var err error
	s.loggerMutex.Do(func() {
		s.logger, err = s.setupLogger()
	})
	if err != nil {
		return nil, err
	}
	if s.logger != nil {
		return s.logger, err
	}
	s.logger, err = s.setupLogger()
	if err != nil {
		return nil, err
	}
	return s.logger, err
}

// LoggerFileWriter 文件日志写手柄
func (s *up) MysqlGormDB() (*gorm.DB, error) {
	var err error
	s.mysqlGormMutex.Do(func() {
		s.mysqlGormDB, err = s.setupMysqlGormDB()
	})
	if err != nil {
		return nil, err
	}
	if s.mysqlGormDB != nil {
		return s.mysqlGormDB, err
	}

	s.mysqlGormDB, err = s.setupMysqlGormDB()
	if err != nil {
		return nil, err
	}
	return s.mysqlGormDB, err
}

// RedisClient redis 客户端
func (s *up) RedisClient() (*redis.Client, error) {
	var err error
	s.redisClientMutex.Do(func() {
		s.redisClient, err = s.setupRedisClient()
	})
	if err != nil {
		return nil, err
	}
	if s.redisClient != nil {
		return s.redisClient, err
	}
	s.redisClient, err = s.setupRedisClient()
	if err != nil {
		return nil, err
	}
	return s.redisClient, err
}

// setupDebugUtil 设置调试工具
func (s *up) setupDebugUtil() error {
	stdlog.Printf("|*** 加载调试工具：%v\n", s.Config.IsDebugMode())
	if !s.Config.IsDebugMode() {
		return nil
	}
	return debugutil.Setup()
}

// setupLogUtil 设置日志工具
func (s *up) setupLogUtil() (err error) {
	logger, err := s.Logger()
	if err != nil {
		return err
	}
	if logger == nil {
		stdlog.Println("|*** 未加载日志工具")
		return err
	}

	loghelper.Setup(logger)
	return err
}

// setupLoggerFileWriter 启动日志文件写手柄
func (s *up) setupLoggerFileWriter() (io.Writer, error) {
	loggerConfig := s.Config.LoggerConfig()
	if !s.Config.EnableLoggingFile() || loggerConfig.File == nil {
		stdlog.Println("|*** 加载日志工具：虚拟的文件写手柄")
		return writerutil.NewDummyWriter()
	}
	rotateConfig := &writerutil.ConfigRotate{
		Dir:            loggerConfig.File.Dir,
		Filename:       loggerConfig.File.Filename,
		RotateTime:     loggerConfig.File.RotateTime.AsDuration(),
		RotateSize:     loggerConfig.File.RotateSize,
		StorageCounter: uint(loggerConfig.File.StorageCounter),
		StorageAge:     loggerConfig.File.StorageAge.AsDuration(),
	}
	return writerutil.NewRotateFile(rotateConfig)
}

// setupLogger 启动日志文件写手柄
func (s *up) setupLogger() (logger log.Logger, err error) {
	// loggers
	var loggers []log.Logger

	// 配置
	loggerConfig := s.Config.LoggerConfig()
	if loggerConfig == nil {
		return logger, err
	}

	// 日志 输出到控制台
	if s.Config.EnableLoggingConsole() && loggerConfig.Console != nil {
		stdlog.Println("|*** 加载日志工具：日志输出到控制台")
		stdLoggerConfig := &logutil.ConfigStd{
			Level:      logutil.ParseLevel(loggerConfig.Console.Level),
			CallerSkip: logutil.DefaultCallerSkip + 2,
		}
		stdLogger, err := logutil.NewStdLogger(stdLoggerConfig)
		if err != nil {
			return logger, err
		}
		loggers = append(loggers, stdLogger)
	}

	// 日志 输出到文件
	if s.Config.EnableLoggingFile() && loggerConfig.File != nil {
		stdlog.Println("|*** 加载日志工具：日志输出到文件")
		// file logger
		fileLoggerConfig := &logutil.ConfigFile{
			Level:      logutil.ParseLevel(loggerConfig.File.Level),
			CallerSkip: logutil.DefaultCallerSkip + 2,

			Dir:      loggerConfig.File.Dir,
			Filename: loggerConfig.File.Filename,

			RotateTime: loggerConfig.File.RotateTime.AsDuration(),
			RotateSize: loggerConfig.File.RotateSize,

			StorageCounter: uint(loggerConfig.File.StorageCounter),
			StorageAge:     loggerConfig.File.StorageAge.AsDuration(),
		}
		writer, err := s.LoggerFileWriter()
		if err != nil {
			return logger, err
		}
		fileLogger, err := logutil.NewFileLogger(
			fileLoggerConfig,
			logutil.WithWriter(writer),
		)
		if err != nil {
			return logger, err
		}
		loggers = append(loggers, fileLogger)
	}

	// 日志工具
	if len(loggers) == 0 {
		return logger, err
	}
	return log.MultiLogger(loggers...), err
}

// setupMysqlGormDB mysql gorm 数据库
func (s *up) setupMysqlGormDB() (*gorm.DB, error) {
	if s.Config.MySQLConfig() == nil {
		stdlog.Println("|*** 加载MySQL-GORM：未初始化")
		return nil, pkgerrors.WithStack(ErrUninitialized)
	}
	stdlog.Println("|*** 加载MySQL-GORM：成功")

	// logger writer
	var (
		writers []logger.Writer
		opts    []mysqlutil.Option
	)
	if s.Config.EnableLoggingConsole() {
		writers = append(writers, mysqlutil.NewStdWriter())
	}
	if s.Config.EnableLoggingFile() {
		writer, err := s.LoggerFileWriter()
		if err != nil {
			return nil, err
		}
		writers = append(writers, mysqlutil.NewJSONWriter(writer))
	}
	if len(writers) > 0 {
		opts = append(opts, mysqlutil.WithWriters(writers...))
	}
	return mysqlutil.NewMysqlDB(s.Config.MySQLConfig(), opts...)
}

// setupRedisClient redis 客户端
func (s *up) setupRedisClient() (*redis.Client, error) {
	if s.Config.RedisConfig() == nil {
		stdlog.Println("|*** 加载Redis客户端：未初始化")
		return nil, pkgerrors.WithStack(ErrUninitialized)
	}
	stdlog.Println("|*** 加载Redis客户端：成功")

	return redisutil.NewDB(s.Config.RedisConfig())
}
