package main

import (
	"errors"
	"fmt"
	"os"

	"os/user"
	"path"

	"github.com/go-ini/ini"
	"github.com/josselinauguste/libnotype/library"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	configurationFile := path.Join(usr.HomeDir, ".libnotype")
	configuration, err := ini.Load(configurationFile)
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
			return addBook(args[1]), nil
		}
		return nil, errors.New("Missing <library> argument")
	}
	if args[0] == "list" {
		return listBooks(), nil
	}
	return nil, errors.New("Unknown command")
}

func addBook(filePath string) func(*library.Library) error {
	return func(library *library.Library) error {
		return library.AddBook(filePath)
	}
}

func listBooks() func(*library.Library) error {
	return func(library *library.Library) error {
		books, err := library.ListBooks()
		if err != nil {
			return err
		}
		if len(books) == 0 {
			fmt.Println("Your library is empty.")
		}
		for _, book := range books {
			fmt.Printf("- %v\n", book.Name)
		}
		return nil
	}
}