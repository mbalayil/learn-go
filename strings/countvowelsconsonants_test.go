package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountVowelsConsonants(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		vCount int
		cCount int
	}{
		{
			name:   "Empty string",
			in:     "",
			vCount: 0,
			cCount: 0,
		}, {
			name:   "Non-empty string - only vowels",
			in:     "aei",
			vCount: 3,
			cCount: 0,
		}, {
			name:   "Non-empty string - only consonants",
			in:     "bcd",
			vCount: 0,
			cCount: 3,
		}, {
			name:   "Non-empty string - vowels and consonants",
			in:     "abcdefghi",
			vCount: 3,
			cCount: 6,
		}, {
			name:   "Non-empty string - only digits",
			in:     "12345",
			vCount: 0,
			cCount: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v, c := countVowelsConsonants(tc.in)
			assert.Equal(t, tc.vCount, v)
			assert.Equal(t, tc.cCount, c)
		})
	}
}
