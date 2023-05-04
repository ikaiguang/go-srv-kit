package mongoutil

import (
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/event"
)

// NewMonitor ...
func NewMonitor(logger log.Logger) *event.CommandMonitor {
	return &event.CommandMonitor{
		Started: func(ctx context.Context, evt *event.CommandStartedEvent) {
			_ = log.WithContext(ctx, logger).Log(log.LevelDebug,
				"request_id",
				evt.RequestID,
				"database",
				evt.DatabaseName,
				"command_name",
				evt.CommandName,
				"query",
				evt.Command.String(),
			)
		},
		Succeeded: func(ctx context.Context, evt *event.CommandSucceededEvent) {
			d := evt.DurationNanos / (1000 * 1000)
			_ = log.WithContext(ctx, logger).Log(log.LevelDebug,
				"request_id",
				evt.RequestID,
				"command_name",
				evt.CommandName,
				"result",
				evt.Reply.String(),
				"duration",
				strconv.FormatInt(d, 10)+"ms",
			)
		},
	}
}
