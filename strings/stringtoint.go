package main

// Convert a given string into int

import "strconv"

func stringToInt(s string) (int, error) {
	return (strconv.Atoi(s))
}
