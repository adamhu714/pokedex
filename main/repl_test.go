package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    " hello ",
			expected: []string{"hello"},
		},
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  Hello World  ",
			expected: []string{"hello", "world"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Lengths do not match: actual - '%v' vs expected - '%v'", actual, c.expected)
			continue
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("Words %d do not match: actual - '%v' vs expected - '%v'", i, actual[i], c.expected[i])
			}
		}
	}
}
