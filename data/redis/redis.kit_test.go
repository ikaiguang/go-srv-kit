package redispkg

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
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

func TestNewClient(t *testing.T) {
	config := testRedisConfig(t)
	db, err := NewClient(config)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
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

func TestClientConstructorsRejectInvalidConfig(t *testing.T) {
	constructors := map[string]func(*Config) (redis.UniversalClient, error){
		"NewClient": NewClient,
		"NewDB":     NewDB,
	}
	for name, constructor := range constructors {
		t.Run(name+"/nil", func(t *testing.T) {
			if db, err := constructor(nil); err == nil || db != nil {
				t.Fatalf("%s(nil) = (%v, %v), want (nil, error)", name, db, err)
			}
		})
		t.Run(name+"/empty", func(t *testing.T) {
			if db, err := constructor(&Config{}); err == nil || db != nil {
				t.Fatalf("%s(empty) = (%v, %v), want (nil, error)", name, db, err)
			}
		})
	}
}
