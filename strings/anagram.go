package main

// Check if two strings are anagrams of each other

func anagram(s1, s2 string) bool {
	// Convert to array of bytes
	array1 := []byte(s1)
	array2 := []byte(s2)

	// A map 'distinct' keeps all the distinct characters of first string with its value
	// calculated as the number of times it is repeated in that string
	distinct := make(map[byte]int)
	for _, v := range array1 {
		if value, ok := distinct[v]; !ok {
			distinct[v] = 1
		} else {
			distinct[v] = value + 1
		}
	}

	for _, v := range array2 {
		// For every character in the second string
		if value, ok := distinct[v]; !ok {
			// If it is not present in the map, its a new character not present in the first
			// string, so not an anagram, return false
			return false
		} else {
			// If character present in map
			if value-1 < 0 {
				// And if value - 1 < 0, it means the character is repeated more number of times
				// in the second string that in the first, so not an anagram, return false
				return false
			}
			// And if value - 1 >= 0, decrement its value
			distinct[v] = value - 1
		}
	}

	// Finally, if all the keys in the 'distinct' map has its corresponding key as 0, then
	// 2 strings are anagrams of each other, else not anagrams of each other
	for _, v := range distinct {
		if v != 0 {
			return false
		}
	}

	return true
}
