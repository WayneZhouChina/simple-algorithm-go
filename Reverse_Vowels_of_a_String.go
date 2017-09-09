package main

import (
	"strings"
)

func reverseVowels(s string) string {
	sStart := 0
	sLength := len(s)

	runeString := []rune(s)
	sEnd := sLength - 1

	for sStart < sEnd {
		if isVowels(string(s[sStart])) == true && isVowels(string(s[sEnd])) == true {
			runeString[sStart], runeString[sEnd] = runeString[sEnd], runeString[sStart]
			sStart++
			sEnd--
		} else if isVowels(string(s[sStart])) == true {
			sEnd--
		} else {
			sStart++
		}
	}
	return string(runeString)
}

func isVowels(c string) bool {
	c = strings.ToLower(c)
	if c == "a" || c == "e" || c == "i" || c == "o" || c == "u" {
		return true
	}
	return false
}
