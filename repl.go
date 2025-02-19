package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/AhGr3y/pokedex-cli/internal/pokeapi"
)

func startRepl() {
	// Initialize config
	defaultLocationURL, err := url.Parse(pokeapi.BaseURL + "location-area/?offset=0&limit=20")
	if err != nil {
		fmt.Printf("Error parsing URL: %v", err)
		os.Exit(1)
	}

	locationCount, err := pokeapi.GetLocationCount()
	if err != nil {
		fmt.Printf("Error retrieving location count: %v", err)
		os.Exit(1)
	}

	config := &pokeapi.Config{
		LocationCount: locationCount,
		Next:          defaultLocationURL,
		Prev:          nil,
	}

	// Initialize scanner
	scanner := bufio.NewScanner(os.Stdin)

	// Start REPL
	for {
		// Prompt user to input data
		fmt.Print("Pokedex > ")

		// Read user input
		if !scanner.Scan() {
			// Handle EOF (e.g. Ctrl+D) or scanner failure
			fmt.Println("\nClosing the Pokedex... Goodbye!")
			break
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading user input: %v\n", err)
			os.Exit(1)
		}

		// Clean user input
		cleanInputs := cleanInput(scanner.Text())

		// Handle empty inputs
		if len(cleanInputs) == 0 {
			continue
		}

		// Check if command exists
		commandName := cleanInputs[0]
		commands := getCommands()
		commandObj, ok := commands[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		// Execute command
		if err := commandObj.callback(config); err != nil {
			log.Fatalf("error during callback: %v", err)
		}

	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "List next location area",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous location area",
			callback:    commandMapb,
		},
	}
}
