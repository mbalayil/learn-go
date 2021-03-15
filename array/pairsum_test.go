package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPairSum(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		s    int
		p    []int
	}{
		{
			name: "empty array",
			a:    []int{},
			s:    6,
			p:    nil,
		}, {
			name: "non empty array - no such pairs",
			a:    []int{1, 2, 3, 4, 5},
			s:    10,
			p:    nil,
		}, {
			name: "non-empty array with >= 2 pairs with given sum",
			a:    []int{2, 3, 4, 5, 6},
			s:    7,
			p:    []int{2, 5, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := pairsum(tt.a, tt.s)
			assert.Equal(t, p, tt.p)
		})
	}
}
