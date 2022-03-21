package day13

import (
	utility "AoC-2021/Utility"
)



func Puzzle13_1(tPaper []utility.Struct13_1, folds []utility.Struct13_1XY) int {
	xMax := 0
	yMax := 0
	for i := range tPaper {
		if tPaper[i].X > xMax {
			xMax = tPaper[i].X
		}
		if tPaper[i].Y > yMax {
			yMax = tPaper[i].Y
		}
	}
	return 0
}
