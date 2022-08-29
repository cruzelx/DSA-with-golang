package main

func LinearSearch(arr []int64, num int64) (bool, int) {
	for i, n := range arr {
		if n == num {
			return true, i
		}
	}
	return false, -1
}
