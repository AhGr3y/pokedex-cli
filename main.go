package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commandMap map[string]cliCommand

func init() {
	// Define commands
	commandMap = map[string]cliCommand{
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
	}
}

func main() {
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
		userCommand := cleanInputs[0]
		command, ok := commandMap[userCommand]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		// Execute command
		if err := command.callback(); err != nil {
			log.Fatalf("error during callback: %v", err)
		}

	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	commands := ""
	for _, v := range commandMap {
		commands += fmt.Sprintf("\n%s: %s", v.name, v.description)
	}

	helpMessage := fmt.Sprintf(`
Welcome to the Pokedex!
Usage:
%s
`, commands)

	fmt.Println(helpMessage)
	return nil
}
