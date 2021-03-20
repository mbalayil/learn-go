package main

// Find the longest palindromic substring in a string

func longestPalindromSubstring(s string) (int, string) {
	l := len(s)
	if l == 0 {
		return 0, ""
	}
	// A table array stores false at table[i][j] if the substring from index i to j is not
	// a palindrome and true otherwise.
	table := make([][]bool, l)
	for i := range table {
		table[i] = make([]bool, l)
	}

	// Substring of length 1 is always a palindrome.
	maxLength := 1
	for i := range table {
		table[i][i] = true
	}

	start := 0
	// Substring of length 2.
	for i := 0; i < l-1; i++ {
		if s[i] == s[i+1] {
			table[i][i+1] = true
			start = i
			maxLength = 2
		}
	}

	// Substring of higher length
	// Let k keep the substring length
	for k := 3; k <= l; k++ {
		for i := 0; i < l-k+1; i++ {
			// Get the ending index.
			j := i + k - 1
			// Checking for sub-string from ith index to jth index iff s[i + 1] to s[(j-1)] is
			// a palindrome.
			if table[i+1][j-1] && s[i] == s[j] {
				table[i][j] = true
				if k > maxLength {
					start = i
					maxLength = k
				}
			}
		}
	}

	// fmt.Printf("%#v\n", table)
	return maxLength, string(s[start : start+maxLength])
}
