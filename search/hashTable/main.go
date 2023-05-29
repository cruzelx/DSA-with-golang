package main

import "fmt"

func main() {
	// Init hash map with buket size, 
	mapper := NewHashTable(2)

	mapper.set("fruits", []string{"orange", "apple"})
	mapper.set("age", 12)
	mapper.set("is_adult", false)
	mapper.set("activity", "swimming")
	mapper.set("flavour", "spicy")
	mapper.set("flavour", "spicy")

	mapper.display()
	fmt.Println()

	mapper.remove("activity")

	mapper.display()
	fmt.Println()

	mapper.set("activity", "football")
	mapper.set("activiyt", "volley ball")
	mapper.display()

}
