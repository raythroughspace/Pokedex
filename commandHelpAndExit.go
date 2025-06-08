package main
import (
	"fmt"
	"os"
)

func commandExit(pg *locationAreaPage, locCache *Cache, location string,
	pokename string, pokedex map[string]pokemon) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(pg *locationAreaPage, locCache *Cache, location string,
	pokename string, pokedex map[string]pokemon) error {
	help := "Welcome to the Pokedex!\nUsage:\n\n"
	for name, command := range getCommands() {
		help = fmt.Sprintf("%v%v: %v\n", help, name, command.description)
	}
	fmt.Print(help)
	return nil
}