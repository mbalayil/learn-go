package main

// Find the smallest positive integer value that cannot be represented as a sum of
// elements of any subset of the given set (given a sorted array (sorted in
// non-decreasing order) of positive numbers) in o(n) complexity.

func subsetSum(n []int) int {
	result := 1
	l := len(n)
	if l == 0 {
		return result
	}

	// If n[0...i-1] can represent values 1 to result-1, including n[i] into the set, the
	// set can represent values from 1 to result-1+n[i]
	for i := 0; i < l && n[i] <= result; i++ {
		result = result + n[i]
	}
	return result
}
