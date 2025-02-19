package pokeapi

import "net/http"

type PokeClient struct {
	Client http.Client
}
