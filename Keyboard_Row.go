package main

import (
	"strings"
)

func findWords(words []string) []string {
	var ret []string
	row1 := []string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p"}
	row2 := []string{"a", "s", "d", "f", "g", "h", "j", "k", "l"}
	row3 := []string{"z", "x", "c", "v", "b", "n", "m"}

	for _, word := range words {
		chWord := strings.Split(word, "")
		var r1, r2, r3 int
		for _, c := range chWord {
			if contains(row1, c) {
				r1 = 1
			} else if contains(row2, c) {
				r2 = 1
			} else if contains(row3, c) {
				r3 = 1
			} else {

			}

		}

		if r1+r2+r3 == 1 {
			ret = append(ret, word)
		}
	}

	return ret
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
