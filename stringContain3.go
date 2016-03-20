package main

func main() {
	a := "ABCD"
	b := "CDEF"
	println(StringContain3(a, b))
}

func StringContain3(a string, b string) bool {

	p := [26]int{
		2, 3, 5, 7, 9,
		11, 13, 17, 19, 23,
		29, 31, 37, 41, 43,
		47, 53, 59, 61, 67,
		71, 73, 79, 83, 89, 97}

	var f int = 1
	for i := 0; i < len(a); i++ {
		value := p[a[i]-'A']
		f *= value
	}

	for j := 0; j < len(b); j++ {
		v := p[b[j]-'A']
		if f%v > 0 {
			return false
		}
	}

	return true
}
