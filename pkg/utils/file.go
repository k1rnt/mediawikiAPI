package utils

import (
	"io/fs"
	"log"
	"path/filepath"
)

// ファイルパスから拡張子を取り除いたファイル名を返す
func GetFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

// dir から再起的にファイルを探索しstring sliceで返す
func FindWikiFiles(dir string) []string {
	var files []string
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return files
}
