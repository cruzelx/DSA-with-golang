package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"time"
)

type TicTacToe struct {
	Board                   [][]string
	Player, Opponent, Empty string
}

func (ttt *TicTacToe) PrintBoard() {
	for _, v := range ttt.Board {
		fmt.Printf("\t%s\t|\t%s\t|\t%s\t\n", v[0], v[1], v[2])
	}
	fmt.Println("\n")
}

func (ttt *TicTacToe) IsMoveLeft() bool {
	for _, i := range ttt.Board {
		for _, j := range i {
			if j == ttt.Empty {
				return true
			}
		}
	}
	return false
}

func (ttt *TicTacToe) EvaluateBoard() int {
	// check rows
	for i := 0; i < 3; i++ {
		if ttt.Board[i][0] == ttt.Board[i][1] && ttt.Board[i][1] == ttt.Board[i][2] {
			if ttt.Board[i][0] == ttt.Opponent {
				return -10
			} else if ttt.Board[i][0] == ttt.Player {
				return 10
			}
		}
	}

	//check columns
	for j := 0; j < 3; j++ {
		if ttt.Board[0][j] == ttt.Board[1][j] && ttt.Board[1][j] == ttt.Board[2][j] {
			if ttt.Board[0][j] == ttt.Opponent {
				return -10
			} else if ttt.Board[0][j] == ttt.Player {
				return 10
			}
		}
	}

	// check diagonals
	if ttt.Board[0][0] == ttt.Board[1][1] && ttt.Board[1][1] == ttt.Board[2][2] {
		if ttt.Board[0][0] == ttt.Opponent {
			return -10
		} else if ttt.Board[0][0] == ttt.Player {
			return 10
		}
	}

	if ttt.Board[0][2] == ttt.Board[1][1] && ttt.Board[1][1] == ttt.Board[2][0] {
		if ttt.Board[0][2] == ttt.Opponent {
			return -10
		} else if ttt.Board[0][2] == ttt.Player {
			return 10
		}
	}

	return 0
}

func (ttt *TicTacToe) Minmax(depth int, isMax bool) int {
	score := ttt.EvaluateBoard()

	if score == 10 || score == -10 {
		return score
	}

	if !ttt.IsMoveLeft() {
		return 0
	}

	if isMax {
		bestMoveVal := math.MinInt

		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if ttt.Board[i][j] == ttt.Empty {
					ttt.Board[i][j] = ttt.Player
					bestMoveVal = Max(bestMoveVal, ttt.Minmax(depth+1, false))
					ttt.Board[i][j] = ttt.Empty
				}
			}
		}
		return bestMoveVal
	} else {
		bestMoveVal := math.MaxInt

		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if ttt.Board[i][j] == ttt.Empty {
					ttt.Board[i][j] = ttt.Opponent
					bestMoveVal = Min(bestMoveVal, ttt.Minmax(depth+1, true))
					ttt.Board[i][j] = ttt.Empty
				}
			}
		}
		return bestMoveVal
	}
}

func (ttt *TicTacToe) FindBestMove(gamer string) []int {
	var bestVal int
	bestIndex := []int{-1, -1}

	if gamer == ttt.Player {
		bestVal = math.MinInt
	} else if gamer == ttt.Opponent {
		bestVal = math.MaxInt
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

			if ttt.Board[i][j] == ttt.Empty {
				var moveVal int

				if gamer == ttt.Player {

					ttt.Board[i][j] = ttt.Player
					moveVal = ttt.Minmax(0, false)

					ttt.Board[i][j] = ttt.Empty

					if moveVal > bestVal {
						bestVal = moveVal
						bestIndex = []int{i, j}
					}
				} else if gamer == ttt.Opponent {

					ttt.Board[i][j] = ttt.Opponent
					moveVal = ttt.Minmax(0, true)

					ttt.Board[i][j] = ttt.Empty

					if moveVal < bestVal {
						bestVal = moveVal
						bestIndex = []int{i, j}
					}
				}
			}
		}
	}
	return bestIndex
}

func (ttt *TicTacToe) Play() {
	endGame := false
	playerTurn := false
	var gamer string

	for !endGame {

		ttt.PrintBoard()

		time.Sleep(2 * time.Second)

		// clear console
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		if !ttt.IsMoveLeft() {
			fmt.Println("Game Finished")
			return
		}

		if ttt.EvaluateBoard() == 10 {
			fmt.Println("Player won...")
			break
		}

		if ttt.EvaluateBoard() == -10 {
			fmt.Println("Opponent won...")
			break
		}

		if playerTurn {
			gamer = ttt.Player
		} else {
			gamer = ttt.Opponent
		}

		bestIndex := ttt.FindBestMove(gamer)

		ttt.Board[bestIndex[0]][bestIndex[1]] = gamer
		playerTurn = !playerTurn

	}

	return
}
