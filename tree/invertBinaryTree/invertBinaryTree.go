package main

func InvertBinaryTree(node *Node) *Node {
	if node == nil {
		return nil
	}

	InvertBinaryTree(node.Left)
	InvertBinaryTree(node.Right)

	node.Left, node.Right = node.Right, node.Left
	return node
}
