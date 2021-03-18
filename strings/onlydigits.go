package main

import (
	"regexp"
)

func onlyDigits(s string) bool {
	return regexp.MustCompile("^[0-9]+$").MatchString(s)
}
