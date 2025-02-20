package main

import (
	"time"

	"github.com/AhGr3y/pokedex-cli/internal/pokeapi"
)

func main() {
	config := config{
		pokeapiClient: pokeapi.NewClient(time.Minute * 5),
	}
	startRepl(&config)
}
