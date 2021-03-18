package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxOccuringChar(t *testing.T) {
	tests := []struct {
		name string
		in   string
		char string
	}{
		{
			name: "Empty string",
			in:   "",
			char: "",
		}, {
			name: "Non-empty string - no repeated characters",
			in:   "a",
			char: "a",
		}, {
			name: "Non-empty string - 1 repeated character",
			in:   "bbc",
			char: "b",
		}, {
			name: "Non-empty string - 3 non-repeated characters",
			in:   "abacadbffff",
			char: "f",
		}, {
			name: "Non-empty string - all non-repeated characters",
			in:   "1234",
			char: "1",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			c := maxOccuringChar(tc.in)
			assert.Equal(t, tc.char, c)
		})
	}
}
