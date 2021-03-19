package main

// Given an array of strings, find the most frequent word in a given array. In the case
// of a tie, the string that is the smallest (lexicographically) is printed.

func mostFrequentWord(list []string) string {
	// A map 'count' keeps each word in the array as its key and the number of times
	// it appears in the array as its value
	count := make(map[string]int)
	l := len(list)
	for i := 0; i < l; i++ {
		if v, ok := count[list[i]]; ok {
			count[list[i]] = v + 1
		} else {
			count[list[i]] = 1
		}
	}

	// Find the frequency in the map that has the maximum value and the corresponding word
	maxFrequency := 0
	var mostFreqWord string
	for k, v := range count {
		if v > maxFrequency {
			maxFrequency = v
			mostFreqWord = k
		}
	}

	// Find all the words whose frequency in the array = maxFrequency other than the
	// word found before, if any word is lexicographically smaller, the mostFreqWord
	// is overwritten by that word. Repeat for all the other words with frequency =
	// maxFrequency
	for k, v := range count {
		if v == maxFrequency && k != mostFreqWord {
			if k < mostFreqWord {
				mostFreqWord = k
			}
		}
	}

	// Return the word that is lexicographically smallest among all the words with
	// frequency = maxFrequency
	return mostFreqWord
}
