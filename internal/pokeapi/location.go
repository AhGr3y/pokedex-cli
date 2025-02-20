package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/AhGr3y/pokedex-cli/internal/pokecache"
)

type LocationData struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (pc PokeClient) GetLocationData(url *string, cache pokecache.Cache) (LocationData, error) {
	fullURL := baseURL + "location-area/"
	if url != nil {
		fullURL = *url
	}

	// Return cached data if available
	cachedData, ok := cache.Get(fullURL)
	if ok {
		var locationData LocationData
		err := json.Unmarshal(cachedData, &locationData)
		if err != nil {
			return LocationData{}, fmt.Errorf("error unmarshalling ")
		}

		return locationData, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationData{}, fmt.Errorf("error creating request: %w", err)
	}

	res, err := pc.Client.Do(req)
	if err != nil {
		return LocationData{}, fmt.Errorf("error doing request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationData{}, fmt.Errorf("error reading response: %w", err)
	}

	// Add data to cache
	// Don't need to disrupt program if fail to add
	err = cache.Add(fullURL, data)
	if err != nil {
		fmt.Printf("error adding location data to cache: %v", err)
	}

	var locationData LocationData
	err = json.Unmarshal(data, &locationData)
	if err != nil {
		return LocationData{}, fmt.Errorf("error unmarshalling ")
	}

	return locationData, nil
}
