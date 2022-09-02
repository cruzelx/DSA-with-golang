package main

import (
	"fmt"
	"time"
)

func main() {
	arr := TestData()

	start := time.Now()
	found, index := LinearSearch(arr, 9999912)
	elapsed := time.Since(start)

	fmt.Printf("Found: %v	Item:%v		Position:%v		Duration:%v", found, arr[index], index, elapsed)
}
