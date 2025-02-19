package pokeapi

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

const BaseURL = "https://pokeapi.co/api/v2/"

// Define custom error
var ErrNilURL = errors.New("url is nil")
var ErrURLWithoutOffset = errors.New("url does not have offset query")
var ErrEmptyQueryParameter = errors.New("empty query parameter")

type Config struct {
	Next *url.URL
	Prev *url.URL
}

func (c *Config) UpdateOnMap(nextURL *url.URL) {
	// Get query offset of current location
	if c.Next != nil {

	}
}

func getOffsetFromURL(url *url.URL) (int, error) {
	if url == nil {
		return 0, ErrNilURL
	}

	// Get offset value from url query
	var offset int
	offset, err := getOffsetFromQuery(url.RawQuery)
	if err != nil {
		return 0, err
	}

	return offset, nil
}

func getOffsetFromQuery(query string) (int, error) {
	if query == "" {
		return 0, ErrEmptyQueryParameter
	}

	if !strings.Contains(query, "offset") {
		return 0, ErrURLWithoutOffset
	}

	parameters := strings.Split(query, "&")
	for _, parameter := range parameters {
		// Handle offset parameter only
		if strings.Contains(parameter, "offset") {
			offsetValueString := strings.Split(parameter, "=")[1]

			// Convert string to int
			offSetValue, err := strconv.Atoi(offsetValueString)
			if err != nil {
				return 0, fmt.Errorf("error converting string to int: %w", err)
			}

			return offSetValue, nil
		}
	}

	return 0, nil
}
