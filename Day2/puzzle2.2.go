// https://adventofcode.com/2021/day/2#part2

package day2

import (
	utility "AoC-2021/Utility"
)

func Puzzle2_2(input []utility.ActionValue) int {
	depth, position, aim := 0, 0, 0
	for _, actVal := range input {
		if actVal.Action == "forward" {
			position += actVal.Val
			depth += aim * actVal.Val
		} else if actVal.Action == "down" {
			aim += actVal.Val
		} else {
			aim -= actVal.Val
		}
	}
	return depth * position
}