package main

import (
	"fmt"
	"time"
)

func main() {
	arr := []int64{1, 2, 3, 4, 6, 9, 8}

	start := time.Now()
	found, index := LinearSearch(arr, 1)
	elapsed := time.Since(start)

	fmt.Printf("Found: %v	Item:%v		Position:%v		Duration:%v", found, arr[index], index, elapsed)
}
