// https://adventofcode.com/2021/day/1

package day1

func Puzzle1(input []int) int {

	count := 0
	lastVal := input[0]
	for _, i := range input {
		if i > lastVal {
			count++
		}
		lastVal = i
	}
	return count
}
