package main

import (
	day1 "AoC-2021/Day1"
	day2 "AoC-2021/Day2"
	day3 "AoC-2021/Day3"
	day4 "AoC-2021/Day4"
	day5 "AoC-2021/Day5"
	day6 "AoC-2021/Day6"

	utility "AoC-2021/Utility"
	"fmt"
	"log"
	"os"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// DAY 1
	fmt.Println("--- Day 1 ---")
	inputInt := utility.ReadNumbers(path + "/Day1/" + "input1.1.txt")
	fmt.Printf("Puzzle 1.1: %d\n", day1.Puzzle1_1(inputInt))
	fmt.Printf("Puzzle 1.2: %d\n", day1.Puzzle1_2(inputInt))

	// DAY 2
	fmt.Println("\n--- Day 2 ---")
	inputString := utility.ReadStrings(path + "/Day2/" + "input2.1.txt")
	inputActionValue := utility.GetActionValue(inputString)
	fmt.Printf("Puzzle 2.1: %d\n", day2.Puzzle2_1(inputActionValue))
	fmt.Printf("Puzzle 2.2: %d\n", day2.Puzzle2_2(inputActionValue))

	// DAY 3
	fmt.Println("\n--- Day 3 ---")
	inputString = utility.ReadStrings(path + "/Day3/" + "input3.1.txt")
	fmt.Printf("Puzzle 3.1: %d\n", day3.Puzzle3_1(inputString))
	fmt.Printf("Puzzle 3.2: %d\n", day3.Puzzle3_2(inputString))

	// DAY4
	fmt.Println("\n--- Day 4 ---")
	rnds, boards := utility.Reading4_1(utility.ReadStrings(path + "/Day4/" + "input4.1.txt"))
	fmt.Printf("Puzzle 4.1: %d\n", day4.Puzzle4_1(rnds, boards))
	fmt.Printf("Puzzle 4.2: %d\n", day4.Puzzle4_2(rnds, boards))

	// DAY 5
	fmt.Println("\n--- Day 5 ---")
	structs5_1 := utility.Reading5_1(utility.ReadStrings(path + "/Day5/" + "test.txt"))
	fmt.Printf("Puzzle 5.1: %d\n", day5.Puzzle5_1(structs5_1))
	fmt.Printf("Puzzle 5.1: %d\n", day5.Puzzle5_2(structs5_1))


	// DAY 6
	fmt.Println("\n--- Day 6 ---")
	inputInt = utility.Reading6_1(utility.ReadStrings(path + "/Day6/" + "test.txt")[0])
	fmt.Printf("Puzzle 6.1: %d\n", day6.Puzzle6_1(inputInt))
	fmt.Printf("Puzzle 6.2: %d\n", day6.Puzzle6_2(inputInt))

	// DAY 7
	fmt.Println("\n--- Day 7 ---")
	// inputInt = utility.Reading6_1(utility.ReadStrings(path + "/Day7/" + "input7.1.txt")[0])
	// fmt.Printf("Puzzle 7.1: %d\n", day7.Puzzle7_1(inputInt))
	// fmt.Printf("Puzzle 7.2: %d\n", day7.Puzzle7_2(inputInt))

	// DAY 8
	fmt.Println("\n--- Day 8 ---")
	// inputString = utility.Reading8_1(utility.ReadStrings(path + "/Day8/" + "input8.1.txt"))
	// fmt.Printf("Puzzle 8.1: %d\n", day8.Puzzle8_1(inputString))
	// fmt.Printf("Puzzle 8.2: %d\n", day8.Puzzle8_2(inputString))

	// DAY 9
	fmt.Println("\n--- Day 9 ---")
	// inputMat := utility.Reading9_1(utility.ReadStrings(path + "/Day9/" + "input9.1.txt"))
	// fmt.Printf("Puzzle 9.1: %d\n", day9.Puzzle9_1(inputMat))
	// fmt.Printf("Puzzle 9.2: %d\n", day8.Puzzle8_2(inputString))

	// DAY 13
	fmt.Println("\n--- Day 13 ---")
	// inputPaper, inputFolds := utility.Reading13_1(utility.ReadStrings(path + "/Day13/" + "input13.1.txt"))
	// fmt.Printf("Puzzle 13.1: %d\n", day13.Puzzle13_1(inputPaper, inputFolds))
	// fmt.Printf("Puzzle 9.2: %d\n", day8.Puzzle8_2(inputString))
}
