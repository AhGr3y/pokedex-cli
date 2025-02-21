package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"strings"

	"github.com/AhGr3y/pokedex-cli/internal/pokeapi"
)

const (
	lowestBaseExp  = 36.0
	highestBaseExp = 608.0
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
	normalizedBaseExp := normalizeBaseExp(pokemon.BaseExperience)
	catchProbability := getCatchProbability(normalizedBaseExp)
	rolledProbability := roundToOneDecimal(rand.Float64() * 100)
	fmt.Printf("catch probability: %v%%\n", catchProbability)
	fmt.Printf("rolled probability: %v%%\n", rolledProbability)
	return rolledProbability <= catchProbability
}

func getCatchProbability(normalizedBaseExp float64) float64 {
	return roundToOneDecimal((1.0 - normalizedBaseExp) * 100)
}

func normalizeBaseExp(baseExp int) float64 {
	return (float64(baseExp) - lowestBaseExp) / (highestBaseExp - lowestBaseExp)
}

// Assume x is between 0 and 100 inclusive
func roundToOneDecimal(x float64) float64 {
	if x > 0 {
		return math.Round(x*10) / 10
	}

	return x
}
