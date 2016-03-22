package main

import (
	"fmt"
)

func main() {
	s := []byte("abc")
	CalcAllPermutation(s, 0, 2)
}

func CalcAllPermutation(perm []byte, from int, to int) {
	if to <= 1 {
		return
	}

	if from == to {
		fmt.Println(string(perm))
	} else {
		for j := from; j <= to; j++ {
			temp := perm[j]
			perm[j] = perm[from]
			perm[from] = temp

			CalcAllPermutation(perm, from+1, to)

			temp1 := perm[j]
			perm[j] = perm[from]
			perm[from] = temp1
		}
	}
}
