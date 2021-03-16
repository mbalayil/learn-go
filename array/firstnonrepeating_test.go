package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstNonRepeating(t *testing.T) {
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
			name: "Non-empty slice with non repeating element",
			in:   []int{2, 3, 2, 3, 4, 3, 7},
			out:  []int{4},
		}, {
			name: "Non-empty slice without non repeating element",
			in:   []int{2, 2, 2, 2, 2},
			out:  nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out := firstNonRepeating(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}
