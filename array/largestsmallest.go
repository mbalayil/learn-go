package main

// Find the largest and smallest number in an unsorted array.

func findSmallest(n []int) []int {
	var smallest []int
	if len(n) == 0 {
		return smallest
	}
	s := n[0]
	for i := range n {
		if n[i] < s {
			s = n[i]
		}
	}
	smallest = append(smallest, s)
	return smallest
}

func findLargest(n []int) []int {
	var largest []int
	if len(n) == 0 {
		return largest
	}
	l := n[0]
	for i := range n {
		if n[i] > l {
			l = n[i]
		}
	}
	largest = append(largest, l)
	return largest
}
