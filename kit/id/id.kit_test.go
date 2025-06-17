package idpkg

import (
	"github.com/stretchr/testify/require"
	"math"
	"strconv"
	"testing"
	"time"
)

// ===== Benchmark =====
// BenchmarkNew_SonySonyflake-8			31080             38894 ns/op               0 B/op          0 allocs/op
// BenchmarkNew_BwmarrinSnowflake-8		76981             15611 ns/op               0 B/op          0 allocs/op
// ===== Benchmark =====

// go test -v -count 1 ./kit/id -test.bench BenchmarkNew_SonySonyflake -test.run BenchmarkNew_SonySonyflake
// BenchmarkNew_SonySonyflake-8       31080             38894 ns/op               0 B/op          0 allocs/op
func BenchmarkNew_SonySonyflake(b *testing.B) {
	//node, err := NewSonySonyflake(1)
	node, err := NewSonySonyflake(math.MaxUint16)
	if err != nil {
		b.Error(err)
		b.FailNow()
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = node.NextID()
	}
}

// go test -v -count 1 ./kit/id -test.bench BenchmarkNew_BwmarrinSnowflake -test.run BenchmarkNew_BwmarrinSnowflake
// BenchmarkNew_BwmarrinSnowflake-8           76981             15611 ns/op               0 B/op          0 allocs/op
func BenchmarkNew_BwmarrinSnowflake(b *testing.B) {
	//node, err := NewBwmarrinSnowflake(1)
	node, err := NewBwmarrinSnowflake(math.MaxUint16)
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

// go test -v -count 1 ./kit/id -run TestAll_NextID
func TestAll_NextID(t *testing.T) {
	var (
		total int = 1e5
	)
	t.Log("==> total: ", strconv.FormatInt(int64(total), 10))

	n1, err := NewBwmarrinSnowflake(1)
	require.Nil(t, err)

	t.Log("==> start bwmarrinSnowflake ...")
	n1s := time.Now()
	for i := 0; i <= total; i++ {
		_, _ = n1.NextID()
	}
	n1latency := time.Since(n1s)
	t.Logf("==> end bwmarrinSnowflake: %s\n", n1latency)

	n2, err := NewSonySonyflake(1)
	require.Nil(t, err)

	t.Log("==> start sonySonyflake ...")
	n2s := time.Now()
	for i := 0; i <= total; i++ {
		_, _ = n2.NextID()
	}
	n2latency := time.Since(n2s)
	t.Logf("==> end sonySonyflake: %s\n", n2latency)

	if n1latency > n2latency {
		t.Log("==> sonySonyflake is better than bwmarrinSnowflake, gt: ", (n1latency - n2latency).String())
	} else {
		t.Log("==> bwmarrinSnowflake is better than sonySonyflake, gt: ", (n2latency - n1latency).String())
	}
}
