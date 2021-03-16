package main

//  Rearrange an array in alternating positive and negative numbers maintaining the
//  order of appearance. Number of positive and negative numbers need not be equal. If
//  there are more positive numbers they appear at the end of the array. If there are
//  more negative numbers, they too appear at the end of the array.

func alternatePositiveNegetive(n []int) []int {
	l := len(n)
	if l == 0 {
		return nil
	}

	start := n[0]
	if start > 0 {
		// First number is positive
		for i := range n {
			if i%2 == 0 {
				// Values at even indices must be positive.
				if n[i] < 0 {
					// If they are negetive,
					for j := i + 1; j < l; j++ {
						if n[j] > 0 {
							// Note the index of the immediate forthcoming positive value - j.
							n = copyAndShift(n, i, j)
							break
						}
					}
				}
			} else {
				// Values at odd indices must be negetive.
				if n[i] > 0 {
					// If they are positive,
					for j := i + 1; j < l; j++ {
						if n[j] < 0 {
							// Note the index of the immediate forthcoming negetive value - j.
							n = copyAndShift(n, i, j)
							break
						}
					}
				}
			}
		}

	} else {
		// First number is negetive
		for i := range n {
			if i%2 == 0 {
				// Values at even indices must be negetive.
				if n[i] > 0 {
					// If they are positive,
					for j := i + 1; j < l; j++ {
						if n[j] < 0 {
							// Note the index of the immediate forthcoming negetive value - j.
							n = copyAndShift(n, i, j)
							break
						}
					}
				}
			} else {
				// Values at odd indices must be positive.
				if n[i] < 0 {
					// If they are negetive,
					for j := i + 1; j < l; j++ {
						if n[j] > 0 {
							// Note the index of the immediate forthcoming positive value - j.
							n = copyAndShift(n, i, j)
							break
						}
					}
				}
			}
		}
	}
	return n
}

// Function will shift n[i+1] to n[j] to the right and make n[i+1] = n[j].
func copyAndShift(n []int, i, j int) []int {
	temp := n[i]
	n[i] = n[j]
	for k := j; k > i+1; k-- {
		n[k] = n[k-1]
	}
	n[i+1] = temp
	return n
}
