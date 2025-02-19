package pokeapi

import "net/url"

const BaseURL = "https://pokeapi.co/api/v2/"

type Config struct {
	Next *url.URL
	Prev *url.URL
}

func (c *Config) updateURLOnMap(next *url.URL) {
	c.Prev = c.Next
	c.Next = next
}
