package main

func main() {
	tree := GenerateBST()
	PrintTree(tree)

	for i := 30; i < 80; i++ {
		_ = tree.DeleteNode(i)
	}

	PrintTree(tree)
}
