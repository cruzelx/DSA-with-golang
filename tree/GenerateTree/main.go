package main

import "fmt"

func main() {
	tree := GenerateBST()
	PrintTree(tree)
	height := heightOfBinaryTree(tree)
	fmt.Println("Height: ", height)

	for i := 30; i < 90; i++ {
		_ = tree.DeleteNode(i)
	}

	PrintTree(tree)
	height = heightOfBinaryTree(tree)
	fmt.Println("Height: ", height)

	var nodesarr *[]*Node
	fmt.Println("Nodes of level: ", NodesOfLevel(tree, 2, height, nodesarr))
	fmt.Println("data: ", nodesarr)
}
