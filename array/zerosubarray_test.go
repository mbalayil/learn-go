package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeroSubarray(t *testing.T) {
	tests := []struct {
		name    string
		in      []int
		answer  bool
		indices []int
	}{
		{
			name:    "Nil slice",
			in:      nil,
			answer:  false,
			indices: nil,
		}, {
			name:    "Empty slice",
			in:      []int{},
			answer:  false,
			indices: nil,
		}, {
			name:    "Non-empty slice-with no subarray whose sum is 0",
			in:      []int{1, 2, 3, 4, 5, 6, -1},
			answer:  false,
			indices: nil,
		}, {
			name:    "Non-empty slice-with a subarray whose sum is 0",
			in:      []int{1, 4, -2, -2, 5, -4, 3},
			answer:  true,
			indices: []int{1, 3},
		}, {
			name:    "Non-empty slice-with a subarray whose sum is 0",
			in:      []int{5, -5},
			answer:  true,
			indices: []int{0, 1},
		}, {
			name:    "Non-empty slice-with a subarray whose sum is 0",
			in:      []int{-5, 1, 2, 2},
			answer:  true,
			indices: []int{0, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			answer, indices := zeroSubarray(tc.in)
			assert.Equal(t, tc.answer, answer)
			assert.Equal(t, tc.indices, indices)
		})
	}
}
