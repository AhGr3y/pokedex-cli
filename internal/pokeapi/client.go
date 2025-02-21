package pokeapi

import (
	"net/http"
	"time"

	"github.com/AhGr3y/pokedex-cli/internal/pokecache"
)

type PokeClient struct {
	Pokedex map[string]Pokemon
	Cache   pokecache.Cache
	Client  http.Client
}

func NewClient(interval time.Duration) PokeClient {
	return PokeClient{
		Pokedex: map[string]Pokemon{},
		Cache:   pokecache.NewCache(interval),
		Client:  http.Client{},
	}
}
