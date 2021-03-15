package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDuplicateNumbers(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		dns  []int
	}{
		{
			name: "empty list",
			a:    []int{},
			dns:  nil,
		}, {
			name: "non-empty list with duplicate entries",
			a:    []int{2, 3, 2, 5, 6, 3},
			dns:  []int{2, 3},
		}, {
			name: "non-empty list withiout duplicate entries",
			a:    []int{2, 3, 5},
			dns:  nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dns := findDuplicateNumbers(tc.a)
			assert.Equal(t, tc.dns, dns)
		})
	}
}
