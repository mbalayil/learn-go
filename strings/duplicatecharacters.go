package main

// Print duplicate characters from a string.

func duplicateCharacters(s string) string {
	// Convert to array of bytes
	r := []byte(s)

	// Byte array 'duplicates' keeps the byte values of all the duplicate characters
	var duplicates []byte

	// A map 'distinct' keeps the byte value of all the distinct characters of the string
	distinct := make(map[byte]bool)

	for _, v := range r {
		// For every character in the string

		if value, ok := distinct[v]; !ok {
			// If it is not present in the map, add the character with value set to true
			distinct[v] = true
		} else {
			// If the character is already present in the map
			if value {
				// And the value of the character is true, it means this is the first
				// repetition, add it to the list of duplicate characters
				duplicates = append(duplicates, v)

				// Set the value of the character to false since we are no longer interested in
				// this character
				distinct[v] = false
			}
		}
	}

	return string(duplicates)
}
