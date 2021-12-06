package day3

import (
	"strconv"
)

func Puzzle3_1(input []string) int {
	var gammaRate string
	var epsilonRate string
	countOnes := make([]int, len(input[0]))
	for _, binStr := range input {
		for i, binDigit := range binStr {
			if binDigit == 49 /* ASCII for 1 */ {
				countOnes[i] += 1
			}
		}
	}

	for _, ones := range countOnes {
		if len(input)-ones >= ones {
			gammaRate += "0"
			epsilonRate += "1"
		} else {
			gammaRate += "1"
			epsilonRate += "0"
		}
	}
	gammaR, _ := strconv.ParseInt(gammaRate, 2, 16)
	epsilonR, _ := strconv.ParseInt(epsilonRate, 2, 16)

	return int(gammaR) * int(epsilonR)
}
