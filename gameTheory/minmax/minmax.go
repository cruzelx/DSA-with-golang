package main

func minmax(currentDepth int, nodeIndex int, maxTurn bool, scores []int, treeDepth int) int {
	if currentDepth == treeDepth {
		return scores[nodeIndex]
	}

	if maxTurn {
		return Max(minmax(currentDepth+1, nodeIndex*2, false, scores, treeDepth), minmax(currentDepth+1, nodeIndex*2+1, false, scores, treeDepth))
	} else {
		return Min(minmax(currentDepth+1, nodeIndex*2, true, scores, treeDepth), minmax(currentDepth+1, nodeIndex*2+1, true, scores, treeDepth))
	}
}

func Max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func Min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
