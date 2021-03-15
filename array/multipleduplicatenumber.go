package main

// Find repeated numbers in an array if it contains multiple duplicates.
// Suppose, an array contains n numbers ranging from 0 to n-1 and there are 5 duplicates on it, how do you find it ?

func findDuplicateNumbers(n []int) []int {
	var dns []int
	l := len(n)
	for i := range n {
		if !valueInSlice(dns, n[i]) {
			for j := i + 1; j < l; j++ {
				if n[i] == n[j] {
					dns = append(dns, n[i])
				}
			}
		}
	}
	return dns
}

// Function to check if a value is present in a slice
func valueInSlice(dns []int, value int) bool {
	for i := range dns {
		if dns[i] == value {
			return true
		}
	}
	return false
}
