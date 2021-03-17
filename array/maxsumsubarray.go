package main

// How to find a subarray with maximum sum in an array of positive and negative numbers ?

func maxSumSubarray(n []int) []int {
	l := len(n)
	if l == 0 {
		return nil
	}

	// Find positive contiguous subarrays (kadane's algorithm)
	// s stores the next index after a negetive prefix sum
	s := 0
	startIndex := 0
	endIndex := 0
	maxSoFar := 0
	currentSum := 0
	for i, v := range n {
		currentSum = currentSum + v
		if currentSum < 0 {
			// Because we are interested in positve contiguous subarray
			currentSum = 0
			s = i + 1
		} else {
			if currentSum > maxSoFar {
				maxSoFar = currentSum
				startIndex = s
				endIndex = i
			}
		}
	}

	var result []int
	for k := startIndex; k <= endIndex; k++ {
		result = append(result, n[k])
	}

	return result
}
