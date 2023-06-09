package operatorpkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTernary(t *testing.T) {
	intSamples := []struct {
		Cond     bool
		V1       int
		V2       int
		Expected int
	}{
		{true, 1, 2, 1},
		{false, 1, 2, 2},
	}
	for _, sample := range intSamples {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, Ternary[int](sample.Cond, sample.V1, sample.V2), sample.Expected)
		})
	}

	stringSamples := []struct {
		Cond     bool
		V1       string
		V2       string
		Expected string
	}{
		{true, "1", "2", "1"},
		{false, "1", "2", "2"},
	}
	for _, sample := range stringSamples {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, Ternary[string](sample.Cond, sample.V1, sample.V2), sample.Expected)
		})
	}
}
