package main

// Check if two strings are a rotation of each other

import "strings"

func rotation(string1, string2 string) bool {
	// Concatenate string1 to string1
	temp := []string{string1, string1}
	joined := strings.Join(temp, "")

	if strings.Contains(joined, string2) {
		// If string2 is a substring of the concatenated string, it is a rotation of string1
		return true
	} else {
		return false
	}
}
