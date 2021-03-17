package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckPalindrome(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  bool
	}{
		{
			name: "nil string",
			in:   "",
			out:  true,
		}, {
			name: "Non nil string - not palindrome",
			in:   "munia balayil",
			out:  false,
		}, {
			name: "Non nil string - palindrome",
			in:   "malayalam",
			out:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out := checkPalindrome(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}
