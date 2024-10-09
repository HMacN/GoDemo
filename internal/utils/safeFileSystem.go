package utils

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type SafeFileSystem struct {
	Files http.FileSystem
}

func (fs SafeFileSystem) Open(path string) (http.File, error) {
	file, err := fs.Files.Open(path)
	if err != nil {
		return nil, err
	}

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if stat.IsDir() {
		index := filepath.Join(path, "index.html")

		isEmpty, err := IsEmpty(index)
		if err != nil {
			return nil, err
		}

		if isEmpty {
			return nil, os.ErrNotExist
		}
	}

	return file, nil
}

func IsEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}
