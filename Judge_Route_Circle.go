package main

import "strings"

func judgeCircle(moves string) bool {
	movesSlice := strings.Split(moves, "")
	var x int = 0
	var y int = 0

	for _, c := range movesSlice {
		switch c {
		case "U":
			x++
			break
		case "D":
			x--
			break
		case "R":
			y++
			break
		case "L":
			y--
			break
		default:
		}
	}

	if x == 0 && y == 0 {
		return true
	} else {
		return false
	}
}
