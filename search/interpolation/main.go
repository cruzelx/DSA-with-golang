package main

import (
	"fmt"
	"time"
)

func main() {
	arr := []int64{1, 2, 3, 4, 6, 8, 9}

	start := time.Now()
	found, index, value := InterpolationSearch(arr, 1)
	elapsed := time.Since(start)

	fmt.Printf("Found: %v	Item:%v		Position:%v		Duration:%v", found, value, index, elapsed)
}
