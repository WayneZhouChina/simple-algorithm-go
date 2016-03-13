package main

import "fmt"

func main() {
	s := []byte("helloString")
	newStr := LeftRotateString(s, len(s), 4)
	fmt.Println(string(newStr))
}

func ReverseString(s []byte, start int, end int) {
	for start < end {
		temp := s[start]
		s[start] = s[end]
		s[end] = temp
		start++
		end--
	}
}

func LeftRotateString(s []byte, StringLen int, offset int) []byte {
	ReverseString(s, 0, offset-1)
	ReverseString(s, offset, StringLen-1)
	ReverseString(s, 0, StringLen-1)
	return s
}
