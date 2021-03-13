package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMissingNumber(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		want []int
	}{
		{
			name: "Non-empty slice with missing number",
			a:    []int{1, 2, 3, 5, 6, 7, 8, 9, 10},
			want: []int{4},
		}, {
			name: "Empty slice",
			a:    []int{},
			want: nil,
		}, {
			name: "Non-empty slice with no missing number",
			a:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			x := findMissingNumber(tc.a)
			assert.Equal(t, tc.want, x)
		})
	}
}
