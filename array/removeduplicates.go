package main

// Write a program to remove duplicates from an array.

func removeDuplicatesMap(n []int) []int {
	if len(n) == 0 {
		return n
	}

	var data []int
	seen := make(map[int]bool)

	for _, v := range n {
		if _, ok := seen[v]; !ok {
			seen[v] = true
			data = append(data, v)
		}
	}

	return data
}
