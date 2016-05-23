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

func TestParseDevices(t *testing.T) {
	deviceName := "DEV1"
	deviceURI := "/dev1"
	ini := fmt.Sprintf("library_path=/path\n[devices]\n%v=%v", deviceName, deviceURI)

	configuration, err := parseConfiguration(strings.NewReader(ini))

	if err != nil {
		t.Errorf("Error expected to be nil, found %v", err.Error())
	}
	device := configuration.Devices[0]
	if device.Name != deviceName {
		t.Errorf("DeviceName expected to be %v, found %v", deviceName, device.Name)
	}
	if device.URI != deviceURI {
		t.Errorf("DeviceURI expected to be %v, found %v", deviceURI, device.URI)
	}
}
