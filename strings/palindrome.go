package main

// Check if a string is palindrome.

func stringReverse(s string) string {
	l := len(s)
	// Convert to array of bytes
	r := []byte(s)
	i := 0
	j := l - 1
	for i < j {
		temp := r[i]
		r[i] = r[j]
		r[j] = temp
		i++
		j--
	}
	return string(r)
}

func checkPalindrome(s string) bool {
	sreverse := stringReverse(s)
	if sreverse == s {
		return true
	} else {
		return false
	}
}
