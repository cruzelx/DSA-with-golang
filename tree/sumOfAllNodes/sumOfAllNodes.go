package main

func SumOfAllNodes(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Val + SumOfAllNodes(node.Left) + SumOfAllNodes(node.Right)

}
