package main

import (
	"sort"
	"strings"
)

func main() {

	a := "FEDCA"
	b := "ECD"

	a = SortString(a)
	b = SortString(b)
	// true: a is contain b, false: a is not contain b
	println(StringContain(a, b))
}

func StringContain(a string, b string) bool {
	var pa, pb int = 0, 0
	for pb < len(b) {
		for pa < len(a) && string(a[pa]) < string(b[pb]) {
			pa++
		}
		if pa >= len(a) || string(a[pa]) > string(b[pb]) {
			return false
		}
		pb++
	}
	return true
}

func SortString(s string) string {
	strs := strings.Split(s, "")
	sort.Strings(strs)
	return strings.Join(strs, "")
}
