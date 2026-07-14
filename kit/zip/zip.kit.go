package zippkg

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	filepkg "github.com/ikaiguang/go-kit/file"
	filepathpkg "github.com/ikaiguang/go-kit/filepath"
)

// Zip 压缩目录
// @param resourcePath 被压缩资源；例: runtime/videos
// @param zipPath 压缩到zip的路径；例: runtime/zip/videos.zip
func Zip(resourcePath string, zipPath string) error {
	fileInfo, err := os.Stat(resourcePath)
	if err != nil {
		return err
	}
	// 压缩文件
	if !fileInfo.IsDir() {
		return ZipFile(resourcePath, zipPath)
	}

	if err := os.MkdirAll(filepath.Dir(zipPath), filepkg.DefaultFileMode); err != nil {
		return err
	}
	targetFile, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer func() { _ = targetFile.Close() }()

	zipWriter := zip.NewWriter(targetFile)
	defer func() { _ = zipWriter.Close() }()

	// 读取文件
	fps, fis, err := filepathpkg.WaldDir(resourcePath)
	if err != nil {
		return err
	}
	for i := range fps {
		if fis[i].IsDir() {
			continue
		}
		zipFilePath, err := filepath.Rel(resourcePath, fps[i])
		if err != nil {
			return err
		}
		err = AddFileToZip(zipWriter, fps[i], zipFilePath)
		if err != nil {
			return err
		}
	}
	return err
}

// ZipFile 压缩目录
// @param filePath 被压缩资源；例: runtime/videos/a.mp4
// @param zipPath 压缩到zip的路径；例: runtime/zip/videos.zip
func ZipFile(filePath string, zipPath string) error {
	if err := os.MkdirAll(filepath.Dir(zipPath), filepkg.DefaultFileMode); err != nil {
		return err
	}
	targetFile, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer func() { _ = targetFile.Close() }()

	zipWriter := zip.NewWriter(targetFile)
	defer func() { _ = zipWriter.Close() }()

	zipFilePath := filepath.Base(filePath)
	err = AddFileToZip(zipWriter, filePath, zipFilePath)
	if err != nil {
		return err
	}
	return err
}

// AddFileToZip 添加文件到zip
// @param srcFilePath 被压缩资源；例: runtime/videos/xxx.mp4
// @param zipFilePath 压缩到zip的路径；例: videos/test.mp4
func AddFileToZip(zipWriter *zip.Writer, srcFilePath, zipFilePath string) error {
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		return err
	}
	defer func() { _ = srcFile.Close() }()

	// 写入文件
	zipFile, err := zipWriter.Create(zipFilePath)
	if err != nil {
		return err
	}

	// // 写入文件
	_, err = io.Copy(zipFile, srcFile)
	if err != nil {
		return err
	}
	return err
}

// Unzip 解压资源
// @param zipPath 压缩资源；例: runtime/zip/videos.zip
// @param unzipResourceDir 解缩到zip的路径；例: runtime/videos
func Unzip(zipPath, unzipResourceDir string) (err error) {
	reader, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer func() { _ = reader.Close() }()

	// 解压文件
	for _, rf := range reader.File {
		err = UnzipFn(rf, unzipResourceDir)
		if err != nil {
			return err
		}
	}
	return err
}

// UnzipFn 解压文件到指定目录
// @param unzipResourceDir 解缩到zip的路径；例: runtime/videos
func UnzipFn(zipFile *zip.File, unzipResourceDir string) (err error) {
	// 输出文件
	outputPath, err := safeUnzipPath(unzipResourceDir, zipFile.Name)
	if err != nil {
		return err
	}

	// 创建文件夹
	if zipFile.FileInfo().IsDir() {
		err = os.MkdirAll(outputPath, filepkg.DefaultFileMode)
		if err != nil {
			return err
		}
		return err
	}

	// 创建输出文件
	outputFile, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, filepkg.DefaultFileMode)
	if err != nil {
		return err
	}
	defer func() { _ = outputFile.Close() }()

	// 打开输入文件
	inputFile, err := zipFile.Open()
	if err != nil {
		return err
	}
	defer func() { _ = inputFile.Close() }()

	// 复制
	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return err
	}
	return err
}

func safeUnzipPath(destDir, zipFileName string) (string, error) {
	cleanDest, err := filepath.Abs(destDir)
	if err != nil {
		return "", err
	}
	outputPath, err := filepath.Abs(filepath.Join(cleanDest, zipFileName))
	if err != nil {
		return "", err
	}
	if outputPath != cleanDest && !strings.HasPrefix(outputPath, cleanDest+string(os.PathSeparator)) {
		return "", fmt.Errorf("illegal file path in zip: %s", zipFileName)
	}
	return outputPath, nil
}
