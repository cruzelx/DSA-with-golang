package main

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
