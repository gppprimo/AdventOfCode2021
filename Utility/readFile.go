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
