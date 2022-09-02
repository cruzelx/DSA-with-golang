package main

import (
	"math/rand"
	"sort"
	"time"
)

func InterpolationSearch(arr []int, num int) (bool, int, int) {
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

func TestData() []int {
	rand.Seed(time.Now().Unix())

	numArr := rand.Perm(10000000)
	sort.Slice(numArr, func(i, j int) bool {
		return numArr[i] < numArr[j]
	})
	return numArr
}

