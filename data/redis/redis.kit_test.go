package redispkg

import (
	"context"
	"os"
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
)

func testRedisConfig(t *testing.T) *Config {
	t.Helper()
	address := os.Getenv("REDIS_TEST_ADDR")
	if address == "" {
		t.Skip("set REDIS_TEST_ADDR to run Redis integration tests")
	}
	return &Config{
		Addresses:       []string{address},
		DialTimeout:     durationpb.New(time.Second),
		ReadTimeout:     durationpb.New(time.Second),
		WriteTimeout:    durationpb.New(time.Second),
		ConnMaxActive:   10,
		ConnMaxLifetime: durationpb.New(time.Minute),
		ConnMaxIdle:     2,
		ConnMinIdle:     1,
		ConnMaxIdleTime: durationpb.New(time.Minute),
	}
}

func TestNewDB(t *testing.T) {
	config := testRedisConfig(t)
	db, err := NewDB(config)
	if err != nil {
		t.Fatalf("NewDB() error = %v", err)
	}
	defer db.Close()

	ctx := context.Background()
	if err = db.Set(ctx, "foo", "bar", 0).Err(); err != nil {
		t.Fatalf("Set() error = %v", err)
	}
	value, err := db.Get(ctx, "foo").Result()
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if value != "bar" {
		t.Fatalf("Get() = %q, want bar", value)
	}
}

func TestNewDBRejectsInvalidConfig(t *testing.T) {
	if db, err := NewDB(nil); err == nil || db != nil {
		t.Fatalf("NewDB(nil) = (%v, %v), want (nil, error)", db, err)
	}
	if db, err := NewDB(&Config{}); err == nil || db != nil {
		t.Fatalf("NewDB(empty) = (%v, %v), want (nil, error)", db, err)
	}
}
