package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func (pc PokeClient) GetLocationData(url *string) (LocationData, error) {
	fullURL := baseURL + "location-area/"
	if url != nil {
		fullURL = *url
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

	var locationData LocationData
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locationData)
	if err != nil {
		return LocationData{}, fmt.Errorf("error decoding response: %w", err)
	}

	return locationData, nil
}
