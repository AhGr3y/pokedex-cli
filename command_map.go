package main

import (
	"errors"
	"fmt"

	"github.com/AhGr3y/pokedex-cli/internal/pokeapi"
)

func commandMap(config *pokeapi.Config) error {
	fmt.Printf("prev: %v next: %v\n", config.Prev, config.Next)
	// Get locations using pokeapi
	locations, err := pokeapi.GetLocations(config.Next)
	if err != nil {
		// Prevent map if user on last page
		if errors.Is(err, pokeapi.ErrEndOfMap) {
			fmt.Println("You're on the last page!")
			return nil
		}
		return fmt.Errorf("error getting location: %w", err)
	}

	// Update config's next and prev url
	err = config.UpdateOnMap(config.Next)
	if err != nil {
		return fmt.Errorf("error updating config on map: %w", err)
	}

	// Display location names to std.out
	for _, location := range locations {
		fmt.Println(location.Name)
	}

	fmt.Printf("prev: %v next: %v\n", config.Prev, config.Next)
	return nil
}

func commandMapb(config *pokeapi.Config) error {
	fmt.Printf("prev: %v next: %v\n", config.Prev, config.Next)
	// Get locations using pokeapi
	locations, err := pokeapi.GetLocations(config.Prev)
	if err != nil {
		// Prevent map back if user on fisrt page
		if errors.Is(err, pokeapi.ErrEndOfMap) {
			fmt.Println("You're on the first page!")
			return nil
		}
		return fmt.Errorf("error getting location: %w", err)
	}

	// Update config's next and prev url
	err = config.UpdateOnMapb(config.Prev)
	if err != nil {
		return fmt.Errorf("error updating config on map: %w", err)
	}

	// Display location names to std.out
	for _, location := range locations {
		fmt.Println(location.Name)
	}

	fmt.Printf("prev: %v next: %v\n", config.Prev, config.Next)
	return nil
}
