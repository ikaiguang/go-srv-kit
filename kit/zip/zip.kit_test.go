package ziputil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// go test -v -count=1 ./pkg/zip -test.run=TestZipFile
func TestZipFile(t *testing.T) {
	var (
		filePath = "./../../test/temp1/temp1.txt"
		zipPath  = "./../../test/temp1/temp1.txt.zip"
	)
	err := ZipFile(filePath, zipPath)
	require.Nil(t, err)
	t.Log(zipPath)
}

// go test -v -count=1 ./pkg/zip -test.run=TestZip_Xxx
func TestZip_Xxx(t *testing.T) {
	var (
		resourcePath = "./../../test/temp2"
		zipPath      = "./../../test/temp3/temp2.zip"
	)
	err := Zip(resourcePath, zipPath)
	require.Nil(t, err)
	t.Log(zipPath)
}

// go test -v -count=1 ./pkg/zip -test.run=TestZip_File
func TestZip_File(t *testing.T) {
	var (
		resourcePath = "./../../test/temp3/temp3.txt"
		zipPath      = "./../../test/temp3/temp3.txt.zip"
	)
	err := Zip(resourcePath, zipPath)
	require.Nil(t, err)
	t.Log(zipPath)
}

// go test -v -count=1 ./pkg/zip -test.run=TestUnzip
func TestUnzip(t *testing.T) {
	var (
		//zipPath          = "./../../test/temp3/temp2.zip"
		zipPath          = "./../../test/temp3/temp3.txt.zip"
		unzipResourceDir = "./../../test/testdata"
	)
	err := Unzip(zipPath, unzipResourceDir)
	require.Nil(t, err)
	t.Log(unzipResourceDir)
}
