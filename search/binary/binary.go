package main

import "math"

func BinarySearch(arr []int64, num int64) (bool, int, int64) {
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := int(math.Floor(float64(start+end) / 2))
		if arr[mid] == num {
			return true, mid, arr[mid]
		}
		if num <= arr[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return false, -1, 0
}
