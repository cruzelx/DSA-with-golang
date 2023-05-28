package main

func main() {
	list := GenerateLinkedList()
	list.Print()

	list.Append(101)
	list.Print()

	list.Prepend(102)
	list.Print()

	list.Reverse()
	list.Print()

}
