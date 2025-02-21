package main

import "fmt"

func commandPokedex(config *config, params ...string) error {
	if len(params) > 0 {
		fmt.Println("pokedex command does not accept arguments")
		return nil
	}

	pokedex := config.pokeapiClient.Pokedex
	if len(pokedex) == 0 {
		fmt.Println("Your Pokedex is empty! Go catch some pokemon!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for name := range pokedex {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}
