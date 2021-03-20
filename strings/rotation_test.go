package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotation(t *testing.T) {
	tests := []struct {
		name    string
		string1 string
		string2 string
		out     bool
	}{
		{
			name:    "Empty string",
			string1: "",
			string2: "",
			out:     true,
		}, {
			name:    "Non-empty string - one element",
			string1: "a",
			string2: "a",
			out:     true,
		}, {
			name:    "Non-empty string - not a rotation",
			string1: "abcd",
			string2: "cdba",
			out:     false,
		}, {
			name:    "Non-empty string - is a rotation",
			string1: "abcd",
			string2: "cdab",
			out:     true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out := rotation(tc.string1, tc.string2)
			assert.Equal(t, tc.out, out)
		})
	}
}
