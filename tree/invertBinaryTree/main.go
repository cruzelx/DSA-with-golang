package main

import "fmt"

func main() {
	tree := GenerateBST()
	PrintTree(tree)

	fmt.Println("\n")

	inverted := InvertBinaryTree(tree)
	PrintTree(inverted)
}
