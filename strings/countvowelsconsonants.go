package main

// Count the number of vowels and consonants in a given string

import (
	"regexp"
	"strings"
)

func countVowelsConsonants(s string) (int, int) {
	vCount := 0
	cCount := 0
	l := len(s)
	s = strings.ToLower(s)
	for i := 0; i < l; i++ {
		if regexp.MustCompile("[aeiou]").MatchString(string(s[i])) {
			vCount++
		} else {
			if regexp.MustCompile("[^aeiou0-9 ]").MatchString(string(s[i])) {
				cCount++
			}
		}
	}

	return vCount, cCount
}
