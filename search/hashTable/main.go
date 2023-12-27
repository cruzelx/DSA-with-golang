package main

import (
	"fmt"
	"hashTable/hasher"
)

func main() {
	// Initial bucket size/array length
	bucketSize := 4

	// %load at which the bucket is resized/doubled and rehashed
	loadFactor := 80

	mapper := NewHashTable(bucketSize, loadFactor, hasher.Djb2)

	// Set few values
	mapper.Set("fruits", []string{"orange", "apple"})
	mapper.Set("age", 12)
	mapper.Set("is_adult", false)

	mapper.Display()
	fmt.Println()

	mapper.Set("activity", KeyVal{"sport", "swimming"})
	mapper.Set("flavour", "spicy")
	mapper.Set("flavour", "spicy")

	mapper.Display()
	fmt.Println()

	mapper.Remove("activity")

	mapper.Set("activity", "football")
	mapper.Set("activiyt", "volley ball")

	mapper.Display()

	// Plot standard deviation vs bucket sizes
	// Plot()
	// WebPlot()

}
