package main

// Find maximum occuring character in a string

import (
	"strings"
)

func maxOccuringChar(s string) string {
	// A map 'count' keeps each character of the string as key with a value that indicates the number of times they occur in the string
	count := make(map[string]int)
	l := len(s)
	for i := 0; i < l; i++ {
		if _, ok := count[string(s[i])]; !ok {
			c := strings.Count(s, string(s[i]))
			count[string(s[i])] = c
		}
	}

	// Find the character with maximum count
	maxOcc := -1
	var maxChar string
	for k, v := range count {
		if v > maxOcc {
			maxOcc = v
			maxChar = k
		}
	}

	return maxChar
}
