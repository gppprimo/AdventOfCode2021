package main

import (
	"fmt"
	"log"
	"os"
	"AoC-2021/Utility"
	"AoC-2021/Day1"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// DAY 1
	input := utility.ReadNumbers(path + "/Day1/" + "input1.txt")
	fmt.Printf("Puzzle 1.1: %d\n", day1.Puzzle1(input))
	fmt.Printf("Puzzle 1.2: %d\n", day1.Puzzle2(input))

	// DAY 2
}
