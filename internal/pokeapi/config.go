package pokeapi

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Define custom error
var ErrNilURL = errors.New("url is nil")
var ErrURLWithoutOffset = errors.New("url does not have offset query")
var ErrEmptyQueryParameter = errors.New("empty query parameter")
var ErrEndOfMap = errors.New("reached end of map")

type Config struct {
	LocationCount int
	Next          *url.URL
	Prev          *url.URL
}

func (c *Config) UpdateOnMap(nextURL *url.URL) error {
	if nextURL == nil {
		return ErrNilURL
	}

	// Get query offset of nextURL
	nextURLOffset, err := getOffsetFromURL(nextURL)
	if err != nil {
		return err
	}

	// Update c.Next and c.Prev based on nextURLOffset
	if nextURLOffset == 0 { // Will only occur when user execute map for the first time
		c.Prev = nil
	} else {
		prevURLString := fmt.Sprintf("%slocation-area/?offset=%d&limit=20", BaseURL, nextURLOffset-20)
		prevURL, err := url.Parse(prevURLString)
		if err != nil {
			return fmt.Errorf("error parsing string to url: %w", err)
		}
		c.Prev = prevURL
	}

	// c.Next will be nil if user on last page
	if nextURLOffset+20 >= c.LocationCount {
		c.Next = nil
	} else {
		nextNextURLString := fmt.Sprintf("%slocation-area/?offset=%d&limit=20", BaseURL, nextURLOffset+20)
		nextNextURL, err := url.Parse(nextNextURLString)
		if err != nil {
			return fmt.Errorf("error parsing string to url: %w", err)
		}
		c.Next = nextNextURL
	}

	return nil
}

func (c *Config) UpdateOnMapb(prevURL *url.URL) error {
	if prevURL == nil {
		return ErrNilURL
	}

	// Get query offset of prevURL
	prevURLOffset, err := getOffsetFromURL(prevURL)
	if err != nil {
		return err
	}

	// Prevent map back if user on first page
	if c.Prev == nil {
		return ErrEndOfMap
	}

	// Update c.Next and c.Prev based on prevURLOffset
	nextURLString := fmt.Sprintf("%slocation-area/?offset=%d&limit=20", BaseURL, prevURLOffset+20)
	nextURL, err := url.Parse(nextURLString)
	if err != nil {
		return fmt.Errorf("error parsing string to url: %w", err)
	}
	c.Next = nextURL

	if prevURLOffset <= 0 { // c.Prev will be nil if user on first page
		c.Prev = nil
	} else {
		var prevPrevURLString string
		if prevURLOffset >= 1 && prevURLOffset <= 20 { // c.Prev will be set to 0 to prevent negative offset value
			prevPrevURLString = fmt.Sprintf("%slocation-area/?offset=0&limit=20", BaseURL)
		} else {
			prevPrevURLString = fmt.Sprintf("%slocation-area/?offset=%d&limit=20", BaseURL, prevURLOffset-20)
		}
		prevPrevURL, err := url.Parse(prevPrevURLString)
		if err != nil {
			return fmt.Errorf("error parsing string to url: %w", err)
		}
		c.Prev = prevPrevURL
	}

	return nil
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
