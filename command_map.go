package main

import (
	"fmt"
)

func commandMap(config *config) error {
	// Get LocationData using pokeapi
	locationData, err := config.pokeapiClient.GetLocationData(config.nextURL)
	if err != nil {
		return fmt.Errorf("error getting location data: %w", err)
	}

	// Handle last page
	if locationData.Next == nil {
		fmt.Println("You're on the last page!")
		return nil
	}

	// Update config's nextURL and prevURL
	config.nextURL = locationData.Next
	config.prevURL = locationData.Previous

	// Display location names to std.out
	for _, location := range locationData.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(config *config) error {
	// Get LocationData using pokeapi
	locationData, err := config.pokeapiClient.GetLocationData(config.prevURL)
	if err != nil {
		return fmt.Errorf("error getting location data: %w", err)
	}

	// Handle first page
	if locationData.Previous == nil {
		fmt.Println("You're on the first page!")
		return nil
	}

	// Update config's nextURL and prevURL
	config.nextURL = locationData.Next
	config.prevURL = locationData.Previous

	// Display location names to std.out
	for _, location := range locationData.Results {
		fmt.Println(location.Name)
	}

	return nil
}
