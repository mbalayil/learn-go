package main

// Reverse words in a given sentence
import (
	"strings"
)

func reverseWords(s string) string {
	list := strings.Split(s, " ")
	l := len(list)
	i := 0
	j := l - 1
	for i < j {
		temp := list[i]
		list[i] = list[j]
		list[j] = temp
		i++
		j--
	}
	reversed := strings.Join(list, " ")
	return reversed
}
