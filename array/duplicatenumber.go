package main

// Find duplicate numbers in an Integer array.
// An array contains n numbers ranging from 0 to n-2. There is exactly one number
// repeated in the array. You need to write a program to find that duplicate number.

func findDuplicateNumber(n []int) []int {
	var duplicate []int
	l := len(n)
	if l == 0 {
		return duplicate
	}
	for i := range n {
		for j := i + 1; j < l; j++ {
			if n[i] == n[j] {
				duplicate = append(duplicate, n[i])
				return duplicate
			}
		}
	}
	return duplicate
}
