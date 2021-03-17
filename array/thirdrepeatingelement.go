package main

// Given an array of integers, find the third repeating element in it.

func thirdRepeatingElement(n []int) []int {
	// Map 'repeated' keeps all distinct entries of the slice with their value set to
	// false if they are non-repeating and true if they are repeating.
	repeated := make(map[int]bool)

	count := 0
	var result []int

	for _, v := range n {
		if _, ok := repeated[v]; !ok {
			// New entry
			repeated[v] = false
		} else {
			// Entry already present in map
			if !repeated[v] {
				// First repetition
				count++
				repeated[v] = true
			}
		}

		if count == 3 {
			result = append(result, v)
			return result
		}
	}
	return result
}
