package day8

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func convertDigit(str string) int {
	res, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func differences(letters []byte, str string) string {
	tmp := ""
	bitArr := make([]int, len(letters), len(letters))
	for i := range letters {
		for j := range str {
			if letters[i] == str[j] {
				bitArr[i] += 1
				break
			}
		}
	}
	for i, ba := range bitArr {
		if ba == 0 {
			tmp += string(letters[i])
		}
	}
	return tmp
}

func checkDigit(s string, notIn map[string]int) int {
	letters := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	switch len(s) {
	case 2:
		return 1
	case 3:
		return 7
	case 4:
		return 4
	case 7:
		return 8
	default:
		tmp := differences(letters, s)
		fmt.Println(tmp)
		return notIn[tmp]
	}
}

func Puzzle8_2(input []string) int {
	notIn := map[string]int{
		"ag": 5,
		"ga": 5,
		"eb": 2,
		"be": 2,
		"eg": 3,
		"ge": 3,
		"g":  9,
		"a":  6,
		"f":  0,
	}

	count := 0
	input2 := []string{"fdgacbe", "cefdb", "cefbgd", "gcbe"}
	for _, s := range input2 {
		strArr := strings.Split(strings.TrimLeft(s, " "), " ")
		res := ""
		for _, str := range strArr {
			res += string(checkDigit(str, notIn) + '0')
		}
		fmt.Printf("res: %v\n", res)
		count += convertDigit(res)
	}
	return count
}
