package main

import (
	"fmt"
	"math/rand"
	"strings"
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

func printBinaryTree(node *Node, level int, isRight bool, indent string) {
	if node == nil {
		return
	}

	if level > 0 {
		var stb strings.Builder

		for i := 0; i < level-1; i++ {
			stb.WriteString("│   ")
		}
		if isRight {
			stb.WriteString("└── ")
			indent += "    "
		} else {
			stb.WriteString("├── ")
			indent += "│   "
		}
		fmt.Print(stb.String())
	}
	fmt.Println(node.Val)

	printBinaryTree(node.Right, level+1, true, indent)
	printBinaryTree(node.Left, level+1, false, indent)
}

func PrintTree(node *Node) {
	printBinaryTree(node, 0, false, "")
	// if node == nil {
	// 	return
	// }

	// fmt.Print(node.Val, " ")

	// PrintTree(node.Left)
	// PrintTree(node.Right)

}

func GenerateBST() *Node {
	rand.Seed(time.Now().UnixNano())

	numList := []int{rand.Intn((100))}

	node := &Node{Val: numList[0]}

	for i := 0; i < 100; i++ {
		randNum := rand.Intn(100)
		node.InsertNode(randNum)

		numList = append(numList, randNum)
	}

	fmt.Println(numList)
	return node

}

func GenerateNonBST() *Node {
	node := &Node{Val: 50}

	node.Left = &Node{Val: 20}
	node.Right = &Node{Val: 31}
	node.Left.Left = &Node{Val: 100}
	node.Left.Right = &Node{Val: 150}
	node.Right.Left = &Node{Val: 1000}
	node.Right.Right = &Node{Val: 1010}

	return node
}
