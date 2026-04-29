//go:build ignore

package snowflake

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// go test -v -count 1 ./kit/snowflake -run TestID
func TestID(t *testing.T) {
	id := ID()
	assert.Greater(t, id, uint64(0), "ID 应大于 0")
}

// go test -v -count 1 ./kit/snowflake -run TestNextID
func TestNextID(t *testing.T) {
	id, err := NextID()
	require.Nil(t, err)
	assert.Greater(t, id, uint64(0))
}

// go test -v -count 1 ./kit/snowflake -run TestNextID_Uniqueness
func TestNextID_Uniqueness(t *testing.T) {
	total := 10000
	idMap := make(map[uint64]struct{}, total)

	for i := 0; i < total; i++ {
		id, err := NextID()
		require.Nil(t, err, "第 %d 次生成 ID 失败", i)
		_, exists := idMap[id]
		assert.False(t, exists, "第 %d 次生成了重复 ID: %d", i, id)
		idMap[id] = struct{}{}
	}
}

// go test -v -count 1 ./kit/snowflake -run TestNewNode
func TestNewNode(t *testing.T) {
	node, err := NewNode(100)
	require.Nil(t, err)
	require.NotNil(t, node)

	id, err := node.NextID()
	require.Nil(t, err)
	assert.Greater(t, id, uint64(0))
}

// go test -v -count 1 ./kit/snowflake -run TestSetNode
func TestSetNode(t *testing.T) {
	node, err := NewNode(200)
	require.Nil(t, err)

	SetNode(node)

	id := ID()
	assert.Greater(t, id, uint64(0))
}

// go test -v -count 1 ./kit/snowflake -run TestID_Monotonic
func TestID_Monotonic(t *testing.T) {
	var prev uint64
	for i := 0; i < 100; i++ {
		id, err := NextID()
		require.Nil(t, err)
		assert.Greater(t, id, prev, "ID 应单调递增")
		prev = id
	}
}
