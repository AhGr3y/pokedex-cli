package main

import "testing"

func TestRoundToOneDecimal(t *testing.T) {
	cases := map[string]struct {
		input    float64
		expected float64
	}{
		"zero": {
			input:    0.0,
			expected: 0.0,
		},
		"negative": {
			input:    -1.2,
			expected: -1.2,
		},
		"less than 10 round up": {
			input:    5.15111,
			expected: 5.2,
		},
		"less than 10 round down": {
			input:    5.11111,
			expected: 5.1,
		},
		"less than 100 round up": {
			input:    50.15111,
			expected: 50.2,
		},
		"less than 100 round down": {
			input:    50.11111,
			expected: 50.1,
		},
		"less than 1000 round up": {
			input:    500.15111,
			expected: 500.2,
		},
		"less than 1000 round down": {
			input:    500.11111,
			expected: 500.1,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			output := roundToOneDecimal(c.input)
			if output != c.expected {
				t.Errorf("unexpected results: output: %v - c.expected: %v", output, c.expected)
			}
		})
	}
}
