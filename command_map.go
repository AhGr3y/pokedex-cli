package main

import (
	"fmt"
)

func commandMap(config *config) error {
	// Handle last page
	if config.nextURL != nil && config.prevURL != nil { // Prevent triggering on first map command
		if config.nextURL == nil {
			fmt.Println("You're on the last page!")
			return nil
		}
	}

	// Get LocationData from pokeapi
	locationData, err := config.pokeapiClient.GetLocationData(config.nextURL)
	if err != nil {
		return fmt.Errorf("error getting location data: %w", err)
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
	// Handle first page
	if config.prevURL == nil {
		fmt.Println("You're on the first page!")
		return nil
	}

	// Get LocationData from pokeapi
	locationData, err := config.pokeapiClient.GetLocationData(config.prevURL)
	if err != nil {
		return fmt.Errorf("error getting location data: %w", err)
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
