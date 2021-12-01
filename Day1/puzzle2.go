package day1

func Puzzle2(input []int) int {
	count := 0
	for i := 0; i < len(input)-3; i++ {
		if input[i] < input[i+3] {
			count += 1
		}
	}
	return count
}
