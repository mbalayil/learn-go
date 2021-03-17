package main

// Write a program to find the length of the longest consecutive sequence in an array of
// integers ?

func longConsecutiveSequence(n []int) (int, []int) {
	l := len(n)
	if l == 0 {
		return 0, nil
	}

	if l == 1 {
		result := []int{0, 0}
		return 1, result
	}

	// s is used to keep the sequence length of the current sequence
	s := 0
	// sequence/length keeps the longest sequence length obtained so far
	sequenceLength := 0
	// start end end is the starting and ending index of the current sequence
	start := 0
	end := 0
	// startIndex & endIndex is the starting and ending index of the longest sequence obtained so far
	startIndex := 0
	endIndex := 0
	// value is the current value under consideration
	value := 0

	for i, v := range n {
		if i == start {
			// You are at the start of a sequence
			value = v
			s = 1
		} else {
			// You are in the middle of a sequence
			if v == value+1 {
				end = i
				s = s + 1
				value = v
			} else {
				// You are at the end of a sequence, update the required fields if the current
				// sequence is longer than the previously recorded maximum length
				if s > sequenceLength {
					sequenceLength = s
					startIndex = start
					endIndex = end
				}
				// And move to the next sequence
				start = i
				value = v
				s = 1
			}
		}
	}

	result := []int{startIndex, endIndex}

	// Return the sequence length of the longest sequence and the start and end index of
	// the sequence
	return sequenceLength, result
}
