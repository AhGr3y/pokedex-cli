package main

import (
	"fmt"
)

func commandHelp(config *config) error {
	var commandUsageText string
	for _, v := range getCommands() {
		commandUsageText += fmt.Sprintf("\n%s: %s", v.name, v.description)
	}

	// Create help message format
	helpMessage := fmt.Sprintf(`
Welcome to the Pokedex!
Usage:
%s
`, commandUsageText)

	fmt.Println(helpMessage)
	return nil
}
