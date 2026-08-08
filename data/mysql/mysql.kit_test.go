package mysqlpkg

import (
	"errors"
	"fmt"
	"testing"
	"time"

	mysqldriver "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
	"gorm.io/gorm"
)

func TestBuildConnOption(t *testing.T) {
	t.Run("nil config", func(t *testing.T) {
		option, err := buildConnOption(nil)
		require.EqualError(t, err, "mysql config is nil")
		require.Nil(t, option)
	})

	t.Run("nil durations", func(t *testing.T) {
		option, err := buildConnOption(&Config{})
		require.NoError(t, err)
		require.Zero(t, option.SlowThreshold)
		require.Zero(t, option.ConnMaxLifetime)
		require.Zero(t, option.ConnMaxIdleTime)
	})

	t.Run("configured pool", func(t *testing.T) {
		option, err := buildConnOption(&Config{
			SlowThreshold:   durationpb.New(200 * time.Millisecond),
			ConnMaxActive:   100,
			ConnMaxLifetime: durationpb.New(30 * time.Minute),
			ConnMaxIdle:     10,
			ConnMaxIdleTime: durationpb.New(time.Hour),
		})
		require.NoError(t, err)
		require.Equal(t, 200*time.Millisecond, option.SlowThreshold)
		require.Equal(t, 100, option.ConnMaxActive)
		require.Equal(t, 30*time.Minute, option.ConnMaxLifetime)
		require.Equal(t, 10, option.ConnMaxIdle)
		require.Equal(t, time.Hour, option.ConnMaxIdleTime)
	})
}

func TestNewDBNilConfig(t *testing.T) {
	db, err := NewDB(nil)
	require.EqualError(t, err, "mysql config is nil")
	require.Nil(t, db)
}

func TestConfigDescriptorPath(t *testing.T) {
	require.Equal(t, "data/mysql/config.proto", File_data_mysql_config_proto.Path())
}

func TestIsErrDuplicatedKey(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{name: "nil", err: nil, want: false},
		{name: "gorm duplicate", err: gorm.ErrDuplicatedKey, want: true},
		{name: "wrapped driver duplicate", err: fmt.Errorf("insert: %w", &mysqldriver.MySQLError{Number: 1062}), want: true},
		{name: "other driver error", err: &mysqldriver.MySQLError{Number: 1045}, want: false},
		{name: "other error", err: errors.New("other"), want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, IsErrDuplicatedKey(tt.err))
		})
	}
}
