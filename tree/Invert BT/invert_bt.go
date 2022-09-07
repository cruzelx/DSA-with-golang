package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

func NewNode(data int) *Node {
	node := &Node{
		val:   data,
		left:  nil,
		right: nil,
	}
	return node
}

func InvertBinaryTree(node *Node) {
	if node == nil {
		return
	}
	temp := node

	node.left = temp.right
	node.right = temp.left

	InvertBinaryTree(node.left)
	InvertBinaryTree(node.right)
}

func Print_BT(node *Node) {
	if node.left != nil {
		Print_BT(node.left)
	}

	fmt.Printf("%d ", node.val)

	if node.right != nil {
		Print_BT(node.right)
	}
}

var count = 0

func GenerateBinaryTree(arr []int) *Node {
	rand.Seed(time.Now().UnixNano())
	count = count + 1

	if len(arr) == 0 {
		return nil
	}

	if len(arr) == 1 {
		return NewNode(arr[0])
	}
	// if len(arr) == 1 {
	// 	node = NewNode(arr[0])
	// 	return node
	// }

	max := len(arr)
	min := 1

	// if max-min < 1 {
	// 	return nil
	// }

	split := rand.Intn(max-min+1) + min
	split = split - 1

	left := arr[:split]
	right := arr[split:]

	node := NewNode(arr[split])

	node.left = GenerateBinaryTree(left)
	node.right = GenerateBinaryTree(right)

	fmt.Println("run count", count, arr, split, left, right, node.left, node.right)
	return node

}
