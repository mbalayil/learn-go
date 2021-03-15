package main

// Check if an array contains a number.

func search(n []int, value int) bool {
	for i := range n {
		if value == n[i] {
			return true
		}
	}
	return false
}
