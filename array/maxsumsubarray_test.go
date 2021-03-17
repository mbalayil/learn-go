package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSumSubarray(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		out  []int
	}{
		{
			name: "Nil slice",
			in:   nil,
			out:  nil,
		}, {
			name: "Empty slice",
			in:   []int{},
			out:  nil,
		}, {
			name: "Non-empty slice",
			in:   []int{-2, -3, 4, -1, -2, 1, 5, -3},
			out:  []int{4, -1, -2, 1, 5},
		}, {
			name: "Non-empty slice",
			in:   []int{1},
			out:  []int{1},
		}, {
			name: "Non-empty slice",
			in:   []int{1, 2, 3, -2, -3, -4},
			out:  []int{1, 2, 3},
		}, {
			name: "Non-empty slice",
			in:   []int{1, -2, 1, 6, -4, -5, 20, 9},
			out:  []int{20, 9},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out := maxSumSubarray(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}
