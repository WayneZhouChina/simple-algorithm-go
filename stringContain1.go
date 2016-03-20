package main

//import "fmt"

func main() {
	a := "ABCD"
	b := "CDE"
	println(StringContain(a, b))
}

// b is a part of a
func StringContain(a string, b string) bool {
	for i := 0; i < len(b); i++ {
		var j int
		for j = 0; j < len(a); j++ {
			if string(b[i]) != string(a[j]) {
				continue
			} else {
				break
			}
		}
		if j >= len(a) {
			return false
		}
	}
	return true
}
