package main

import (
	"fmt"
	"strings"
)

func commandInspect(config *config, params ...string) error {
	if len(params) < 1 {
		fmt.Println("Please input a pokemon to inspect.")
		return nil
	}

	pokemonName := strings.ToLower(params[0])
	pokemon, ok := config.pokeapiClient.Pokedex[pokemonName]
	if !ok {
		fmt.Printf("You have not caught %s yet!\n", pokemonName)
		return nil
	}

	pokemon.DisplayStats()

	return nil
}
