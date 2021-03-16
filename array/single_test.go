package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleMap(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		s    []int
	}{
		{
			name: "nil list",
			a:    nil,
			s:    []int{},
		}, {
			name: "Empty list",
			a:    []int{},
			s:    []int{},
		}, {
			name: "Non-empty list with singleton element",
			a:    []int{5},
			s:    []int{5},
		}, {
			name: "Non-empty list without singleton element",
			a:    []int{5, 5},
			s:    []int{},
		}, {
			name: "Non-empty list with multiple singleton elements",
			a:    []int{5, 6},
			s:    []int{5, 6},
		}, {
			name: "Non-empty list without singleton element",
			a:    []int{9, 2, 3, 9, 2, 3},
			s:    []int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := singleMap(tc.a)
			assert.Equal(t, tc.s, s)
		})
	}
}
