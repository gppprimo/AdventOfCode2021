package day6

func Puzzle6_1(input []int) int {
	res := input
	for loop := 0; loop < 80; loop++ {
		end := len(res)

		for j := 0; j < end; j++ {
			if res[j] == 0 {
				res = append(res, 8)
				res[j] = 6
				continue
			}
			res[j] -= 1
		}
	}
	return len(res)
}
