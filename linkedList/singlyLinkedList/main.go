package main

func main() {
	list := GenerateLinkedList()
	list.PrintLinkedList()

	list.AppendNode(101)
	list.PrintLinkedList()
	
	list.PushNode(102)
	// list.InsertAfterNode()
	list.PrintLinkedList()

}
