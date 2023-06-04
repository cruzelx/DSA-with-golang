package main

import "fmt"

func main() {
	lru := NewLRU(3)
	lru.Print()

	lru.Put(0, 12)
	lru.Print()

	lru.Put(3, 555)
	lru.Print()

	lru.Put(4, 999)
	lru.Print()

	lru.Put(11, 1111)
	lru.Print()

	fmt.Println(lru.Get(0))
	lru.Print()

	fmt.Println(lru.Get(3))
	lru.Print()

	fmt.Println(lru.Get(4))
	lru.Print()

	fmt.Println(lru.Get(11))
	lru.Print()
}
