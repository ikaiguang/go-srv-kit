package mongopkg

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"strings"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/event"
)

func TestMonitorStartedUsesSourceAndOmitsCommandByDefault(t *testing.T) {
	var output bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&output, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))
	monitor := NewMonitor(logger)
	monitor.Started(context.Background(), commandStartedEvent(t))

	record := decodeLogRecord(t, &output)
	if got := record["level"]; got != "DEBUG" {
		t.Fatalf("level = %v, want DEBUG", got)
	}
	if _, ok := record["command"]; ok {
		t.Fatalf("command = %v, want omitted", record["command"])
	}
	source, ok := record["source"].(map[string]any)
	if !ok {
		t.Fatalf("source = %v, want object", record["source"])
	}
	file, _ := source["file"].(string)
	if !strings.HasSuffix(file, "mongo_monitor.kit.go") {
		t.Fatalf("source.file = %q, want mongo_monitor.kit.go suffix", file)
	}
}

func TestMonitorSucceededLogsSlowCommandAtWarn(t *testing.T) {
	var output bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&output, &slog.HandlerOptions{Level: slog.LevelDebug}))
	monitor := NewMonitor(logger, WithSlowThreshold(100*time.Millisecond))
	monitor.Succeeded(context.Background(), &event.CommandSucceededEvent{
		CommandFinishedEvent: event.CommandFinishedEvent{
			Duration:     100 * time.Millisecond,
			DatabaseName: "app",
			CommandName:  "find",
			RequestID:    42,
		},
	})

	record := decodeLogRecord(t, &output)
	if got := record["level"]; got != "WARN" {
		t.Fatalf("level = %v, want WARN", got)
	}
}

func TestMonitorStartedIncludesCommandWhenEnabled(t *testing.T) {
	var output bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&output, &slog.HandlerOptions{Level: slog.LevelDebug}))
	monitor := NewMonitor(logger, WithCommandLogging(true))
	monitor.Started(context.Background(), commandStartedEvent(t))

	record := decodeLogRecord(t, &output)
	command, _ := record["command"].(string)
	if !strings.Contains(command, "find") {
		t.Fatalf("command = %q, want serialized command", command)
	}
}

func commandStartedEvent(t *testing.T) *event.CommandStartedEvent {
	t.Helper()
	command, err := bson.Marshal(bson.D{{Key: "find", Value: "users"}})
	if err != nil {
		t.Fatalf("bson.Marshal() error = %v", err)
	}
	return &event.CommandStartedEvent{
		Command:      command,
		DatabaseName: "app",
		CommandName:  "find",
		RequestID:    42,
	}
}

func decodeLogRecord(t *testing.T, output *bytes.Buffer) map[string]any {
	t.Helper()
	var record map[string]any
	if err := json.Unmarshal(bytes.TrimSpace(output.Bytes()), &record); err != nil {
		t.Fatalf("json.Unmarshal() error = %v; output = %q", err, output.String())
	}
	return record
}
