package library

import (
	"io"
	"os"
	"path"
)

type Library struct {
	path string
}

func New(path string) *Library {
	return &Library{path}
}

func (l Library) AddBook(bookPath string) error {
	return copyFile(l.path, bookPath)
}

func copyFile(dstFolder, src string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	dst := path.Join(dstFolder, path.Base(src))
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	cerr := out.Close()
	if err != nil {
		return err
	}
	return cerr
}
