package day8

import (
	"strings"
)

func Puzzle8_1(input []string) int {
	count := 0
	for _, e := range input {
		tmp := strings.Split(e, " ")
		for _, e := range tmp {
			if len(e) == 7 || len(e) == 4 || len(e) == 3 || len(e) == 2 {
				count += 1
			}
		}
	}
	return count
}
