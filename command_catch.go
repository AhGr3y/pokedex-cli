package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/AhGr3y/pokedex-cli/internal/pokeapi"
)

func commandCatch(config *config, params ...string) error {
	if len(params) < 1 {
		fmt.Println("Please input to the pokemon to catch.")
		return nil
	}

	pokemonName := strings.ToLower(params[0])
	pokemon, err := config.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		fmt.Printf("error catching pokemon: %s\n", pokemonName)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	caught := catchPokemon(pokemon)
	if caught {
		fmt.Printf("%s was caught!\n", pokemonName)
		config.pokeapiClient.Pokedex[pokemonName] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}

func catchPokemon(pokemon pokeapi.Pokemon) bool {
	catchProbability := getCatchProbability(pokemon.BaseExperience)
	rolledProbability := rand.Intn(101)
	fmt.Printf("%v has %v base experience\n", pokemon.Name, pokemon.BaseExperience)
	fmt.Printf("catch probability: %v%%\n", catchProbability)
	fmt.Printf("rolled probability: %v%%\n", rolledProbability)
	return rolledProbability <= catchProbability
}

func getCatchProbability(baseExp int) int {
	baseModifier := 100
	if baseExp < 50 {
		return baseModifier - 15
	} else if baseExp < 100 {
		return baseModifier - 20
	} else if baseExp < 150 {
		return baseModifier - 25
	} else if baseExp < 200 {
		return baseModifier - 30
	} else if baseExp < 250 {
		return baseModifier - 35
	} else if baseExp < 300 {
		return baseModifier - 40
	} else if baseExp < 350 {
		return baseModifier - 45
	} else if baseExp < 400 {
		return baseModifier - 50
	} else if baseExp < 450 {
		return baseModifier - 55
	} else if baseExp < 500 {
		return baseModifier - 60
	} else if baseExp < 550 {
		return baseModifier - 65
	} else if baseExp < 600 {
		return baseModifier - 70
	} else if baseExp < 650 {
		return baseModifier - 75
	}

	return 0
}
