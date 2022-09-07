package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := rand.Perm(20)
	fmt.Println(arr)
	Print_BT(GenerateBinaryTree(arr))
}
