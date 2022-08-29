package main

func InterpolationSearch(arr []int64, num int64) (bool, int, int64) {
	start := 0
	end := len(arr) - 1

	for start <= end && num <= arr[end] && num >= arr[start] {
		mid := int(float64(start) + (float64(num-arr[start]) * (float64(end-start) / float64(arr[end]-arr[start]))))

		if arr[mid] > num {
			end = mid - 1
		} else if arr[mid] < num {
			start = mid + 1
		} else {
			return true, mid, arr[mid]
		}

	}

	return false, -1, 0
}
