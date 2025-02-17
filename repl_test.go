package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected []string
	}{
		"empty input": {
			input:    "",
			expected: []string{},
		},
		"all caps": {
			input:    "HELLO WORLD",
			expected: []string{"hello", "world"},
		},
		"whitespace around": {
			input:    "   Hello   World   ",
			expected: []string{"hello", "world"},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			output := cleanInput(c.input)

			if len(output) != len(c.expected) {
				t.Errorf("len(output): %d != len(c.expected): %d", len(output), len(c.expected))
			} else {
				for i := range c.expected {
					if output[i] != c.expected[i] {
						t.Errorf("output[%d]: %s != c.expected[%d]: %s", i, output[i], i, c.expected[i])
					}
				}
			}

		})
	}
}
