package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected []string
	}{
		"trim spaces": {
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		"no spaces": {
			input:    "helloworld",
			expected: []string{"helloworld"},
		},
		"multiple spaces": {
			input:    "hello   world",
			expected: []string{"hello", "world"},
		},
		"empty string": {
			input:    "",
			expected: []string{},
		},
		"only spaces": {
			input:    "     ",
			expected: []string{},
		},
		"tabs and newlines": {
			input:    "\thello\tworld\n",
			expected: []string{"hello", "world"},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			actual := cleanInput(c.input)
			if diff := cmp.Diff(c.expected, actual); diff != "" {
				t.Errorf("cleanInput(%q) mismatch (-want +got):\n%s \nactual: %v", c.input, diff, actual)
			}
		})
	}
}
