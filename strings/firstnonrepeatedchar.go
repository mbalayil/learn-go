package main

// Print the first non repeated character from a string

import (
	"strings"
)

// Function to count the number of times a character appears in a string
func countCharacter(s string, c string) int {
	count := strings.Count(s, c)
	if count == -1 {
		return 0
	}
	return count
}

func firstNonRepeatedChar(s string) string {
	l := len(s)
	for i := 0; i < l; i++ {
		if c := countCharacter(s, string(s[i])); c == 1 {
			return string(s[i])
		}
	}
	return ""
}
