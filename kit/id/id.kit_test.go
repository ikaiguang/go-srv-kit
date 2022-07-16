package idutil

import "testing"

// go test -v -count=2 ./kit/id -bench=BenchmarkNew_Xxx -run=BenchmarkNew_Xxx
func BenchmarkNew_Xxx(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = New()
	}
}

// go test -v -count=1 ./kit/id -run=TestNew_Xxx
func TestNew_Xxx(t *testing.T) {
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
		total int = 10e6
		idMap     = make(map[int64]int64, total)
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < total; i++ {
				got := New()
				if _, ok := idMap[got]; ok {
					t.Errorf("重复ID：%d\n", got)
					return
				}
				idMap[got] = got
			}
		})
	}
}
