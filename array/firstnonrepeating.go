package main

// Find the first non-repeating element in an array of integers

func firstNonRepeating(n []int) []int {
	l := len(n)
	if l == 0 {
		return nil
	}

	var result []int
	// A map "seen" is used to store the values that are once seen
	seen := make(map[int]bool)
	for i, v := range n {
		// For every value v in the array
		if _, ok := seen[v]; !ok {
			// If v is not an entry in the map, add it to the map
			seen[v] = true
			match := 0
			// Now that v is seen for the first time, see if it has a duplicate in the array
			for j := i + 1; j < l; j++ {
				if v == n[j] {
					// If v has a duplicate entry in the array, break from loop as we are no
					// longer interested in this value
					match = 1
					break
				}
			}
			if match == 0 {
				// If v does not have a duplicate entry in the array, it means this is the entry
				// that we want, add it to the result and return result
				result = append(result, n[i])
				return result
			}
		}
	}

	// If after traversing the whole array, you did not find any non repeating element,
	// result will have nil value. So return nil
	return result
}
