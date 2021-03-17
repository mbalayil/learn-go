package main

// Reverse a given string - using recursion.

func stringReverseRecursion(s string) string {
	l := len(s)
	sByteArray := []byte(s)
	reverse := sreverse(sByteArray, 0, l-1)
	return string(reverse)
}

func sreverse(s []byte, i, j int) string {
	if i >= j {
		return string(s)
	} else {
		temp := s[i]
		s[i] = s[j]
		s[j] = temp
		return sreverse(s, i+1, j-1)
	}
}
