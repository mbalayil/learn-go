package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMostFrequentWord(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		word string
	}{
		{
			name: "Empty slice",
			in:   []string{},
			word: "",
		}, {
			name: "Non-empty slice - no repeated words",
			in:   []string{"and", "a"},
			word: "a",
		}, {
			name: "Non-empty slice - 1 repeated word",
			in:   []string{"and", "bye", "and"},
			word: "and",
		}, {
			name: "Non-empty slice - 2 repeated words",
			in:   []string{"abc", "def", "abc", "def", "fgh"},
			word: "abc",
		}, {
			name: "Non-empty slice - 3 repeated words",
			in:   []string{"ghi", "def", "ghi", "abc", "abc", "def", "ghi"},
			word: "ghi",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			w := mostFrequentWord(tc.in)
			assert.Equal(t, tc.word, w)
		})
	}
}
