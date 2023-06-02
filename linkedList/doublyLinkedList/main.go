package main

func main() {
	dll := GenerateDoublyLinkedList()
	dll.PrintForward()
	dll.PrintBackward()

	dll.Append(1024)
	dll.PrintForward()

	dll.Prepend(2022)
	dll.PrintForward()

	dll.InsertAfter(9999, 2)
	dll.PrintForward()

	dll.Remove(100)
	dll.PrintForward()
}
