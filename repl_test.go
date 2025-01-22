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
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "helloWorld",
			expected: []string{"helloWorld"},
		},
		{
			input:    "bonjour le monde    ",
			expected: []string{"bonjour", "le", "monde"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Fatalf("actual and expected different lengths")
		}
		for i := range actual {
			word := actual[i]
			expected := c.expected[i]
			if word != expected {
				t.Fatalf("%s != %s", word, expected)
			}
		}
	}
}
