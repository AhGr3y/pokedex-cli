package main

import (
	"fmt"
	"strings"
)

func commandExplore(config *config, params ...string) error {
	if len(params) < 1 {
		fmt.Println("Please provide location area to explore.")
		return nil
	}
	locationArea := params[0]

	locationAreaData, err := config.pokeapiClient.GetLocationAreaData(locationArea)
	if err != nil {
		return fmt.Errorf("error getting location area data: %w", err)
	}

	fmt.Printf("Exploring %s...\n", locationArea)
	fmt.Println("Found Pokemon:")
	pokemonEncounters := locationAreaData.PokemonEncounters
	for _, encounter := range pokemonEncounters {
		fmt.Println(formatPokemonName(encounter.Pokemon.Name))
	}

	return nil
}

func formatPokemonName(name string) string {
	if name == "" {
		return ""
	}

	return " - " + strings.ToLower(name)
}
