package main

// Given two sorted integer arrays A and B, merge B into A as one sorted array

func mergeSortedArrays(list1, list2 []int) []int {
	l1 := len(list1)
	l2 := len(list2)
	if l1 == 0 && l2 == 0 {
		// If both arrays are empty
		return nil
	}
	if l1 == 0 {
		// If list1 is empty
		return list2
	}
	if l2 == 0 {
		// If list2 is empty
		return list1
	}

	var result []int
	i := 0
	j := 0
	// If list1 and list2 are non-empty
	for i < l1 && j < l2 {
		if list1[i] == list2[j] {
			result = append(result, list1[i], list2[j])
			i++
			j++
		} else {
			if list1[i] < list2[j] {
				result = append(result, list1[i])
				i++
			} else {
				result = append(result, list2[j])
				j++
			}
		}
	}

	// If list1 is larger than list2, list1 will still have more values to process
	for ; i < l1; i++ {
		result = append(result, list1[i])
	}

	// If list2 is larger than list1, list2 will still have more values to process
	for ; j < l2; j++ {
		result = append(result, list2[j])
	}

	return result
}
