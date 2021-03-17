package main

// 1) Reverse a given string

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
