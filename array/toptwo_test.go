package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopTwo(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		top  []int
	}{
		{
			name: "Nil slice",
			in:   nil,
			top:  nil,
		}, {
			name: "Empty slice",
			in:   []int{},
			top:  nil,
		}, {
			name: "Non-empty slice",
			in:   []int{8, 1, 3, 4, 9, 2},
			top:  []int{9, 8},
		}, {
			name: "Non-empty slice",
			in:   []int{1, 2},
			top:  []int{2, 1},
		}, {
			name: "Non-empty slice",
			in:   []int{1},
			top:  []int{1, 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			top := topTwo(tc.in)
			assert.Equal(t, tc.top, top)
		})
	}
}
