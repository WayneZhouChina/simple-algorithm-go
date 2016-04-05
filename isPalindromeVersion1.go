package main

import (
	"fmt"
)

func main() {
	s := []byte("acabada")
	fmt.Println(isPalindrome(s, len(s)))
}

func isPalindrome(s []byte, n int) bool {
	if n < 1 {
		return false
	}

	front := 0
	back := n - 1

	for front < back {
		if s[front] != s[back] {
			return false
		}
		front++
		back--
	}
	return true
}
