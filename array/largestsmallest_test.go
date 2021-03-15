package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSmallest(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		out  []int
	}{
		{
			name: "Nil slice",
			a:    nil,
			out:  nil,
		}, {
			name: "empty slice",
			a:    []int{},
			out:  nil,
		}, {
			name: "non empty slice - exists",
			a:    []int{2, 3, 9, 10, 20},
			out:  []int{2},
		}, {
			name: "non empty slice - exists",
			a:    []int{10, 15, 7, 5, 9},
			out:  []int{5},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			x := findSmallest(tc.a)
			assert.Equal(t, tc.out, x)
		})
	}
}

func TestFindLargest(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		out  []int
	}{
		{
			name: "Nil slice",
			a:    nil,
			out:  nil,
		}, {
			name: "empty slice",
			a:    []int{},
			out:  nil,
		}, {
			name: "non empty slice - exists",
			a:    []int{2, 3, 9, 10, 20},
			out:  []int{20},
		}, {
			name: "non empty slice - exists",
			a:    []int{10, 15, 7, 5, 9},
			out:  []int{15},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			x := findLargest(tc.a)
			assert.Equal(t, tc.out, x)
		})
	}
}
