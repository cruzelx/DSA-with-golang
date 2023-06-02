package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func newNode(data int) *Node {
	return &Node{
		Data: data,
		Next: nil,
	}
}

func newList() *LinkedList {
	return &LinkedList{
		Head:   nil,
		Length: 0,
	}
}

func (list *LinkedList) Append(data int) {

	newNode := newNode(data)
	if list.Head == nil {
		list.Head = newNode
	} else {
		curr := list.Head
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = newNode
	}
	list.Length++
}

func (list *LinkedList) Prepend(data int) {
	newNode := &Node{Data: data, Next: list.Head}
	list.Head = newNode
	list.Length++
}

func (list *LinkedList) InsertAfter(data int, index int) {
	curr := list.Head

	if index < 0 {
		return
	}
	if index == 0 {
		list.Head.Next = &Node{Data: data, Next: curr.Next}
		list.Length++
		return
	}

	for i := 0; i < index; i++ {
		if index >= list.Length {
			return
		}
		curr = curr.Next
	}

	if curr == nil {
		list.Append(data)
		list.Length++
		return
	}

	curr.Next = &Node{Data: data, Next: curr.Next}
	list.Length++

	// for curr != nil {
	// 	if position == index {
	// 		next := curr.Next
	// 		curr.Next = &Node{Data: data, Next: next}
	// 	}
	// 	curr = curr.Next
	// 	position++
	// }
}

func (list *LinkedList) Remove(index int) {
	curr := list.Head

	if index < 0 {
		return
	}
	if index == 0 {
		list.Head = list.Head.Next
		list.Length--
		return
	}

	prev := curr
	for i := 0; i < index; i++ {
		if index >= list.Length {
			return
		}

		prev = curr
		curr = curr.Next
	}
	prev.Next = curr.Next
	list.Length--

}

func (list *LinkedList) Print() {
	curr := list.Head
	for curr != nil {
		fmt.Print(curr.Data, " ")
		curr = curr.Next
	}
	fmt.Println()
}

func (list *LinkedList) Reverse() {
	curr := list.Head
	var prev, next *Node

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	list.Head = prev
}

func (list *LinkedList) FindAtIndex(index int) *Node {
	curr := list.Head
	count := 0
	for curr != nil {
		if count == index {
			return curr
		}
		curr = curr.Next
		count++

	}
	return nil
}

func ZipperList(list1 *LinkedList, list2 *LinkedList) *LinkedList {

	curr1 := list1.Head
	curr2 := list2.Head

	for curr1 != nil && curr2 != nil {

		temp := curr1.Next
		curr1.Next = curr2

		curr2 = curr2.Next
		curr1.Next.Next = temp
		curr1 = temp
	}

	if curr2 != nil {
		curr1.Next = curr2
	}

	return list1

}

func (ll *LinkedList) RemoveAlternateNodes(head *Node) {
	if head == nil {
		return
	}

	curr := head.Next
	if curr == nil {
		return
	}

	head.Next = curr.Next
	ll.RemoveAlternateNodes(head.Next)
}

func GenerateLinkedList() *LinkedList {
	rand.Seed(time.Now().UnixNano())
	randArr := rand.Perm(5)

	list := newList()

	for _, v := range randArr {

		list.Append(v)
	}

	return list
}
