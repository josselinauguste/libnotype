package library

import (
	"os"
	"path"
	"testing"
)

func TestAddBookToLibrary(t *testing.T) {
	library := New(testPath(""))

	err := library.AddBook(getTestPDF())

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if _, err := os.Stat(testPath("book.pdf")); os.IsNotExist(err) {
		t.Error("Book is not found in library path")
	}
}

func TestList(t *testing.T) {
	library := New(getFixturesPath())

	books, err := library.ListBooks()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(books) != 2 {
		t.Errorf("Books length expected to be %v, found %v", 2, len(books))
	}
}

func getTestPDF() string {
	return path.Join(getFixturesPath(), "book.pdf")
}

func getFixturesPath() string {
	currentPath, _ := os.Getwd()
	return path.Join(currentPath, "fixtures")
}

func testPath(fileName string) string {
	basePath := "/tmp"
	if fileName != "" {
		return path.Join(basePath, fileName)
	}
	return basePath
}
