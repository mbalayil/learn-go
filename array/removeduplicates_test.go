package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicatesMap(t *testing.T) {
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
			out:  []int{},
		}, {
			name: "Non empty - no duplicates",
			in:   []int{2},
			out:  []int{2},
		}, {
			name: "Non empty - duplicates",
			in:   []int{2, 2},
			out:  []int{2},
		}, {
			name: "Non empty - duplicates",
			in:   []int{2, 3, 2, 5, 6, 5, 6},
			out:  []int{2, 3, 5, 6},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := removeDuplicatesMap(tc.in)

			assert.Equal(t, tc.out, v)
		})
	}
}
