package utility

import "math"

func Min(input []int) int {
	min := math.MaxInt
	for _, e := range input {
		if e < min {
			min = e
		}
	}
	return min
}

func Max(input []int) int {
	max := math.MinInt
	for _, e := range input {
		if e > max {
			max = e
		}
	}
	return max
}
