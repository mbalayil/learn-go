package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindromSubstring(t *testing.T) {
	tests := []struct {
		name      string
		in        string
		substring string
		length    int
	}{
		{
			name:      "Empty string",
			in:        "",
			substring: "",
			length:    0,
		}, {
			name:      "Non-empty string - one element",
			in:        "a",
			substring: "a",
			length:    1,
		}, {
			name:      "Non-empty string - 2 elements",
			in:        "abac",
			substring: "aba",
			length:    3,
		}, {
			name:      "Non-empty string - multiple elements",
			in:        "malayalamnews",
			substring: "malayalam",
			length:    9,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			length, substring := longestPalindromSubstring(tc.in)
			assert.Equal(t, tc.length, length)
			assert.Equal(t, tc.substring, substring)
		})
	}
}
