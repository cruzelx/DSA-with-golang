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
	Head   *Node
	Tail   *Node
	Length int
}

func NewNode(data int) *Node {
	return &Node{Data: data, Next: nil, Prev: nil}
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		Head:   nil,
		Tail:   nil,
		Length: 0,
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
		dll.Length++
		return
	}

	node.Next = dll.Head
	dll.Head.Prev = node
	dll.Head = node
	dll.Length++
}

func (dll *DoublyLinkedList) Append(data int) {
	node := &Node{Data: data}

	if dll.Head == nil {
		dll.Head = node
		dll.Tail = node
		node.Next = nil
		node.Prev = nil
		dll.Length++
		return
	}

	node.Prev = dll.Tail
	dll.Tail.Next = node
	dll.Tail = node
	dll.Length++

}

func (dll *DoublyLinkedList) InsertAfter(data int, index int) {

	if index < 0 {
		return
	}

	if index == 0 {
		dll.Prepend(data)
		dll.Length++
		return
	}

	curr := dll.Head

	for i := 0; i < index; i++ {
		if index >= dll.Length {
			return
		}
		if curr == nil {
			return
		}
		curr = curr.Next

	}

	if curr == nil {
		dll.Append(data)
		dll.Length++
		return
	}

	curr.Next = &Node{Data: data, Prev: curr, Next: curr.Next}
	dll.Length++
}

func (dll *DoublyLinkedList) Remove(index int) {
	curr := dll.Head

	if index < 0 {
		return
	}

	if index == 0 {
		dll.Head = dll.Head.Next
		dll.Head.Prev = nil
		dll.Length--
		return
	}

	prev := curr
	for i := 0; i < index; i++ {
		if index >= dll.Length {
			return
		}
		prev = curr
		curr = curr.Next
	}

	prev.Next = curr.Next
	if curr.Next != nil {
		prev.Next.Prev = prev
	}
	dll.Length--

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
