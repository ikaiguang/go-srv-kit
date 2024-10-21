package loggerutil

import (
	"io"
	stdlog "log"
	"sync"

	writerpkg "github.com/ikaiguang/go-srv-kit/kit/writer"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

func (s *loggerManager) GetWriter() (io.Writer, error) {
	var err error
	s.writerOnce.Do(func() {
		s.writer, err = s.getWriter()
	})
	if err != nil {
		s.writerOnce = sync.Once{}
	}
	return s.writer, err
}

func (s *loggerManager) getWriter() (io.Writer, error) {
	fileLoggerConfig := s.conf.GetFile()
	if fileLoggerConfig == nil || !fileLoggerConfig.GetEnable() {
		stdlog.Println("|*** LOADING: FakeLogger: ...")
		writer, err := writerpkg.NewDummyWriter()
		if err != nil {
			e := errorpkg.ErrorInternalError(err.Error())
			return nil, errorpkg.WithStack(e)
		}
		return writer, nil
	}

	// rotate write
	rotateConfig := &writerpkg.ConfigRotate{
		Dir:            fileLoggerConfig.GetDir(),
		Filename:       fileLoggerConfig.GetFilename(),
		RotateTime:     fileLoggerConfig.GetRotateTime().AsDuration(),
		RotateSize:     fileLoggerConfig.GetRotateSize(),
		StorageCounter: uint(fileLoggerConfig.GetStorageCounter()),
		StorageAge:     fileLoggerConfig.GetStorageAge().AsDuration(),
	}
	writer, err := writerpkg.NewRotateFile(rotateConfig)
	if err != nil {
		e := errorpkg.ErrorInternalError(err.Error())
		return nil, errorpkg.WithStack(e)
	}
	return writer, nil
}
