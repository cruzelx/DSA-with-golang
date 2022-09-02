package main

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

func TernarySearch(arr []int, num int) (bool, int, int) {
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

func TestData() []int {
	rand.Seed(time.Now().Unix())

	numArr := rand.Perm(10000000)
	sort.Slice(numArr, func(i, j int) bool {
		return numArr[i] < numArr[j]
	})
	return numArr
}
