package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type LocationData struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Location struct {
	Name string
	URL  string
}

func GetLocations(locationURL *url.URL) ([]Location, error) {
	// Retrieve location data from PokeAPI
	resp, err := http.Get(locationURL.String())
	if err != nil {
		return nil, fmt.Errorf("error fetching resource: %w", err)
	}
	defer resp.Body.Close()

	// Decode json data into LocationData struct
	var locationData LocationData
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&locationData)
	if err != nil {
		return nil, fmt.Errorf("error decoding json data: %w", err)
	}

	// Convert locationData.Results to []Location struct type
	var locations []Location
	for _, result := range locationData.Results {
		locations = append(locations, Location{Name: result.Name, URL: result.URL})
	}

	return locations, nil
}
