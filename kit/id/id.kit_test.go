package idpkg

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testSnowflake struct {
	next atomic.Uint64
}

func (s *testSnowflake) NextID() (uint64, error) {
	return s.next.Add(1), nil
}

func TestConcurrentSetNodeAndNextID(t *testing.T) {
	original := nodeHandler
	t.Cleanup(func() { SetNode(original) })

	var wg sync.WaitGroup
	for range 8 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				SetNode(&testSnowflake{})
				if _, err := NextID(); err != nil {
					t.Errorf("NextID: %v", err)
					return
				}
			}
		}()
	}
	wg.Wait()
}

// ===== Benchmark =====
// BenchmarkNew_BwmarrinSnowflake-8		76981             15611 ns/op               0 B/op          0 allocs/op
// ===== Benchmark =====

// go test -v -count 1 ./kit/id -test.bench BenchmarkNew_BwmarrinSnowflake -test.run BenchmarkNew_BwmarrinSnowflake
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

// go test -v -count 1 ./kit/id -run TestMy_NextID
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

func TestIPV4ToNodeID(t *testing.T) {
	got, err := IPv4ToNodeID("192.168.1.2")
	require.NoError(t, err)
	assert.Equal(t, uint16(258), got)

	got, err = IPv4ToNodeID("192.168.255.255")
	require.NoError(t, err)
	assert.Equal(t, uint16(snowflakeMaxNode), got)

	_, err = IPv4ToNodeID("2001:db8::1")
	require.Error(t, err)

	_, err = IPv4ToNodeID("invalid")
	require.Error(t, err)
}

func TestNewBwmarrinSnowflakeNodeRange(t *testing.T) {
	_, err := NewBwmarrinSnowflake(snowflakeMaxNode)
	require.NoError(t, err)

	_, err = NewBwmarrinSnowflake(snowflakeMaxNode + 1)
	require.Error(t, err)

	_, err = NewBwmarrinSnowflake(-1)
	require.Error(t, err)
}

func TestConcurrentNewBwmarrinSnowflake(t *testing.T) {
	var wg sync.WaitGroup
	for i := range 16 {
		wg.Add(1)
		go func(nodeID int) {
			defer wg.Done()
			for range 100 {
				if _, err := NewBwmarrinSnowflake(int64(nodeID)); err != nil {
					t.Errorf("NewBwmarrinSnowflake: %v", err)
					return
				}
			}
		}(i)
	}
	wg.Wait()
}
