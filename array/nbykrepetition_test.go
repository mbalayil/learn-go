package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepetition(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		n    int
		k    int
		out  []int
	}{
		{
			name: "Nil slice",
			in:   nil,
			n:    0,
			k:    3,
			out:  nil,
		}, {
			name: "Empty slice",
			in:   []int{},
			n:    0,
			k:    3,
			out:  nil,
		}, {
			name: "Non-empty slice",
			in:   []int{4, 2, 4, 4, 6, 1, 6, 3, 6},
			n:    10,
			k:    5,
			out:  []int{4, 6},
		}, {
			name: "Non-empty slice",
			in:   []int{4, 2, 4, 4, 3, 6, 1, 6, 3, 6},
			n:    12,
			k:    4,
			out:  nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out := repetition(tc.in, tc.n, tc.k)
			assert.Equal(t, tc.out, out)
		})
	}
}
