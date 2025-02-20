package main

import "testing"

func TestFormatPokemonName(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected string
	}{
		"empty": {
			input:    "",
			expected: "",
		},
		"one word all lower": {
			input:    "pikachu",
			expected: " - Pikachu",
		},
		"one word all upper": {
			input:    "PIKACHU",
			expected: " - PIKACHU",
		},
		"many word all lower": {
			input:    "pikachu is not cute",
			expected: " - Pikachu is not cute",
		},
		"many word all upper": {
			input:    "PIKACHU IS NOT CUTE",
			expected: " - PIKACHU IS NOT CUTE",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			output := formatPokemonName(c.input)

			if output != c.expected {
				t.Errorf("unexpected results: output: %v != c.expected: %v", output, c.expected)
			}
		})
	}
}
