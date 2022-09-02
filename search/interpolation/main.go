package main

import (
	"fmt"
	"time"
)

func main() {
	arr := TestData()

	start := time.Now()
	found, index, value := InterpolationSearch(arr, 9999912)
	elapsed := time.Since(start)

	fmt.Printf("Found: %v	Item:%v		Position:%v		Duration:%v", found, value, index, elapsed)
}
