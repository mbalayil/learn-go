package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "Empty string",
			in:   "",
			out:  "",
		}, {
			name: "Non-empty string - one word - one element",
			in:   "a",
			out:  "a",
		}, {
			name: "Non-empty string - one word multiple elements",
			in:   "abcdefg",
			out:  "abcdefg",
		}, {
			name: "Non-empty string - multiple words",
			in:   "I need to reverse the words of this string",
			out:  "string this of words the reverse to need I",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out := reverseWords(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}
