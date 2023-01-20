package utils

import (
	"os"
	"path/filepath"
)

func OpenFile_A(path string) (*os.File, error) {
	dir, _ := filepath.Split(path)
	os.MkdirAll(dir, os.ModePerm)
	writer, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		return nil, err
	} else {
		defer writer.Close()
		return writer, err
	}
}
