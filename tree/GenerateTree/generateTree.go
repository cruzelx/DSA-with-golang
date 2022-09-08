package main

import (
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

	if node.Val == val {
		return true
	}

	if node.Val > val {
		node.Left.SearchNode(val)
	} else if node.Val < val {
		node.Right.SearchNode(val)
	}

}

func GenerateBST() *Node {
	rand.Seed(time.Now().UnixNano())

	node := &Node{Val: rand.Intn(100)}

	for i := 0; i < 20; i++ {
		node.InsertNode(rand.Intn(100))
	}
	return node

}
