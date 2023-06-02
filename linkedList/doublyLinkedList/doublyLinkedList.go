package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	Data int
	Next *Node
	Prev *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

func NewNode(data int) *Node {
	return &Node{Data: data, Next: nil, Prev: nil}
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		Head: nil,
		Tail: nil,
	}
}

func (dll *DoublyLinkedList) PrintForward() {
	curr := dll.Head

	for curr != nil {
		fmt.Print(curr.Data, " ")
		curr = curr.Next
	}
	fmt.Println()
}

func (dll *DoublyLinkedList) PrintBackward() {
	curr := dll.Tail

	for curr != nil {
		fmt.Print(curr.Data, " ")
		curr = curr.Prev
	}

	fmt.Println()

}

func (dll *DoublyLinkedList) Prepend(data int) {
	node := &Node{Data: data}

	if dll.Head == nil {
		dll.Head = node
		dll.Tail = node
		node.Next = nil
		node.Prev = nil
		return
	}

	node.Next = dll.Head
	dll.Head.Prev = node
	dll.Head = node
}

func (dll *DoublyLinkedList) Append(data int) {
	node := &Node{Data: data}

	if dll.Head == nil {
		dll.Head = node
		dll.Tail = node
		node.Next = nil
		node.Prev = nil
		return
	}

	node.Prev = dll.Tail
	dll.Tail.Next = node
	dll.Tail = node

}

func (dll *DoublyLinkedList) InsertAfter(data int, index int) {
	if index < 0 {
		return
	}

	if index == 0 {
		dll.Prepend(data)
	}

	curr := dll.Head

	for i := 0; i < index; i++ {
		if curr == nil {
			return
		}
		curr = curr.Next
	}

	if curr == nil {
		dll.Append(data)
	}

	curr.Next = &Node{Data: data, Prev: curr, Next: curr.Next}
}

func GenerateDoublyLinkedList() *DoublyLinkedList {
	rand.Seed(time.Now().UnixNano())
	randArr := rand.Perm(5)

	list := NewDoublyLinkedList()

	for _, v := range randArr {

		list.Append(v)
	}

	return list
}
