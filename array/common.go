package main

// Find common elements in three sorted arrays.

func common(in1, in2, in3 []int) []int {
	// return empty slice if any of the input slice is empty.
	if len(in1) == 0 || len(in2) == 0 || len(in3) == 0 {
		return nil
	}

	// Find common elements in first 2 slices, store them
	// in a map called common2.
	var common []int
	common2 := make(map[int]bool)
	i := 0
	j := 0
	for i < len(in1) && j < len(in2) {
		if in1[i] == in2[j] {
			common2[in1[i]] = true
			i++
			j++
		} else {
			if in1[i] < in2[j] {
				i++
			} else {
				j++
			}
		}
	}

	// Find common elements in 3rd slice and common2,
	// store in another slice common.
	for _, v := range in3 {
		if _, ok := common2[v]; ok {
			common = append(common, v)
		}
	}

	return common
}
