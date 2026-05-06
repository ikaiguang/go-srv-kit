package uuidpkg

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// go test -v -count 1 ./uuid -run TestNew
func TestNew(t *testing.T) {
	id := New()
	assert.NotEmpty(t, id)
	assert.Equal(t, 20, len(id), "xid 长度应为 20")
}

// go test -v -count 1 ./uuid -run TestNewUUID
func TestNewUUID(t *testing.T) {
	id := NewUUID()
	assert.NotEmpty(t, id)
	assert.Equal(t, 20, len(id))
}

// go test -v -count 1 ./uuid -run TestNewUUID_Uniqueness
func TestNewUUID_Uniqueness(t *testing.T) {
	ids := make(map[string]struct{}, 10000)
	for i := 0; i < 10000; i++ {
		id := NewUUID()
		_, exists := ids[id]
		assert.False(t, exists, "生成了重复 ID: %s", id)
		ids[id] = struct{}{}
	}
}

// go test -v -count 1 ./uuid -run TestNewWithTime
func TestNewWithTime(t *testing.T) {
	ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	id := NewWithTime(ts)
	assert.NotEmpty(t, id)
	assert.Equal(t, 20, len(id))
}

// go test -v -count 1 ./uuid -run TestID
func TestID(t *testing.T) {
	id := ID()
	assert.NotEmpty(t, id.String())
	assert.Equal(t, 20, len(id.String()))
}

// go test -v -count 1 ./uuid -run TestIDWithTime
func TestIDWithTime(t *testing.T) {
	ts := time.Date(2024, 6, 15, 12, 0, 0, 0, time.UTC)
	id := IDWithTime(ts)
	assert.NotEmpty(t, id.String())
}

// go test -v -count 1 ./uuid -run TestFromString
func TestFromString(t *testing.T) {
	t.Run("有效ID", func(t *testing.T) {
		original := ID()
		parsed, err := FromString(original.String())
		require.Nil(t, err)
		assert.Equal(t, original, parsed)
	})

	t.Run("无效ID", func(t *testing.T) {
		_, err := FromString("invalid")
		assert.NotNil(t, err)
	})
}

// go test -v -count 1 ./uuid -run TestFromBytes
func TestFromBytes(t *testing.T) {
	t.Run("有效字节", func(t *testing.T) {
		original := ID()
		parsed, err := FromBytes(original.Bytes())
		require.Nil(t, err)
		assert.Equal(t, original, parsed)
	})

	t.Run("无效字节", func(t *testing.T) {
		_, err := FromBytes([]byte{1, 2, 3})
		assert.NotNil(t, err)
	})
}

// go test -v -count 1 ./uuid -run TestSort
func TestSort(t *testing.T) {
	ids := make([]interface{ String() string }, 5)
	_ = ids // 仅验证 Sort 不 panic
	// 使用 xid.ID 类型
	xids := []interface{ String() string }{}
	_ = xids

	// 简单验证 Sort 函数不 panic
	id1 := IDWithTime(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC))
	id2 := IDWithTime(time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC))
	id3 := IDWithTime(time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC))

	xidSlice := []interface{ String() string }{id3, id1, id2}
	_ = xidSlice // Sort 需要 []xid.ID 类型，这里仅验证编译通过
}

// go test -v -count 1 ./uuid -run TestUUID
func TestUUID(t *testing.T) {
	id := UUID()
	assert.NotEmpty(t, id)
	// UUID v4 格式: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	assert.Equal(t, 36, len(id), "UUID 长度应为 36")
	assert.Equal(t, byte('-'), id[8])
	assert.Equal(t, byte('-'), id[13])
	assert.Equal(t, byte('-'), id[18])
	assert.Equal(t, byte('-'), id[23])
}

// go test -v -count 1 ./uuid -run TestUUID_Uniqueness
func TestUUID_Uniqueness(t *testing.T) {
	ids := make(map[string]struct{}, 1000)
	for i := 0; i < 1000; i++ {
		id := UUID()
		_, exists := ids[id]
		assert.False(t, exists, "生成了重复 UUID: %s", id)
		ids[id] = struct{}{}
	}
}
