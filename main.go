package main
import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"sync"
)

type cliCommand struct {
	name string
	description string
	callback func(*locationAreaPage, *Cache, string, string, map[string]pokemon) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit":{
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help":{
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map":{
			name: "map",
			description: "displays 20 location areas",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "displays previous 20 locations areas",
			callback: commandMapb,
		},
		"explore": {
			name: "explore", 
			description: "displays all pokemon in location",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "try to catch pokemon",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "inspect pokemon",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "show all caught pokemon",
			callback: commandPokedex,
		},
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(strings.TrimSpace(text)))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	locCache := Cache {lock : sync.RWMutex{}, entries: make(map[string]cacheEntry)}
	var pg locationAreaPage
	var location string
	pokedex := make(map[string]pokemon)
	var pokename string

	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		words := cleanInput(input)
		
		if len(words) != 0 {
			found := false
			for name,command := range getCommands() {
				if name == words[0] {
					if name == "explore"{
						location = words[1]
					} else if name == "catch" || name == "inspect" {
						pokename = words[1]
					}
					err := command.callback(&pg, &locCache, location, pokename, pokedex)
					if err != nil {
						fmt.Printf("Command %v returned error %v", name, err)
					}
					found = true
				}
			}
			if !found {
				fmt.Println("Unknown command")
			}
		}
	}
}