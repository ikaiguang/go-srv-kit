package idpkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ===== Benchmark =====
// BenchmarkNew_BwmarrinSnowflake-8		76981             15611 ns/op               0 B/op          0 allocs/op
// ===== Benchmark =====

// go test -v -count 1 ./id -test.bench BenchmarkNew_BwmarrinSnowflake -test.run BenchmarkNew_BwmarrinSnowflake
// BenchmarkNew_BwmarrinSnowflake-8           76981             15611 ns/op               0 B/op          0 allocs/op
func BenchmarkNew_BwmarrinSnowflake(b *testing.B) {
	//node, err := NewBwmarrinSnowflake(1)
	node, err := NewBwmarrinSnowflake(snowflakeMaxNode)
	if err != nil {
		b.Error(err)
		b.FailNow()
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = node.NextID()
	}
}

// go test -v -count 1 ./id -run TestMy_NextID
func TestMy_NextID(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{
			name: "#唯一ID",
			want: 0,
		},
	}

	var (
		total int = 1e6 // 10^6
		idMap     = make(map[uint64]uint64, total)
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < total; i++ {
				got, err := NextID()
				if err != nil {
					t.Error(err)
					t.FailNow()
				}
				//t.Logf("i: %d, id: %d \n", i, got)
				if _, ok := idMap[got]; ok {
					t.Errorf("重复ID：%d\n", got)
					t.FailNow()
				}
				idMap[got] = got
			}
		})
	}
}

// go test -v -count 1 ./id -run TestIPV4ToNodeID
func TestIPV4ToNodeID(t *testing.T) {
	got, err := IPV4ToNodeID("192.168.1.2")
	require.NoError(t, err)
	assert.Equal(t, uint16(258), got)

	got, err = IPV4ToNodeID("192.168.255.255")
	require.NoError(t, err)
	assert.Equal(t, uint16(snowflakeMaxNode), got)

	_, err = IPV4ToNodeID("2001:db8::1")
	require.Error(t, err)

	_, err = IPV4ToNodeID("invalid")
	require.Error(t, err)
}

// go test -v -count 1 ./id -run TestNewBwmarrinSnowflakeNodeRange
func TestNewBwmarrinSnowflakeNodeRange(t *testing.T) {
	_, err := NewBwmarrinSnowflake(snowflakeMaxNode)
	require.NoError(t, err)

	_, err = NewBwmarrinSnowflake(snowflakeMaxNode + 1)
	require.Error(t, err)

	_, err = NewBwmarrinSnowflake(-1)
	require.Error(t, err)
}
