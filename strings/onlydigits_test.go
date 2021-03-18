package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnlyDigits(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  bool
	}{
		{
			name: "Empty string",
			in:   "",
			out:  false,
		}, {
			name: "Non-empty string - only digits",
			in:   "12786",
			out:  true,
		}, {
			name: "Non-empty string - not only digits",
			in:   "12ab86",
			out:  false,
		}, {
			name: "Non-empty string - only digits",
			in:   "1",
			out:  true,
		}, {
			name: "Non-empty string - no digits",
			in:   "abcd",
			out:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out := onlyDigits(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}
