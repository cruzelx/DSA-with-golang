package main

import (
	"math/rand"
	"time"
)

func main() {
	// scores := []int{1, 2, 7, 4, 9, 3, 8, 11, 52, 0, 99, 12, 13, 14, 15, 5}
	// treeDepth := int(math.Log2(float64(len(scores))))

	// score := minmax(0, 0, true, scores, treeDepth)
	// fmt.Println("Score: ", score)

	rand.Seed(time.Now().UnixNano())

	board := [][]string{{"_", "_", "_"}, {"_", "_", "_"}, {"_", "_", "_"}}
	i, j := rand.Intn(3), rand.Intn(3)
	board[i][j] = "X"

	ttt := TicTacToe{Board: board, Opponent: "O", Player: "X", Empty: "_"}

	ttt.Play()

}
