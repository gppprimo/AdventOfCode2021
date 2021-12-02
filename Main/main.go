package main

import (
	day1 "AoC-2021/Day1"
	day2 "AoC-2021/Day2"
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

}
