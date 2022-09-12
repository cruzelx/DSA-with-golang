package main

import (
	"fmt"
	"math"
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

func (node *Node) DeleteNode(val int) *Node {
	if node == nil {
		return node
	}

	if val > node.Val {
		node.Right = node.Right.DeleteNode(val)
		return node
	}
	if val < node.Val {
		node.Left = node.Left.DeleteNode(val)
		return node
	}

	if node.Left == nil && node.Right == nil {
		return nil
	}

	if node.Left == nil {
		temp := node.Right
		node = nil
		return temp
	}

	if node.Right == nil {
		temp := node.Left
		node = nil
		return temp
	}

	parent := node
	child := parent.Right

	for child.Left != nil {
		parent = child
		child = child.Left
	}

	if parent != node {
		parent.Left = child.Right
	} else {
		parent.Right = child.Right
	}

	node.Val = child.Val
	return node

}

func NodesOfLevel(node *Node, level int, height int, nodesarr *[]*Node) *[]*Node {
	var nodes []*Node
	if node == nil {
		return nil
	}

	NodesOfLevel(node.Left, level+1, height, nodesarr)
	NodesOfLevel(node.Right, level+1, height, nodesarr)

	if level == height {
		fmt.Println("Level Nodes: ", node)
		*nodesarr = append(*nodesarr, node)
		fmt.Println("Level Nodes arr: ", nodes)
	}
	return nodesarr
}

func PrintTree(node *Node) {
	height := heightOfBinaryTree(node)
	levels := height + 1

	// convert each level of tree to array
	// arr := [][]int{}
	for i := 0; i < levels; i++ {
		for j := 0; j < int(math.Pow(2, float64(i))); j++ {
			fmt.Print(j)
		}
		fmt.Println("\n")
	}

	for i := 0; i < levels; i++ {
		printSpace(int(math.Pow(2, float64(levels))) - 1)
		// fmt.Print(node.Val)
	}
}

func printSpace(x int) {
	for i := 0; i < x; i++ {
		fmt.Print(" ")
	}
}

// func PrintTree(node *Node) {
// 	if node == nil {
// 		return
// 	}

// 	fmt.Print(node.Val, " ")

// 	PrintTree(node.Left)
// 	PrintTree(node.Right)

// }
func heightOfBinaryTree(node *Node) int {
	if node == nil {
		return -1
	}

	lCount := heightOfBinaryTree(node.Left)
	rCount := heightOfBinaryTree(node.Right)

	if lCount > rCount {
		lCount += 1
		return lCount
	} else {
		rCount += 1
		return rCount
	}

	return -1
}

func GenerateBST() *Node {
	rand.Seed(time.Now().UnixNano())

	numList := []int{rand.Intn((100))}

	node := &Node{Val: numList[0]}

	for i := 0; i < 10; i++ {
		randNum := rand.Intn(100)
		node.InsertNode(randNum)

		numList = append(numList, randNum)
	}

	fmt.Println(numList)
	return node

}
