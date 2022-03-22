package setup

import (
	"io"
	stdlog "log"
	"sync"

	writerutil "github.com/ikaiguang/go-srv-kit/kit/writer"
)

// LoggerFileWriter 文件日志写手柄
func (s *modules) LoggerFileWriter() (io.Writer, error) {
	var err error
	s.loggerFileWriterMutex.Do(func() {
		s.loggerFileWriter, err = s.loadingLoggerFileWriter()
	})
	if err != nil {
		s.loggerFileWriterMutex = sync.Once{}
		return nil, err
	}
	if s.loggerFileWriter != nil {
		return s.loggerFileWriter, err
	}

	s.loggerFileWriter, err = s.loadingLoggerFileWriter()
	if err != nil {
		return nil, err
	}
	return s.loggerFileWriter, err
}

// loadingLoggerFileWriter 启动日志文件写手柄
func (s *modules) loadingLoggerFileWriter() (io.Writer, error) {
	fileLoggerConfig := s.Config.LoggerConfigForFile()
	if !s.Config.EnableLoggingFile() || fileLoggerConfig == nil {
		stdlog.Println("|*** 加载日志工具：虚拟的文件写手柄")
		return writerutil.NewDummyWriter()
	}
	rotateConfig := &writerutil.ConfigRotate{
		Dir:            fileLoggerConfig.Dir,
		Filename:       fileLoggerConfig.Filename,
		RotateTime:     fileLoggerConfig.RotateTime.AsDuration(),
		RotateSize:     fileLoggerConfig.RotateSize,
		StorageCounter: uint(fileLoggerConfig.StorageCounter),
		StorageAge:     fileLoggerConfig.StorageAge.AsDuration(),
	}
	return writerutil.NewRotateFile(rotateConfig)
}
