package zippkg

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	filepkg "github.com/ikaiguang/go-srv-kit/kit/file"
	filepathpkg "github.com/ikaiguang/go-srv-kit/kit/filepath"
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
	fps, entries, err := filepathpkg.WalkDir(resourcePath)
	if err != nil {
		return err
	}
	for i := range fps {
		if entries[i].IsDir() {
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
	if zipWriter == nil {
		return errors.New("zip writer is nil")
	}
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
		err = ExtractZipEntry(rf, unzipResourceDir)
		if err != nil {
			return err
		}
	}
	return err
}

// ExtractZipEntry extracts one ZIP entry into the destination directory.
func ExtractZipEntry(zipFile *zip.File, unzipResourceDir string) (err error) {
	if zipFile == nil {
		return errors.New("zip file entry is nil")
	}
	// 输出文件
	outputPath, err := safeUnzipPath(unzipResourceDir, zipFile.Name)
	if err != nil {
		return err
	}

	// 创建文件夹
	if zipFile.FileInfo().IsDir() {
		if err = rejectSymlinkPath(unzipResourceDir, outputPath); err != nil {
			return err
		}
		err = os.MkdirAll(outputPath, 0o755)
		if err != nil {
			return err
		}
		return err
	}

	parentDir := filepath.Dir(outputPath)
	if err = rejectSymlinkPath(unzipResourceDir, parentDir); err != nil {
		return err
	}
	if err = os.MkdirAll(parentDir, 0o755); err != nil {
		return err
	}
	if err = rejectSymlinkPath(unzipResourceDir, outputPath); err != nil {
		return err
	}
	mode := zipFile.Mode().Perm()
	if mode == 0 {
		mode = 0o644
	}
	// 创建输出文件
	outputFile, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, mode)
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

// Deprecated: use ExtractZipEntry instead.
func UnzipFn(zipFile *zip.File, unzipResourceDir string) error {
	return ExtractZipEntry(zipFile, unzipResourceDir)
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

func rejectSymlinkPath(destDir, targetPath string) error {
	cleanDest, err := filepath.Abs(destDir)
	if err != nil {
		return err
	}
	cleanTarget, err := filepath.Abs(targetPath)
	if err != nil {
		return err
	}
	rel, err := filepath.Rel(cleanDest, cleanTarget)
	if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(os.PathSeparator)) {
		return fmt.Errorf("illegal file path in zip: %s", targetPath)
	}

	current := cleanDest
	parts := []string{}
	if rel != "." {
		parts = strings.Split(rel, string(os.PathSeparator))
	}
	for i := -1; i < len(parts); i++ {
		if i >= 0 {
			current = filepath.Join(current, parts[i])
		}
		info, statErr := os.Lstat(current)
		if os.IsNotExist(statErr) {
			return nil
		}
		if statErr != nil {
			return statErr
		}
		if info.Mode()&os.ModeSymlink != 0 {
			return fmt.Errorf("refusing to extract through symlink: %s", current)
		}
	}
	return nil
}
