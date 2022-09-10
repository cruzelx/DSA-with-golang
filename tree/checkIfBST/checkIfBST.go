package main

// Check if the tree is Binary Search Tree using min-max algorithm
func CheckIfBST(root *Node, min int, max int) bool {
	if root == nil {
		return true
	}

	if root.Val > max || root.Val < min {
		return false
	}

	return CheckIfBST(root.Left, min, root.Val) && CheckIfBST(root.Right, root.Val, max)
}
