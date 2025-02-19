package main

import (
	"fmt"

	"github.com/AhGr3y/pokedex-cli/internal/pokeapi"
)

func commandMap(config *pokeapi.Config) error {
	// Get locations using pokeapi
	locations, err := pokeapi.GetLocations(config.Next)
	if err != nil {
		return fmt.Errorf("error getting location: %w", err)
	}

	// Update config's next and prev url
	config.UpdateOnMap(config.Next)

	// Display location names to std.out
	for _, location := range locations {
		fmt.Println(location.Name)
	}

	return nil
}
