package main

// You are given an array of size n, find all elements in the array that appear more than n/k times.

// The function takes 3 arguments - array, its size and a value k
func repetition(list []int, n, k int) []int {

	if n == 0 {
		return nil
	}

	times := n / k
	// A map "repeat" keeps every element in the array as keys with its number of
	// repetition in the array as value
	repeat := make(map[int]int)
	for _, v := range list {
		// For every value v in the array
		if r, ok := repeat[v]; ok {
			// If it is present in the map, add 1 to its previous value
			repeat[v] = r + 1
		} else {
			// If not, add the element as a new entry with its value initialized as 1
			repeat[v] = 1
		}
	}

	var result []int
	// Store all the keys in the map with value > n/k in result
	for k, v := range repeat {
		if v > times {
			result = append(result, k)
		}
	}

	// Sort the result array
	for i := range result {
		for j := 0; j < len(result)-1-i; j++ {
			if result[j] > result[j+1] {
				temp := result[j]
				result[j] = result[j+1]
				result[j+1] = temp
			}
		}
	}

	return result
}
