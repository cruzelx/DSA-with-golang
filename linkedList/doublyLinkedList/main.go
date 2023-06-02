package main

func main() {
	dll := GenerateDoublyLinkedList()
	dll.PrintForward()
	dll.PrintBackward()

	dll.Append(1024)
	dll.PrintForward()

	dll.Prepend(2022)
	dll.PrintForward()

	// dll.InsertAfter(2, 9999)
	// dll.PrintForward()
}
