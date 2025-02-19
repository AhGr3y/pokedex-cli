package main

import "github.com/AhGr3y/pokedex-cli/internal/pokeapi"

func main() {
	config := config{
		pokeapiClient: pokeapi.PokeClient{},
	}
	startRepl(&config)
}
