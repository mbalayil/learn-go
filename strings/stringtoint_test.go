package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToInt(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  int
		err  string
	}{
		{
			name: "Empty string",
			in:   "",
			err:  "strconv.Atoi: parsing \"\": invalid syntax",
		}, {
			name: "Non-empty string - one element - invalid entry",
			in:   "a",
			err:  "strconv.Atoi: parsing \"a\": invalid syntax",
		}, {
			name: "Non-empty string - 2 elements - invalid entry",
			in:   "ab",
			err:  "strconv.Atoi: parsing \"ab\": invalid syntax",
		}, {
			name: "Non-empty string - valid entry",
			in:   "23",
			out:  23,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out, err := stringToInt(tc.in)
			assert.Equal(t, tc.out, out)
			if tc.err != "" {
				assert.EqualError(t, err, tc.err)
			}
		})
	}
}
