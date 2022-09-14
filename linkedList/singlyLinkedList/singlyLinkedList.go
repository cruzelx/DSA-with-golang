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

func (node *Node) AppendNode(val int) {
	last := node
	for last.Next != nil {
		last = last.Next
	}
	last.Next = &Node{Data: val}
}

func (node *Node) PushNode(val int) {
	temp := node
	newNode := &Node{Data: val, Next: temp}
	temp = newNode
}

// func (node *Node) InsertAfterNode(prevNode *Node, val int) {
// 	if prevNode == nil {
// 		return
// 	}
// 	prevNode.Next = &Node{Next: prevNode.Next, Data: val}
// }

func (node *Node) PrintLinkedList() {
	temp := node
	fmt.Print(temp.Data, " ")
	for temp.Next != nil {
		temp = temp.Next
		fmt.Print(temp.Data, " ")
	}
	fmt.Println("\n")
}

func GenerateLinkedList() *Node {
	rand.Seed(time.Now().UnixNano())
	randArr := rand.Perm(20)

	node := &Node{Data: rand.Intn(100)}

	for _, v := range randArr {
		node.AppendNode(v)
	}
	return node
}
