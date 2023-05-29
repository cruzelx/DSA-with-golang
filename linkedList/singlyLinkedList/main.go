package main

import "fmt"

func main() {
	list := GenerateLinkedList()
	// list.Print()

	list.Append(101)
	// list.Print()

	list.Prepend(102)
	list.Print()

	list.Remove(6)
	list.Print()

	list.Reverse()
	// list.Print()

	node := list.FindAtIndex(2)
	fmt.Println(node)

	list1 := GenerateLinkedList()
	// list1.Print()
	list2 := GenerateLinkedList()
	// list2.Print()

	zipped := ZipperList(list1, list2)
	zipped.Print()

}
