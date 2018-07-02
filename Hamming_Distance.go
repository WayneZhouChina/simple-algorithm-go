package main

func hammingDistance(x int, y int) int {
	var res int = 0
	var i uint = 0
	for i = 0; i < 32; i++ {
		x = x & (1 << i)
		y = y & (1 << i)

		if (x ^ y) > 0 {
			res++
		}
	}

	return res
}
