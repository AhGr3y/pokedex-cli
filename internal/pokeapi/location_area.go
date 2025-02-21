package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreaData struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (pc PokeClient) GetLocationAreaData(locationArea string) (LocationAreaData, error) {
	if locationArea == "" {
		return LocationAreaData{}, fmt.Errorf("missing url")
	}

	fullURL := fmt.Sprintf("%slocation-area/%s/", baseURL, locationArea)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaData{}, fmt.Errorf("error creating request: %w", err)
	}

	// Return cached data if available
	if cacheData, ok := pc.Cache.Get(fullURL); ok {
		var locationAreaData LocationAreaData
		err = json.Unmarshal(cacheData, &locationAreaData)
		if err != nil {
			return LocationAreaData{}, fmt.Errorf("error unmarshalling location area data: %w", err)
		}

		return locationAreaData, nil
	}

	res, err := pc.Client.Do(req)
	if err != nil {
		return LocationAreaData{}, fmt.Errorf("error doing request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaData{}, fmt.Errorf("error reading response: %w", err)
	}

	var locationAreaData LocationAreaData
	err = json.Unmarshal(data, &locationAreaData)
	if err != nil {
		return LocationAreaData{}, fmt.Errorf("error unmarshalling location area data: %w", err)
	}

	pc.Cache.Add(fullURL, data)

	return locationAreaData, nil
}
