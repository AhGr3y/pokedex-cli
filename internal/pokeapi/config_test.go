package pokeapi

import (
	"errors"
	"log"
	"net/url"
	"testing"
)

func TestGetOffsetFromURL(t *testing.T) {
	// Initialize test urls
	urlWithOffset, err := url.Parse("https://pokeapi.co/api/v2/location-area/?offset=20&limit=20")
	if err != nil {
		log.Fatalf("error parsing string to URL: %v", err)
	}

	urlWithoutOffset, err := url.Parse("https://pokeapi.co/api/v2/location-area/?limit=20")
	if err != nil {
		log.Fatalf("error parsing string to URL: %v", err)
	}

	// Define test cases
	cases := map[string]struct {
		input    *url.URL
		expected int
		err      error
	}{
		"nil URL": {
			input:    nil,
			expected: 0,
			err:      ErrNilURL,
		},
		"query with offset": {
			input:    urlWithOffset,
			expected: 20,
			err:      nil,
		},
		"query without offset": {
			input:    urlWithoutOffset,
			expected: 0,
			err:      ErrURLWithoutOffset,
		},
	}

	// Test the test cases
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			output, err := getOffsetFromURL(c.input)

			// Test for expected error message
			if err != nil {
				if !errors.Is(err, c.err) {
					t.Errorf("unexpected error: %v", err)
				}
			}

			// Test for expected output
			if output != c.expected {
				t.Errorf("unexpected output: output: %d != c.expected: %d\n", output, c.expected)
			}
		})
	}
}

func TestGetOffsetFromQuery(t *testing.T) {
	// Define test cases
	cases := map[string]struct {
		input    string
		expected int
		err      error
	}{
		"empty string": {
			input:    "",
			expected: 0,
			err:      ErrEmptyQueryParameter,
		},
		"with offset one": {
			input:    "?offset=20",
			expected: 20,
			err:      nil,
		},
		"with offset many": {
			input:    "?offset=20&limit=20",
			expected: 20,
			err:      nil,
		},
		"without offset one": {
			input:    "?limit=20",
			expected: 0,
			err:      ErrURLWithoutOffset,
		},
		"without offset many": {
			input:    "?limit=20&example=69",
			expected: 0,
			err:      ErrURLWithoutOffset,
		},
	}

	// Test the test cases
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			output, err := getOffsetFromQuery(c.input)

			// Test for expected error message
			if err != nil {
				if !errors.Is(err, c.err) {
					t.Errorf("unexpected error: %v", err)
				}
			}

			// Test for expected output
			if output != c.expected {
				t.Errorf("unexpected output: output: %v != c.expected: %v\n", output, c.expected)
			}
		})
	}
}
