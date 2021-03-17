package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDuplicateCharacters(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "nil string",
			in:   "",
			out:  "",
		}, {
			name: "Non nil string",
			in:   "munia balayil",
			out:  "ail",
		}, {
			name: "Non nil string",
			in:   "munia",
			out:  "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out := duplicateCharacters(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}
