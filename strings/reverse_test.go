package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringReverse(t *testing.T) {
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
			out:  "liyalab ainum",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out := stringReverse(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}
