package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKLargest(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		k    int
		out  []int
	}{
		{
			name: "Nil slice",
			a:    nil,
			k:    2,
			out:  nil,
		}, {
			name: "Empty slice",
			a:    []int{},
			k:    3,
			out:  nil,
		}, {
			name: "Non-empty slice",
			a:    []int{5, 8, 9, 6, 4, 3, 6},
			k:    2,
			out:  []int{8},
		}, {
			name: "Non-empty slice - k > size of slice",
			a:    []int{5},
			k:    2,
			out:  nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out := kLargest(tc.a, tc.k)
			assert.Equal(t, tc.out, out)
		})
	}
}
