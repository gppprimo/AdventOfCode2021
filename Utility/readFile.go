package utility

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadInput(filepath string) []byte {
	var data []byte
	var err error
	data, err = ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	return data
}

func ReadStrings(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var text []string
	for scanner.Scan() {
		text = append(text, strings.TrimRight(scanner.Text(), "\n"))
	}
	return text
}

func ReadNumbers(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer file.Close()

	Scanner := bufio.NewScanner(file)

	var numbers []int
	for Scanner.Scan() {
		numbers = append(numbers, toInt(Scanner.Text()))
	}
	return numbers
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	return result
}

func stringToArrayInt(sSlice string) []int {
	var ints []int
	for _, s := range sSlice {
		ints = append(ints, toInt(string(s)))
	}
	return ints
}

func arrayStringToArrayInt(sSlice []string) []int {
	var ints []int
	for _, s := range sSlice {
		if s != "" {
			ints = append(ints, toInt(string(s)))
		}
	}
	return ints
}


func Reading4_1(input []string) ([]int, [][][]int) {
	randomNumbs := make([]int, len(input[0]))
	inputString := strings.Split(input[0], ",")
	randomNumbs = arrayStringToArrayInt(inputString)

	var boards [][][]int
	var board [][]int

	for i := 2; i < len(input); i++ {
		if input[i] != "" {
			inputString := strings.Split(input[i], " ")
			boardLine := arrayStringToArrayInt(inputString)
			board = append(board, boardLine)
		} else {
			boards = append(boards, board)
			board = make([][]int, 0)
		}
	}
	boards = append(boards, board)
	return randomNumbs, boards
}

func Reading5_1(input []string) []Struct5_1 {
	var result []Struct5_1
	for _, s := range input {
		sSplitted := strings.Split(s, "->")
		sSplitted[0] = strings.TrimRight(sSplitted[0], " ")
		sSplitted[1] = strings.TrimLeft(sSplitted[1], " ")
		lSplitted := strings.Split(sSplitted[0], ",")
		rSplitted := strings.Split(sSplitted[1], ",")
		x1, err := strconv.Atoi(lSplitted[0])
		if err != nil {
			log.Fatal(err)
		}
		y1, err := strconv.Atoi(lSplitted[1])
		if err != nil {
			log.Fatal(err)
		}
		x2, err := strconv.Atoi(rSplitted[0])
		if err != nil {
			log.Fatal(err)
		}
		y2, err := strconv.Atoi(rSplitted[1])
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, Struct5_1{x1, y1, x2, y2})
	}
	return result
}

func Reading6_1(input string) []int {
	inputString := strings.Split(input, ",")
	return arrayStringToArrayInt(inputString)
}

func Reading8_1(input []string) []string {
	var filteredInput []string
	for _, s := range input {
		sFiltered := strings.Split(s, "|")[1]
		filteredInput = append(filteredInput, sFiltered)
	}
	return filteredInput
}

func Reading9_1(input []string) [][]int {
	mat := make([][]int, len(input))
	for i := range input {
		mat[i] = make([]int, len(input[i]))
		mat[i] = stringToArrayInt(input[i])
		fmt.Println(mat[i])
	}
	return mat
}


func Reading13_1(input []string) ([]Struct13_1, []Struct13_1XY) {
	end := 0
	for i := range input {
		if input[i] == "" {
			end = i
		}
	}

	var struct13_1 []Struct13_1
	for i := range input[0:end] {
		tmp := strings.Split(input[i], ",")

		x, err := strconv.Atoi(tmp[0])
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(tmp[1])
		if err != nil {
			log.Fatal(err)
		}
		struct13_1 = append(struct13_1, Struct13_1{x, y})
	}

	var fold []Struct13_1XY
	for i := end+1; i < len(input); i++{
		tmp := strings.Split(input[i], " ")
		tmpXY := strings.Split(tmp[2], "=")
		val, err := strconv.Atoi(tmpXY[1])
		if err != nil {
			log.Fatal(err)
		}
		fold = append(fold, Struct13_1XY{tmpXY[0], val})
	}

	return struct13_1, fold
}
