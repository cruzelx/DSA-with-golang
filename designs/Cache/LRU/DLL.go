package main

type Node struct {
	Key   int
	Value int
	Prev  *Node
	Next  *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

func NewNode(key int, value int) *Node {
	return &Node{
		Key:   key,
		Value: value,
		Prev:  nil,
		Next:  nil,
	}
}

func NewDLL() *DoublyLinkedList {
	sentinal := &Node{}
	return &DoublyLinkedList{
		Head: sentinal,
		Tail: sentinal,
	}
}

func (dll *DoublyLinkedList) Prepend(node *Node) {

	node.Next = dll.Head
	node.Prev = nil
	dll.Head.Prev = node
	dll.Head = node

}

func (dll *DoublyLinkedList) Remove(node *Node) {

	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (dll *DoublyLinkedList) RemoveTail() {

	currTail := dll.Tail
	dll.Tail = dll.Tail.Prev
	dll.Tail.Next = nil
	currTail.Prev = nil

}
