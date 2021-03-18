package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermute(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  []string
	}{
		{
			name: "nil string",
			in:   "",
			out:  nil,
		}, {
			name: "Non nil string",
			in:   "a",
			out:  []string{"a"},
		}, {
			name: "Non nil string",
			in:   "ab",
			out:  []string{"ab", "ba"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out := permute(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}
