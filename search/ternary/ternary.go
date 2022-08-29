package main

import "math"

func TernarySearch(arr []int64, num int64) (bool, int, int64) {
	start := 0
	end := len(arr) - 1

	for start <= end {

		mid1 := start + int(math.Floor(float64(end-start)/3))
		mid2 := end - int(math.Floor(float64(end-start)/3))

		if num == arr[mid1] {
			return true, mid1, arr[mid1]
		}

		if num == arr[mid2] {
			return true, mid2, arr[mid2]
		}

		if num < arr[mid1] {
			end = mid1 - 1
		} else if num > arr[mid2] {
			start = mid2 + 1
		} else {
			start = mid1 + 1
			end = mid2 - 1
		}

	}
	return false, -1, 0

}
