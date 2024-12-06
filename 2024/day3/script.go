package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Filepath is a required argument for this script")
		return
	}

	dataSet, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Failed to read the given file", os.Args[1], err)
		return
	}

	sumOfMultiplications := calculateSumOfMultiplications(dataSet)
	fmt.Println("Sum of multiplications:", sumOfMultiplications)
}

func calculateSumOfMultiplications(dataSet []byte) int {
	var total int = 0

	pattern := regexp.MustCompile(`mul\((\d+,\d+)\)`)
	matches := pattern.FindAllSubmatch(dataSet, -1)

	for _, match := range matches {
		if len(match) < 1 {
			continue
		}

		numbers := strings.Split(string(match[1]), ",")

		total += multiplyNumbersInStringSlice(numbers)
	}

	return total
}

func multiplyNumbersInStringSlice(numbersAsString []string) int {
	if len(numbersAsString) < 2 {
		return 0
	}

	var multiplication int = 1

	for _, numberAsString := range numbersAsString {
		number, err := strconv.Atoi(numberAsString)
		if err != nil {
			continue
		}

		multiplication = multiplication * number
	}

	return multiplication
}
