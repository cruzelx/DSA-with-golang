package main

import (
	"fmt"
	"math"
)

func main() {
	tree := GenerateBST()
	PrintTree(tree)

	isBST := CheckIfBST(tree, math.MinInt, math.MaxInt)
	fmt.Println("IS BST ? : ", isBST)

	tree = GenerateNonBST()
	PrintTree(tree)

	isBST = CheckIfBST(tree, math.MinInt, math.MaxInt)
	fmt.Println("IS BST ? : ", isBST)

}
