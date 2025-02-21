package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/AhGr3y/pokedex-cli/internal/pokeapi"
)

func commandInspect(config *config, params ...string) error {
	if len(params) < 1 {
		fmt.Println("Please input a pokemon to inspect.")
		return nil
	}

	pokemonName := strings.ToLower(params[0])
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokemonName)
	data, ok := config.pokeapiClient.Cache.Get(url)
	if !ok {
		fmt.Printf("You have not caught %s yet!\n", pokemonName)
		return nil
	}

	var pokemon pokeapi.Pokemon
	err := json.Unmarshal(data, &pokemon)
	if err != nil {
		return fmt.Errorf("error unmarshalling data: %w", err)
	}

	pokemon.DisplayStats()

	return nil
}
