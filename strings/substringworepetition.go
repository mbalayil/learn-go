package main

// Find the length of the longest substring without repeating characters.

func substringWoRepeat(s string) string {
	l := len(s)
	if l == 0 {
		return ""
	}

	var maxStartIndex int
	var maxEndIndex int
	currStart := 0
	var currLength int
	maxLength := 0
	// A map is used to track if a value is repeated in every substring we are considering
	distinct := make(map[string]bool)
	i := currStart
	var update bool

	for ; i <= l; i++ {
		if i == l {
			// You have reached the end of array
			update = true
		} else {
			if _, ok := distinct[string(s[i])]; !ok {
				// Value not repeated.
				distinct[string(s[i])] = true
			} else {
				// Some value is repeated in the current substring.
				update = true
			}
		}

		if update {
			// You enter this block either when you are at the end of array or some value is
			// repeated in the current substring, so the max values are updated if required
			currLength = i - currStart
			if currLength > maxLength {
				// If the current substring length > already found maxLength, update maxLength
				maxStartIndex = currStart
				maxEndIndex = i - 1
				maxLength = currLength
			}
			if i == l {
				// If you are at the end of array, return the longest substring with no
				// repetition
				return string(s[maxStartIndex : maxEndIndex+1])
			} else {
				// Move to next substring.
				currStart = currStart + 1
				i = currStart
				// Clear the map entries for the new substring.
				for k := range distinct {
					delete(distinct, k)
				}
				// Add the first entry of the next substring
				distinct[string(s[i])] = true
				update = false
			}
		}

	}
	return string(s[maxStartIndex : maxEndIndex+1])
}
