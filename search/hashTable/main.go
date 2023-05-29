package main

func main() {
	mapper := NewHashTable()

	mapper.set("Alex", 24)
	mapper.set("Taste", "Sweet")

	mapper.display()

	mapper.get("Alex")

	mapper.remove("Alex")
	mapper.display()
}
