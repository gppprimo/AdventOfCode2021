package day3

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
)

// 1 -> 1 is mostCommon
// 0 -> 0 is mostCommon
// 2 -> 1 and 0 are equally common
func mostCommonVal(input []string, index int) int {
	count := 0
	for _, binStr := range input {
		if binStr[index]-'0' == 1 {
			count += 1
		}
	}

	if len(input)-count > count {
		return 0
	} else if len(input)-count < count {
		return 1
	}
	return 2
}

func lessCommonVal(input []string, index int) int {
	if mostCommonVal(input, index) == 1 {
		return 0
	} else if mostCommonVal(input, index) == 0 {
		return 1
	}
	return 2
}

func findBinaryString(input []string, commonVal func([]string, int) int) int {
	fmt.Println(runtime.FuncForPC(reflect.ValueOf(commonVal).Pointer()).Name())
	tempInputStr := input
	for i := 0; i < len(input[0]); i++ {
		if len(tempInputStr) == 1 {
			break
		}
		tempInputStr2 := make([]string, 0)
		common := commonVal(tempInputStr, i)
		fmt.Printf("Bin n. %d. Common %d\n", i+1, common)
		for _, binStr := range tempInputStr {
			if int(binStr[i]-'0') == common {
				tempInputStr2 = append(tempInputStr2, binStr)
			} else if common == 2 {
				if (runtime.FuncForPC(reflect.ValueOf(commonVal).Pointer()).Name() == "AoC-2021/Day3.mostCommonVal" && int(binStr[i]-'0') == 1) || 
				(runtime.FuncForPC(reflect.ValueOf(commonVal).Pointer()).Name() == "AoC-2021/Day3.lessCommonVal" && int(binStr[i]-'0') == 0) {
					tempInputStr2 = append(tempInputStr2, binStr)
				}
			}
		}
		tempInputStr = tempInputStr2
	}
	oxy, _ := strconv.ParseInt(tempInputStr[0], 2, 16)
	return int(oxy)
}

func computeOxygen(input []string) int {
	return findBinaryString(input, mostCommonVal)
}

func computeCO2(input []string) int {
	return findBinaryString(input, lessCommonVal)
}

func Puzzle3_2(input []string) int {
	oxygenGeneratorRating := computeOxygen(input)
	CO2ScrubberRating := computeCO2(input)

	return oxygenGeneratorRating * CO2ScrubberRating
}
