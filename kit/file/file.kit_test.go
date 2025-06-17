package filepkg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// go test -v -count 1 ./pkg/file -run TestMoveFileToDir
func TestMoveFileToDir(t *testing.T) {
	var (
		filePath = "./../../test/temp1/temp1.txt"
		fileDir  = "./../../test/temp2"
	)
	targetPath, err := MoveFileToDir(filePath, fileDir)
	require.Nil(t, err)
	t.Log(targetPath)
}
