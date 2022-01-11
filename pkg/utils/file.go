package utils

import (
	"io/fs"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
)

// ファイルパスから拡張子を取り除いたファイル名を返す
func GetFileNameWithoutExt(path string) (string, error) {
	if !Exists(path) {
		return "", xerrors.Errorf("%s does not exist", path)
	}
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))]), nil
}

// dir から再起的にファイルを探索しstring sliceで返す
func FindWikiFiles(dir string) ([]string, error) {
	if !Exists(dir) {
		return nil, xerrors.Errorf("%s does not exist", dir)
	}
	var files []string
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func Exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}
