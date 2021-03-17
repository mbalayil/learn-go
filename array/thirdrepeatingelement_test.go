package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThirdRepeatingElement(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		out  []int
	}{
		{
			name: "Nil slice",
			in:   nil,
			out:  nil,
		}, {
			name: "Empty slice",
			in:   []int{},
			out:  nil,
		}, {
			name: "Non-empty slice - >= 3 repeated elements",
			in:   []int{7, 6, 2, 3, 2, 3, 2, 3, 6, 6},
			out:  []int{6},
		}, {
			name: "Non-empty slice - < 3 repeated elements",
			in:   []int{7},
			out:  nil,
		}, {
			name: "Non-empty slice -  all elements are same",
			in:   []int{2, 2, 2, 2, 2, 2},
			out:  nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out := thirdRepeatingElement(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}
