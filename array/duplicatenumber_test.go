package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDuplicateNumber(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		want []int
	}{
		{
			name: "Nil slice",
			a:    nil,
			want: nil,
		}, {
			name: "Empty slice",
			a:    []int{},
			want: nil,
		}, {
			name: "Non-empty slice with repetition",
			a:    []int{0, 1, 2, 3, 4, 5, 5, 6, 7, 8},
			want: []int{5},
		}, {
			name: "Non-empty slice with repetition",
			a:    []int{2, 0, 4, 1, 7, 8, 3, 5, 6, 4},
			want: []int{4},
		}, {
			name: "Non-empty slice without repetition",
			a:    []int{0, 2, 1, 4, 3, 6, 5, 7, 8, 9},
			want: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			x := findDuplicateNumber(tc.a)
			assert.Equal(t, tc.want, x)
		})
	}
}
