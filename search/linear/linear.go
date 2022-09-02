package main

import (
	"math/rand"
	"sort"
	"time"
)

func LinearSearch(arr []int, num int) (bool, int) {
	for i, n := range arr {
		if n == num {
			return true, i
		}
	}
	return false, -1
}

func TestData() []int {
	rand.Seed(time.Now().Unix())

	numArr := rand.Perm(10000000)
	sort.Slice(numArr, func(i, j int) bool {
		return numArr[i] < numArr[j]
	})
	return numArr
}

