package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/AhGr3y/pokedex-cli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.PokeClient
	nextURL       *string
	prevURL       *string
}

func startRepl(config *config) {
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

		// Retrieve paramaters
		var params []string
		if len(cleanInputs) > 1 {
			params = cleanInputs[1:]
		}

		// Execute command
		if err := commandObj.callback(config, params...); err != nil {
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
	callback    func(*config, ...string) error
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
		"explore": {
			name:        "explore <location-name>",
			description: "List all pokemon in location explored",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon-name>",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon-name>",
			description: "List stats of a pokemon",
			callback:    commandInspect,
		},
	}
}
