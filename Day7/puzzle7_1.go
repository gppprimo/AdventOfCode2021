package day7

import (
	utility "AoC-2021/Utility"
	"math"
)

func Puzzle7_1(input []int) int {
	min := utility.Min(input)
	max := utility.Max(input)

	lessFuel := math.MaxInt
	for i := min; i <= max; i++ {
		sum := 0
		for _, e := range input {
			sum += int(math.Abs(float64(e - i)))
		}
		if sum < lessFuel {
			lessFuel = sum
		}
	}
	return lessFuel
}

func Puzzle7_2(input []int) int {
	min := utility.Min(input)
	max := utility.Max(input)
	lessFuel := math.MaxInt
	for i := min; i <= max; i++ {
		sum := 0
		for _, e := range input {
			n := int(math.Abs(float64(e - i)))
			sum += (n * (n + 1)) / 2
		}
		if sum < lessFuel {
			lessFuel = sum
		}
	}
	return lessFuel
}
