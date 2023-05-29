package main

import "fmt"

func main() {
	mapper := NewHashTable(3)

	mapper.set("fruits", []string{"orange", "apple"})
	mapper.set("age", 12)
	mapper.set("is_adult", false)
	mapper.set("activity", "swimming")

	mapper.display()
	fmt.Println()

	mapper.remove("activity")

	mapper.display()
	fmt.Println()
	fmt.Println(mapper.get("fruits"))
	fmt.Println()

	mapper.set("activity", "football")
	mapper.display()

}
