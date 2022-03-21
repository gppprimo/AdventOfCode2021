package day4

func Puzzle4_2(randoms []int, boards [][][]int) int {
	examined := make([]bool, len(boards))
	res := 0
	for _, rnd := range randoms {
		for i, board := range boards {
			checkNumnberOnBoar(board, rnd)
			if checkBingo(board) && !examined[i] {
				res = rnd * sumUnmarked(board)
				examined[i] = true
			}
		}
	}
	return res
}
