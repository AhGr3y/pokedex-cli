package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// start REPL
	for {
		// command line prompt
		fmt.Print("Pokedex > ")

		// get user input
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Printf("error reading user input: %w", err)
			os.Exit(0)
		}

		// handle user input
		cleanInputs := cleanInput(scanner.Text())
		if len(cleanInputs) == 0 {
			// do nothing
		} else {
			firstCommand := cleanInputs[0]
			if firstCommand == "exit" { // handle 'exit' command
				if len(cleanInputs) == 1 {
					fmt.Println("Goodbye!")
					os.Exit(0)
				} else { // handle 'exit' command with options
					var options string
					for i := 1; i < len(cleanInputs); i++ {
						options += " " + cleanInputs[i]
					}
					fmt.Printf("unrecognized options for '%s' command:%s\n", firstCommand, options)
				}
			} else { // handle other commands
				// handle other commands
				fmt.Printf("Your command was: %s\n", firstCommand)
			}
		}

	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
