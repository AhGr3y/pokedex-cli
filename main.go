package main

import (
	"time"

	"github.com/AhGr3y/pokedex-cli/internal/pokeapi"
	"github.com/AhGr3y/pokedex-cli/internal/pokecache"
)

func main() {
	config := config{
		pokeapiClient: pokeapi.PokeClient{},
		pokecache:     pokecache.NewCache(time.Second * 5),
	}
	startRepl(&config)
}
