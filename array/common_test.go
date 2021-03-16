package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommon(t *testing.T) {
	tests := []struct {
		name string
		in1  []int
		in2  []int
		in3  []int
		out  []int
	}{
		{
			name: "Nil slices",
		}, {
			name: "2 empty slices",
			in1:  []int{},
			in2:  []int{},
			in3:  []int{4, 5, 8, 9},
		}, {
			name: "3 non-empty strings with no common elements",
			in1:  []int{1, 2, 3},
			in2:  []int{6, 7},
			in3:  []int{4, 5, 8, 9},
		}, {
			name: "3 non-empty strings with some common elements",
			in1:  []int{1, 2, 3, 4, 8},
			in2:  []int{4, 6, 7, 8},
			in3:  []int{4, 5, 8, 9},
			out:  []int{4, 8},
		}, {
			name: "Same array",
			in1:  []int{1, 2, 3},
			in2:  []int{1, 2, 3},
			in3:  []int{1, 2, 3},
			out:  []int{1, 2, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out := common(tc.in1, tc.in2, tc.in3)
			assert.Equal(t, tc.out, out)
		})
	}
}
