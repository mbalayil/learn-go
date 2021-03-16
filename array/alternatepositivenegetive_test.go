package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlternatePositiveNegetive(t *testing.T) {
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
			name: "Non-empty slice - positive start",
			in:   []int{8, 5, -3, -2, 6, 7},
			out:  []int{8, -3, 5, -2, 6, 7},
		}, {
			name: "Non-empty slice - negetive start",
			in:   []int{-5, -3, -2, 6, 7},
			out:  []int{-5, 6, -3, 7, -2},
		}, {
			name: "Non-empty slice- more positive numbers",
			in:   []int{9, -5, -3, -2, 6, 7, 10, 11, 12},
			out:  []int{9, -5, 6, -3, 7, -2, 10, 11, 12},
		}, {
			name: "Non-empty slice- more negetive numbers",
			in:   []int{-1, -4, 9, -9, -5, -3, -2, 6, 7},
			out:  []int{-1, 9, -4, 6, -9, 7, -5, -3, -2},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out := alternatePositiveNegetive(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}
