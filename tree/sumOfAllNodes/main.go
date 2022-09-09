package main

import "fmt"

func main() {
	tree := GenerateBST()
	PrintTree(tree)

	fmt.Println("\n")

	sum := SumOfAllNodes(tree)
	fmt.Print("Sum of all nodes: ", sum)
}
