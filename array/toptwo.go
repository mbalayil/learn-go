package main

// Find the top two numbers from an integer array

func topTwo(n []int) []int {
	l := len(n)
	if l == 0 {
		return nil
	}

	top1 := n[0]
	top2 := n[0]
	for _, v := range n {
		if v > top1 {
			top2 = top1
			top1 = v
		}
	}

	top := []int{top1, top2}
	return top
}
