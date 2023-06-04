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
	return &DoublyLinkedList{
		Head: nil,
		Tail: nil,
	}
}

func (dll *DoublyLinkedList) Prepend(node *Node) {

	if dll.Head == nil {
		dll.Head = node
		dll.Tail = node
		dll.Head.Prev = nil
		dll.Tail.Next = nil
		return
	}

	node.Next = dll.Head
	dll.Head.Prev = node
	dll.Head = node

}

func (dll *DoublyLinkedList) Remove(node *Node) {
	if node == dll.Head {
		dll.Head = dll.Head.Next
		dll.Head.Prev = nil
		return
	}

	if node == dll.Tail {
		dll.Tail = dll.Tail.Prev
		dll.Tail.Next = nil
		return
	}

	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (dll *DoublyLinkedList) RemoveTail() {
	if dll.Head == nil {
		return
	}

	if dll.Tail == dll.Head {
		dll.Head = nil
		dll.Tail = nil
		return
	}

	currTail := dll.Tail
	dll.Tail = dll.Tail.Prev
	dll.Tail.Next = nil
	currTail.Prev = nil

}
