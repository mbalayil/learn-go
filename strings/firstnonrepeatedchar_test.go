package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstNonRepeatedChar(t *testing.T) {
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
			name: "Non-empty string - no non-repeated characters",
			in:   "aeiaei",
			char: "",
		}, {
			name: "Non-empty string - 1 non-repeated characters",
			in:   "bbccd",
			char: "d",
		}, {
			name: "Non-empty string - 1 non-repeated characters",
			in:   "b",
			char: "b",
		}, {
			name: "Non-empty string - all non-repeated characters",
			in:   "1234",
			char: "1",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			c := firstNonRepeatedChar(tc.in)
			assert.Equal(t, tc.char, c)
		})
	}
}
