package main

// Count the occurence of a given character in a string

import (
	"strings"
)

func countCharacter(s string, c string) int {
	count := strings.Count(s, c)
	if count == -1 {
		return 0
	}
	return count
}
