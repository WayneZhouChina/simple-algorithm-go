package main

import (
	"strings"
)

//

func numJewelsInStones(J string, S string) int {
	var count int = 0

	ss := strings.Split(S, "")
	for _, s := range ss {
		if strings.Contains(J, s) == true {
			count++
		}
	}

	return count

}
