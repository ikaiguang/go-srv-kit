package snowflake

import "testing"

// go test -v -count 1 ./kit/snowflake -run TestNextID
func TestNextID(t *testing.T) {
	tests := []struct {
		name string
		want uint64
	}{
		{
			name: "#TestNextID",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 10; i++ {
				t.Logf("==> id: %d\n", ID())
			}
		})
	}
}
