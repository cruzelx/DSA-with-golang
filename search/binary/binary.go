package main

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

func BinarySearch(arr []int, num int) (bool, int, int) {
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

func TestData() []int {
	rand.Seed(time.Now().Unix())

	numArr := rand.Perm(10000000)
	sort.Slice(numArr, func(i, j int) bool {
		return numArr[i] < numArr[j]
	})
	return numArr
}
