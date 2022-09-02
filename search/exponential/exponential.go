package main

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

func binarySearch(arr []int, start int, end int, num int) (bool, int) {
	for start <= end {
		mid := start + int(math.Floor(float64(end-start)/2))

		if arr[mid] == num {
			return true, mid
		}

		if num <= arr[mid] {
			end = mid - 1
		}

		if num > arr[mid] {
			start = mid + 1
		}

	}
	return false, -1
}

func ExponentialSearch(arr []int, num int) (bool, int, int) {
	arrLength := len(arr)
	if num == arr[0] {
		return true, 0, arr[0]
	}

	i := 1
	for i < arrLength && arr[i] <= num {
		i = i * 2
	}
	found, index := binarySearch(arr, i/2, int(math.Min(float64(i), float64(arrLength-1))), num)
	if !found {
		return false, -1, 0
	}
	return true, index, arr[index]
}

func TestData() []int {
	rand.Seed(time.Now().Unix())

	numArr := rand.Perm(10000000)
	sort.Slice(numArr, func(i, j int) bool {
		return numArr[i] < numArr[j]
	})
	return numArr
}
