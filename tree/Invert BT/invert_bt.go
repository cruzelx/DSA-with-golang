package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (node *Node) InsertNode(val int) {
	if node.Val > val {
		if node.Left == nil {
			node.Left = &Node{Val: val}
		} else {
			node.Left.InsertNode(val)
		}
	} else if node.Val <= val {
		if node.Right == nil {
			node.Right = &Node{Val: val}
		} else {
			node.Right.InsertNode(val)
		}
	}
}

func (node *Node) SearchNode(val int) bool {
	if node == nil {
		return false
	}

	if node.Val > val {
		node.Left.SearchNode(val)
	} else if node.Val < val {
		node.Right.SearchNode(val)
	}
	return true
}

func InvertBinaryTree(node *Node) *Node {
	if node == nil {
		return nil
	}

	InvertBinaryTree(node.Left)
	InvertBinaryTree(node.Right)

	node.Left, node.Right = node.Right, node.Left
	return node
}

func PrintTree(node *Node) {
	if node == nil {
		return
	}

	// pre-order traversal
	fmt.Print(node.Val, " ")

	PrintTree(node.Left)
	PrintTree(node.Right)

}

func GenerateBST() *Node {
	rand.Seed(time.Now().UnixNano())

	numList := []int{rand.Intn((100))}

	node := &Node{Val: numList[0]}

	for i := 0; i < 5; i++ {
		randNum := rand.Intn(100)
		node.InsertNode(randNum)

		numList = append(numList, randNum)
	}

	fmt.Println(numList)
	return node

}
