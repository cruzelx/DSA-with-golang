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
	Head *Node
}

func newNode(data int) *Node {
	return &Node{
		Data: data,
		Next: nil,
	}
}

func newList() *LinkedList {
	return &LinkedList{
		Head: nil,
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
}

func (list *LinkedList) Prepend(data int) {
	newNode := &Node{Data: data, Next: list.Head}
	list.Head = newNode
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

func GenerateLinkedList() *LinkedList {
	rand.Seed(time.Now().UnixNano())
	randArr := rand.Perm(5)

	list := newList()

	for _, v := range randArr {

		list.Append(v)
	}

	return list
}
