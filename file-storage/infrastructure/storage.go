package infrastructure

import (
	"io"
	//"fmt"
	//"file-storing/models"
	// "golang.org/x/xerrors"
	"os"
	"path/filepath"
)

type FileStorer interface {
	Save(filename string, content io.Reader) (string, error)
	Load(pathname string) (io.Reader, error)
}

type LocalStorage struct {
	root string
}

func NewLocalStorage(root string) *LocalStorage {
	os.MkdirAll(root, os.ModePerm)

	return &LocalStorage{
		root: root,
	}
}

func (s *LocalStorage) Save(filename string, content io.Reader) (string, error) {
	pathname := s.fullPath(filename)

	file, err := os.Create(pathname)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, content)
	if err != nil {
		return "", err
	}

	return pathname, nil
}

func (s *LocalStorage) Load(pathname string) (content io.Reader, err error) {
	return os.Open(pathname)
}

func (s *LocalStorage) fullPath(filename string) string {
	return filepath.Join(s.root, filename)
}
