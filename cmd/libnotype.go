package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/go-ini/ini"
	"github.com/josselinauguste/libnotype/library"
)

func main() {
	configuration, err := ini.Load("~/.libnotype")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(2)
	}
	libraryPath, err := configuration.Section("").GetKey("library_path")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(3)
	}
	library := library.New(libraryPath.String())
	command, err := parseCommand(os.Args[1:])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	err = command(library)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(4)
	}
}

func parseCommand(args []string) (func(*library.Library) error, error) {
	if len(args) == 0 {
		return nil, errors.New("Missing command")
	}
	if args[0] == "add" {
		if len(args) > 1 {
			return func(library *library.Library) error {
				return library.AddBook(args[1])
			}, nil
		}
		return nil, errors.New("Missing <library> argument")
	}
	return nil, errors.New("Unknown command")
}
