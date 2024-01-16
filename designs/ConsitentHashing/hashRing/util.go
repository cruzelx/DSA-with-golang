package main

import (
	"fmt"
	"sync"
)


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
