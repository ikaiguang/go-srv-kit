package logpkg

import (
	"context"
	"log/slog"
)

type discardHandler struct{}

func (discardHandler) Enabled(context.Context, slog.Level) bool  { return false }
func (discardHandler) Handle(context.Context, slog.Record) error { return nil }
func (h discardHandler) WithAttrs([]slog.Attr) slog.Handler      { return h }
func (h discardHandler) WithGroup(string) slog.Handler           { return h }

// NewDummyLogger ...
func NewDummyLogger() (*slog.Logger, error) {
	return NewNopLogger(), nil
}

// NewNopLogger ...
func NewNopLogger() *slog.Logger {
	return slog.New(discardHandler{})
}
