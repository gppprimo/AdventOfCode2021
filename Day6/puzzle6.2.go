package day6

import utility "AoC-2021/Utility"

func checkZeros(input []int) bool {
	res := false
	for j := range input {
		if input[j] == 0 {
			res = true
			input = append(input, 8)
			input[j] = 6
		}
	}
	return res
}

func Puzzle6_2(input []int) int {
	for loop := 0; loop < 256; loop++ {
		end := len(input)
		min := utility.Min(input)
		if checkZeros(input) {
			continue
		}
		for j := 0; j < end; j++ {
			input[j] -= min
		}
	}
	return len(input)
}
