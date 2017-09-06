package main

func repeatedSubstringPattern(s string) bool {
	length := len(s) / 2
	for i := 0; i < length; i++ {
		if s[i] != s[length-1] {
			return false
		}
	}

	return true
}
