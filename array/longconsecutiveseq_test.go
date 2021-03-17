package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongConsecutiveSequence(t *testing.T) {
	tests := []struct {
		name    string
		in      []int
		length  int
		indices []int
	}{
		{
			name:    "Nil slice",
			in:      nil,
			length:  0,
			indices: nil,
		}, {
			name:    "Empty slice",
			in:      []int{},
			length:  0,
			indices: nil,
		}, {
			name:    "Non-empty slice",
			in:      []int{1, 2, 3, 6, 7, 9, 10, 11, 12, 13, 19},
			length:  5,
			indices: []int{5, 9},
		}, {
			name:    "Non-empty slice",
			in:      []int{1},
			length:  1,
			indices: []int{0, 0},
		}, {
			name:    "Non-empty slice",
			in:      []int{8, 1, 2, 3, 6},
			length:  3,
			indices: []int{1, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			length, indices := longConsecutiveSequence(tc.in)
			assert.Equal(t, tc.length, length)
			assert.Equal(t, tc.indices, indices)
		})
	}

}
