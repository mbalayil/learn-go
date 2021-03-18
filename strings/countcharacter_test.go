package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountCharacter(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		char   string
		cCount int
	}{
		{
			name:   "Empty string",
			in:     "",
			char:   "",
			cCount: 1,
		}, {
			name:   "Non-empty string - character not present",
			in:     "aei",
			char:   "b",
			cCount: 0,
		}, {
			name:   "Non-empty string - character present once",
			in:     "bcd",
			char:   "b",
			cCount: 1,
		}, {
			name:   "Non-empty string - character present more than once",
			in:     "abbcdb",
			char:   "b",
			cCount: 3,
		}, {
			name:   "Non-empty string - character present once",
			in:     "1",
			char:   "1",
			cCount: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			c := countCharacter(tc.in, tc.char)
			assert.Equal(t, tc.cCount, c)
		})
	}
}
