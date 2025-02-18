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

// define commands
var commandMap = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
}

func main() {
	// start REPL
	for {
		// command line prompt
		fmt.Print("Pokedex > ")

		// get user input
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Printf("error reading user input: %v", err)
			os.Exit(0)
		}

		// handle user input
		cleanInputs := cleanInput(scanner.Text())
		if len(cleanInputs) == 0 {
			continue
		} else {
			userCommand := cleanInputs[0]

			// check if command exists
			command, ok := commandMap[userCommand]
			if !ok {
				fmt.Println("Unknown command")
				continue
			}

			// execute command
			if err := command.callback(); err != nil {
				log.Fatalf("error during callback: %v", err)
			}

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
