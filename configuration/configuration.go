package configuration

import (
	"bytes"
	"io"
	"os"

	"github.com/go-ini/ini"
)

type Configuration struct {
	LibraryPath string
}

func ParseConfigurationFile(filePath string) (Configuration, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return Configuration{}, err
	}
	return parseConfiguration(file)
}

func parseConfiguration(file io.Reader) (Configuration, error) {
	configuration := Configuration{}
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(file)
	ini, err := ini.Load(buffer.Bytes())
	if err != nil {
		return configuration, err
	}
	libraryPath, err := ini.Section("").GetKey("library_path")
	if err != nil {
		return configuration, err
	}
	configuration.LibraryPath = libraryPath.String()
	return configuration, nil
}
