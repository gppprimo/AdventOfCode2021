// https://adventofcode.com/2021/day/2

package day2

import (
	utility "AoC-2021/Utility"
)

func Puzzle2_1(input []utility.ActionValue) int {
	depth, position := 0, 0
	for _, actVal := range input {
		if actVal.Action == "forward" {
			position += actVal.Val
		} else if actVal.Action == "down" {
			depth += actVal.Val
		} else {
			depth -= actVal.Val
		}
	}
	return depth * position
}
