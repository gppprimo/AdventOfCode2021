package day4

func getTable(boards [][]int, i int, j int) [][]int {
	return boards[i:j]
}

func checkTableBitmap(boards [][]int, bitmap [][]int, rnd int, startIdx int, stopIdx int) {
	for i := startIdx; i < stopIdx; i++ {
		for j, elem := range boards[i] {
			if elem == rnd {
				bitmap[i][j] = 1
			}
		}
	}
}

func checkOnes(arr []int) bool {
	for _, e := range arr {
		if e != 1 {
			return false
		}
	}
	return true
}

func checkBingo(bitmap [][]int, startIdx int, stopIdx int) bool {
	for i := startIdx; i < stopIdx; i++ {
		row := bitmap[i]
		var column []int
		for j := 0; j < len(row); j++ {
			column = append(column, bitmap[j][i-startIdx])
		}
		if checkOnes(row) || checkOnes(column) {
			return true
		}
	}
	return false
}

func sumUnmarked(table [][]int, bitmap [][]int, startIdx int, stopIdx int) int {
	sum := 0
	for i := startIdx; i < stopIdx; i++ {
		for j := 0; j < len(bitmap[i]); j++ {
			if bitmap[i][j] == 0 {
				sum += table[i][j]
			}
		}
	}
	return sum
}

func Puzzle4_1(randoms []int, boards [][]int) int {
	bitmap := make([][]int, len(boards))
	for i := 0; i < len(boards); i++ {
		bitmap[i] = make([]int, len(boards[0]))
	}

	for _, rnd := range randoms {
		for i := 0; i <= len(boards)-len(boards[0]); i += len(boards[0]) {
			j := i + len(boards[0])
			checkTableBitmap(boards, bitmap, rnd, i, j)
			if checkBingo(bitmap, i, j) {
				return sumUnmarked(boards, bitmap, i, j) * rnd
			}

		}
	}
	return 0
}
