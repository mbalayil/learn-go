package main

// Find the missing number in an integer array of 1 to 100.
// Given an integer array which contains numbers from 1 to 100 but one number is
// missing, you need to write a program to find that missing number in an array.

func findMissingNumber(values []int) []int {
	n := 1
	var missing []int
	// Assuming the array is sorted
	for i := range values {
		if values[i] == n {
			n++
		} else {
			missing = append(missing, n)
			return missing
		}
	}
	return missing
}
