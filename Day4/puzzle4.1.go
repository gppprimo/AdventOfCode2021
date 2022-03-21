// https://adventofcode.com/2021/day/4

package day4

func getTable(boards [][]int, i int, j int) [][]int {
	return boards[i:j]
}

func checkOnes(arr []int) bool {
	for _, e := range arr {
		if e != -1 {
			return false
		}
	}
	return true
}

func checkBingo(board [][]int) bool {
	for i := 0; i < len(board); i++ {
		row := board[i]
		var column []int
		for j := 0; j < len(board); j++ {
			column = append(column, board[j][i])
		}
		if checkOnes(row) || checkOnes(column) {
			return true
		}
	}
	return false
}

func sumUnmarked(board [][]int) int {
	sum := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != -1 {
				sum += board[i][j]
			}
		}
	}
	return sum
}

func checkNumnberOnBoar(board [][]int, rnd int) {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == rnd {
				board[i][j] = -1
			}
		}
	}
}

func Puzzle4_1(randoms []int, boards [][][]int) int {
	for _, rnd := range randoms {
		for _, board := range boards {
			checkNumnberOnBoar(board, rnd)
			if checkBingo(board) {
				return rnd * sumUnmarked(board)
			}
		}
	}
	return 0
}
