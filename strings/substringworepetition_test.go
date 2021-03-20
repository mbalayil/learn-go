package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubstringWoRepeat(t *testing.T) {
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
			name: "Non-empty string - one element",
			in:   "a",
			out:  "a",
		}, {
			name: "Non-empty string - no repeated elements",
			in:   "abcdefg",
			out:  "abcdefg",
		}, {
			name: "Non-empty string - with repeated elements",
			in:   "abcadeafg",
			out:  "bcade",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out := substringWoRepeat(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}
