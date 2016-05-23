package library

import (
	"os"

	"github.com/josselinauguste/libnotype/filesystem"
)

type Library struct {
	path string
}

func New(path string) *Library {
	return &Library{path}
}

func (l Library) AddFromFile(bookPath string) error {
	return filesystem.CopyFile(l.path, bookPath)
}

func (l Library) List() ([]Book, error) {
	files, err := filesystem.SelectFiles(l.path, bookPredicate)
	if err != nil {
		return nil, err
	}
	var books []Book
	for _, file := range files {
		book := NewBook(file.Name())
		books = append(books, *book)
	}
	return books, err
}

func bookPredicate(file os.FileInfo) bool {
	extensions := []string{".pdf", ".epub"}
	return filesystem.HasExtension(file, extensions)
}

type Book struct {
	Name string
}

func NewBook(name string) *Book {
	return &Book{name}
}
