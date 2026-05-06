package pathpkg

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count 1 ./path -run TestPath
func TestPath(t *testing.T) {
	path := Path()
	t.Log("==> path :", path)
	assert.NotEmpty(t, path, "Path() 不应返回空字符串")
	assert.NotEqual(t, ".", path, "Path() 不应返回 '.'")
	// Path() 返回的是当前源文件所在目录，应以 path 结尾
	assert.True(t, strings.HasSuffix(path, "path"), "Path() 应以 'path' 结尾，实际: %s", path)
}
