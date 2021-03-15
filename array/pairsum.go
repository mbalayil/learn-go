package main

// Find all pairs on an integer array whose sum is equal to a given number.
// You need to write a program to find all elements in the array whose sum is equal to the given number. Remember, the array may contain both positive and negative numbers, so your solution should consider that.

func pairsum(n []int, s int) []int {
	var p []int
	l := len(n)
	for i := range n {
		for j := i + 1; j < l; j++ {
			if n[i]+n[j] == s {
				p = append(p, n[i], n[j])
			}
		}
	}
	return p
}
