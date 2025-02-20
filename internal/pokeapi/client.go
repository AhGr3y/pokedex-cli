package pokeapi

import (
	"net/http"
	"time"

	"github.com/AhGr3y/pokedex-cli/internal/pokecache"
)

type PokeClient struct {
	Cache  pokecache.Cache
	Client http.Client
}

func NewClient(interval time.Duration) PokeClient {
	return PokeClient{
		Cache:  pokecache.NewCache(interval),
		Client: http.Client{},
	}
}
