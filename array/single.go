package main

// There is an array with every element repeated twice except one. Find that element.

func singleMap(n []int) []int {
	// If the slice is nil or empty, no non-repeated elements.
	s := []int{}
	if len(n) == 0 {
		return s
	}

	// Create an empty map "singleton"
	singleton := make(map[int]bool)
	for _, v := range n {
		// For every v in n
		if _, ok := singleton[v]; ok {
			// If v is present in the map, the value has a duplicate, so we are not interested
			// in this value, so delete that entry from the map
			delete(singleton, v)
		} else {
			// If the value is not present in the map, add it to the map, as it may be a
			// unique element
			singleton[v] = true
		}
	}

	// After the full iteration through the array, we are left with all unique elements in
	// the array, with no duplicates at all
	for k := range singleton {
		s = append(s, k)
	}
	return s
}
