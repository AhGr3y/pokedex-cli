package main

import "fmt"

func commandHelp() error {
	commands := ""
	for _, v := range getCommands() {
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
