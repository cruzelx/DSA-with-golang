package main

import "fmt"

func main() {
	tree := GenerateBST()
	PrintTree(tree)

	fmt.Println("\n")

	height := heightOfBinaryTree(tree)
	fmt.Print("Height: ", height)

}
