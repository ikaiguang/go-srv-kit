package mongopkg

import (
	"context"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/v2/event"
)

type monitorOption struct {
	slowThreshold time.Duration
	logCommand    bool
}

// WithCommandLogging controls whether started events include the full Mongo
// command. It should only be enabled when command contents are safe to log.
func WithCommandLogging(enabled bool) MonitorOption {
	return func(o *monitorOption) {
		o.logCommand = enabled
	}
}

type MonitorOption func(*monitorOption)

// WithSlowThreshold 慢查询阈值
func WithSlowThreshold(slowThreshold time.Duration) MonitorOption {
	return func(o *monitorOption) {
		o.slowThreshold = slowThreshold
	}
}

// NewMonitor ...
func NewMonitor(logger *slog.Logger, opts ...MonitorOption) *event.CommandMonitor {
	options := &monitorOption{}
	for i := range opts {
		opts[i](options)
	}
	loggerHandler := loggerWithDefault(logger).With("module", "mongo-driver-monitor")
	return &event.CommandMonitor{
		Started: func(ctx context.Context, evt *event.CommandStartedEvent) {
			attrs := []slog.Attr{
				slog.Int64("request_id", evt.RequestID),
				slog.String("database", evt.DatabaseName),
				slog.String("command_name", evt.CommandName),
			}
			if options.logCommand {
				attrs = append(attrs, slog.String("command", evt.Command.String()))
			}
			loggerHandler.LogAttrs(ctx, slog.LevelDebug, "mongo command started", attrs...)
		},
		Succeeded: func(ctx context.Context, evt *event.CommandSucceededEvent) {
			attrs := []slog.Attr{
				slog.Int64("request_id", evt.RequestID),
				slog.String("database", evt.DatabaseName),
				slog.String("command_name", evt.CommandName),
				slog.String("duration", evt.Duration.String()),
				//"result", evt.Reply.String(),
			}

			if options.slowThreshold > 0 && evt.Duration >= options.slowThreshold {
				loggerHandler.LogAttrs(ctx, slog.LevelWarn, "mongo command succeeded", attrs...)
			} else {
				loggerHandler.LogAttrs(ctx, slog.LevelDebug, "mongo command succeeded", attrs...)
			}
		},
		Failed: func(ctx context.Context, evt *event.CommandFailedEvent) {
			loggerHandler.LogAttrs(
				ctx,
				slog.LevelWarn,
				"mongo command failed",
				slog.Int64("request_id", evt.RequestID),
				slog.String("database", evt.DatabaseName),
				slog.String("command_name", evt.CommandName),
				slog.String("duration", evt.Duration.String()),
				slog.Any("error", evt.Failure),
			)
		},
	}
}

func loggerWithDefault(logger *slog.Logger) *slog.Logger {
	if logger != nil {
		return logger
	}
	return slog.Default()
}
