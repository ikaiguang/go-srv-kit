package testping

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// 初始化必要逻辑
	os.Exit(m.Run())
}
