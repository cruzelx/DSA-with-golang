package main

import (
	"fmt"
	"sync"
)

// func PrintMapAsTable(m map[string]string) {
// 	// Calculate column widths based on keys and values
// 	keyWidth := 0
// 	valueWidth := 0
// 	for key, value := range m {
// 		keyWidth = Max(keyWidth, len(key))
// 		valueWidth = Max(valueWidth, len(value))
// 	}

// 	// Print header
// 	fmt.Println(strings.Repeat("-", keyWidth+valueWidth+3))
// 	fmt.Printf("| %-*s | %-*s |\n", keyWidth, "Key", valueWidth, "Value")
// 	fmt.Println(strings.Repeat("-", keyWidth+valueWidth+3))

// 	// Print each key-value pair
// 	for key, value := range m {
// 		fmt.Printf("| %-*s | %-*s |\n", keyWidth, key, valueWidth, value)
// 	}

// 	fmt.Println(strings.Repeat("-", keyWidth+valueWidth+3))
// }

func PrintMapAsTable(m *sync.Map) {
	fmt.Println("Key | Value")
	fmt.Println("-------------")

	m.Range(func(key, value interface{}) bool {
		fmt.Printf("%v \t\t| %d\n", key, len(value.([]KeyVal)))
		return true
	})
}
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
