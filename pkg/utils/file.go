package utils

import "path/filepath"

// ファイルパスから拡張子を取り除いたファイル名を返す
func GetFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}
