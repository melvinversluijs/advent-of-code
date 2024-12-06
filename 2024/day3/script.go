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

	sumOfMultiplications, sumOfMultiplicationsWithInstructions := calculateSumOfMultiplications(dataSet)
	fmt.Println("Sum of multiplications:", sumOfMultiplications)
	fmt.Println("Sum of multiplications with instructions:", sumOfMultiplicationsWithInstructions)
}

func calculateSumOfMultiplications(dataSet []byte) (int, int) {
	var total int = 0
	var totalWithInstructions int = 0

	// Retrieve the multiplications and instructions in 3 capture groups.
	// First is: numbers, seperated by comma
	// Second is: do
	// Third is: dont't
	pattern := regexp.MustCompile(`mul\((\d+,\d+)\)|(do)\(\)|(don't)\(\)`)
	matches := pattern.FindAllSubmatch(dataSet, -1)

	enabled := true
	for _, match := range matches {
		numbersString := string(match[1]) // First capture group
		doString := string(match[2])      // Second capture group
		dontString := string(match[3])    // Third capture group

		if doString == "do" {
			enabled = true
			continue
		}

		if dontString == "don't" {
			enabled = false
			continue
		}

		numbers := strings.Split(string(numbersString), ",")

		multiplication := multiplyNumbersInStringSlice(numbers)
		total += multiplication

		if enabled {
			totalWithInstructions += multiplication
		}
	}

	return total, totalWithInstructions
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
