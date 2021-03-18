package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnagram(t *testing.T) {
	tests := []struct {
		name    string
		string1 string
		string2 string
		out     bool
	}{
		{
			name:    "nil string",
			string1: "",
			string2: "",
			out:     true,
		}, {
			name:    "Non nil string - anagram",
			string1: "binary",
			string2: "brainy",
			out:     true,
		}, {
			name:    "Not anagram - some values in string1 not used by string2",
			string1: "munia",
			string2: "muni",
			out:     false,
		}, {
			name:    "Not anagram - some values in string2 not in string1",
			string1: "munia",
			string2: "munia ",
			out:     false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out := anagram(tc.string1, tc.string2)
			assert.Equal(t, tc.out, out)
		})
	}
}
