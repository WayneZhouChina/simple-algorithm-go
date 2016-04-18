package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 7, 11, 12, 15}
	length := len(a)
	sum := 15
	towSumMy(a, length, sum)
}

func towSumMy(array []int, length int, sum int) {
	begin := 0
	end := length - 1

	for begin < end {
		curSum := array[begin] + array[end]
		if curSum == sum {
			fmt.Printf("Value %d and value %d which is we want\n", array[begin], array[end])
			// if you want to print all satisfy value
			begin += 1
			end -= 1
		} else {
			if curSum < sum {
				begin += 1
			} else {
				end -= 1
			}
		}
	}
}
