package configuration

import (
	"fmt"
	"strings"
	"testing"
)

func TestParseLibraryPath(t *testing.T) {
	libraryPath := "/Users/spongebob/Books"
	ini := fmt.Sprintf("library_path=%v\n", libraryPath)

	configuration, err := parseConfiguration(strings.NewReader(ini))

	if err != nil {
		t.Errorf("Error expected to be nil, found %v", err.Error())
	}
	if configuration.LibraryPath != libraryPath {
		t.Errorf("LibraryPath expected to be %v, found %v", libraryPath, configuration.LibraryPath)
	}
}
