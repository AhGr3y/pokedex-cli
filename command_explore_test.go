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
			expected: " - pikachu",
		},
		"one word all upper": {
			input:    "PIKACHU",
			expected: " - pikachu",
		},
		"many word all lower": {
			input:    "pikachu is not cute",
			expected: " - pikachu is not cute",
		},
		"many word all upper": {
			input:    "PIKACHU IS NOT CUTE",
			expected: " - pikachu is not cute",
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
