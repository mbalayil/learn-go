package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		v    int
		want bool
	}{
		{
			name: "Nil slice",
			a:    nil,
			v:    0,
			want: false,
		}, {
			name: "Empty slice",
			a:    []int{},
			v:    1,
			want: false,
		}, {
			name: "Non-empty slice with search value not present",
			a:    []int{1, 2, 3, 4, 5, 6},
			v:    10,
			want: false,
		}, {
			name: "Non-empty slice with search value present",
			a:    []int{5, 10, 3, 7, 8, 8},
			v:    8,
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			x := search(tc.a, tc.v)
			assert.Equal(t, tc.want, x)
		})
	}
}
