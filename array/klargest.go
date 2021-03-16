package main

// Find kth largest element in an unsorted array.

func kLargest(n []int, k int) []int {
	l := len(n)
	// For nil and empty slice and when k > length of slice, we cannot find kth largest
	// element
	if l == 0 || k > l {
		return nil
	}

	// kth step of bubble sort puts the kth largest element in its position in the sorted
	// version. So repeat the bubble sort algorith for k steps only
	for i := 0; i < k; i++ {
		for j := 0; j < l-i-1; j++ {
			if n[j] > n[j+1] {
				n = swap(n, j, j+1)
			}
		}
	}

	// Return (l-k)th element
	out := []int{n[l-k]}
	return out
}

func swap(n []int, a int, b int) []int {
	temp := n[a]
	n[a] = n[b]
	n[b] = temp
	return n
}
