package utils

import (
	"os"
	"path/filepath"
)

func OpenFile_A(path string) (*os.File, error) {
	dir, _ := filepath.Split(path)
	os.MkdirAll(dir, os.ModePerm)
	return os.OpenFile(path, os.O_CREATE|os.O_APPEND, 0666)
}
