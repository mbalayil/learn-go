package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsetSum(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		out  int
	}{
		{
			name: "nil slice",
			in:   nil,
			out:  1,
		}, {
			name: "Empty slice",
			in:   []int{},
			out:  1,
		}, {
			name: "Non-mpty slice",
			in:   []int{1},
			out:  2,
		}, {
			name: "Non-mpty slice",
			in:   []int{1, 2, 4, 6},
			out:  14,
		}, {
			name: "Non-mpty slice",
			in:   []int{1, 1, 1, 5},
			out:  4,
		}, {
			name: "Non-mpty slice",
			in:   []int{5, 6, 7, 8},
			out:  1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out := subsetSum(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}
