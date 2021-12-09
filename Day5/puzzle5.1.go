package day5

import (
	utility "AoC-2021/Utility"
	"fmt"
)

func countOverlaps(bitmap [][]int) int {
	count := 0
	for i := 0; i < len(bitmap); i++ {
		for j := 0; j < len(bitmap[0]); j++ {
			if bitmap[i][j] >= 2 {
				count += 1
			}
		}
	}
	return count
}

func checkBitmap(bitmap [][]int, x1 int, y1 int, x2 int, y2 int) {
	if x1 >= x2 {
		tmp := x1
		x1 = x2
		x2 = tmp
	}

	if y1 >= y2 {
		tmp := y1
		y1 = y2
		y2 = tmp
	}

	for j := x1; j <= x2; j++ {
		for i := y1; i <= y2; i++ {
			bitmap[i][j] += 1
		}
	}
}


func checkBitmapDiag(bitmap [][]int, x1 int, y1 int, x2 int, y2 int) {

}

func Puzzle5_1(input []utility.Struct5_1) int {
	maxX, maxY := 0, 0
	for _, elem := range input {
		if elem.X1 >= maxX {
			maxX = elem.X1
		} else if elem.X2 > maxX {
			maxX = elem.X2
		}
		if elem.Y1 >= maxY {
			maxY = elem.Y1
		} else if elem.Y2 > maxY {
			maxY = elem.Y2
		}
	}
	bitmap := make([][]int, maxX +1)
	for i := range bitmap {
		bitmap[i] = make([]int, maxY+2)
	}


	for _, inputLine := range input {
		if inputLine.Y1 == inputLine.Y2 || inputLine.X1 == inputLine.X2 {
			
			checkBitmap(bitmap, inputLine.X1, inputLine.Y1, inputLine.X2, inputLine.Y2)
		} else {
			fmt.Println(inputLine)
			checkBitmapDiag(bitmap, inputLine.X1, inputLine.Y1, inputLine.X2, inputLine.Y2)
		}
	}
	return countOverlaps(bitmap)
}
