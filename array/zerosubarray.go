package main

// You are given an array of positive and negative numbers, find if there is  a
// non-empty sub-array with 0 sum.

func zeroSubarray(n []int) (bool, []int) {
	l := len(n)
	if l == 0 {
		return false, nil
	}

	// Create a map to keep all the prefix sums and their corresponding index.
	// Prefix sum at index i = sum of all the values from 0th index to ith index
	prefixSum := make(map[int]int)
	prefixSum[0] = -1

	sum := 0
	for i, v := range n {
		sum = sum + v
		if _, ok := prefixSum[sum]; !ok {
			// If entry corresponding to a sum is not present in the map, add that sum to the
			// map with its value set to the index upto which the sum is taken
			prefixSum[sum] = i
		} else {
			// If entry corresponding to a sum is present in the map, it means the sum is
			// repeated, and hence all values from its previous occurence till now sums to
			// zero, note the starting and ending index of this slot
			var indices []int
			startIndex := prefixSum[sum] + 1
			EndIndex := i
			indices = append(indices, startIndex, EndIndex)
			return true, indices
		}
	}

	// If all the prefix sums are unique it means no subarray sums to zero
	return false, nil
}
