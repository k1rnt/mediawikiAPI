package utils

import (
	"io/fs"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
)

// pathから拡張子を取り除いたファイル名を返す
func GetFileNameWithoutExt(path string) (string, error) {
	if !Exists(path) {
		return "", xerrors.Errorf("%s does not exist", path)
	}
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))]), nil
}

// dir から再起的にwikiファイルを探索しstring sliceで返す
func FindWikiFiles(dir string) ([]string, error) {
	if !Exists(dir) {
		return nil, xerrors.Errorf("%s does not exist", dir)
	}
	var files []string
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// mediawikiファイル以外は飛ばす
		if !isMediawikiFile(path) {
			return nil
		}

		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

// pathが存在するかどうかを真偽値で返す
func Exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

// pathの拡張子がmediawikiかどうかを真偽値で返す
func isMediawikiFile(path string) bool {
	return filepath.Ext(path) == ".mediawiki"
}
