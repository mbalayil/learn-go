package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSortedArrays(t *testing.T) {
	tests := []struct {
		name   string
		list1  []int
		list2  []int
		result []int
	}{
		{
			name:   "Both inputs nil slices",
			list1:  nil,
			list2:  nil,
			result: nil,
		}, {
			name:   "One input empty slice, other nil slice",
			list1:  []int{},
			list2:  nil,
			result: nil,
		}, {
			name:   "Both non-empty slices",
			list1:  []int{1, 5, 8, 9, 10},
			list2:  []int{2, 3, 4, 6, 7},
			result: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		}, {
			name:   "one non-empty slice, other empty slice",
			list1:  []int{1, 5, 8, 9, 10},
			list2:  []int{},
			result: []int{1, 5, 8, 9, 10},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := mergeSortedArrays(tc.list1, tc.list2)
			assert.Equal(t, tc.result, result)
		})
	}
}
