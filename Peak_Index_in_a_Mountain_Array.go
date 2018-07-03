package main

func peakIndexInMountainArray(A []int) int {
	var left, right, mid int = 0, len(A) - 1, 0

	for left < right {
		mid = (left + right) / 2

		if A[mid] > A[mid-1] && A[mid+1] > A[mid] {
			left = mid
		} else if (A[mid] > A[mid+1]) && (A[mid-1] > A[mid]) {
			right = mid
		} else {
			break
		}
	}

	return mid
}
