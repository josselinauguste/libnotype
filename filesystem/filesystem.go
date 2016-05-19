package filesystem

import (
	"os"
	"path"
	"io"
	"io/ioutil"
)

func CopyFile(dstFolder, src string) error {
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

func HasExtension(file os.FileInfo, extensions []string) bool {
	for _, extension := range extensions {
		if path.Ext(file.Name()) == extension {
			return true
		}
	}
	return false
}

func SelectFiles(folderPath string, predicate func(os.FileInfo) bool) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return  nil, err
	}
	selectedFiles := make([]os.FileInfo, 0)
	for _, file := range files {
		if file.Mode().IsDir() {
			innerFiles, err := SelectFiles(path.Join(folderPath, file.Name()), predicate)
			if err != nil {
				return  nil, err
			}
			selectedFiles = append(selectedFiles, innerFiles...)
		}
		if file.Mode().IsRegular() && predicate(file) {
			selectedFiles = append(selectedFiles, file)
		}
	}
	return selectedFiles, nil
}
