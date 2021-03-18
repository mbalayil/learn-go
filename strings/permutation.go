package main

// Find all permutations of a string.

// Create a global array of strings
var List []string

func permute(s string) []string {
	// Convert string to byte array
	r := []byte(s)
	l := len(s)
	List = nil
	// Call the permutation function with the byte array, first and last index
	permutation(r, 0, l-1)
	return List
}

func permutation(s []byte, l, r int) {
	if l == r {
		List = append(List, string(s))
	} else {
		for i := l; i <= r; i++ {
			s = Doswap(s, l, i)
			permutation(s, l+1, r)
			s = Doswap(s, l, i)
		}
	}
}

// Function 'Doswap' swaps the values in s at i and j locations
func Doswap(s []byte, i, j int) []byte {
	temp := s[i]
	s[i] = s[j]
	s[j] = temp
	return s
}
