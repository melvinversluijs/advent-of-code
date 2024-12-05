package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Invalid amount of arguments given, expected one filepath as input")
		return
	}

	dataSetInBytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Failed to read the given file", err)
		return
	}

	dataSet := string(dataSetInBytes[:])
	leftNumbers, rightNumbers := splitDataSetIntoTwoArrays(dataSet)

	totalDistanceBetweenNumbers := calculateTotalDistanceBetweenNumbers(leftNumbers, rightNumbers)
	fmt.Println("Total distance between numbers:", totalDistanceBetweenNumbers)

	similarityScore := calculateSimilarityScore(leftNumbers, rightNumbers)
	fmt.Println("Similarity score:", similarityScore)
}

func splitDataSetIntoTwoArrays(dataSet string) ([]int, []int) {
	lines := strings.Split(dataSet, "\n")

	var leftNumbers, rightNumbers []int
	for _, line := range lines {
		numbers := strings.Split(line, "   ")
		leftNumber, leftError := strconv.Atoi(numbers[0])
		rightNumber, rightError := strconv.Atoi(numbers[1])

		if leftError != nil || rightError != nil {
			continue
		}

		leftNumbers = append(leftNumbers, leftNumber)
		rightNumbers = append(rightNumbers, rightNumber)
	}

	return leftNumbers, rightNumbers
}

func calculateTotalDistanceBetweenNumbers(leftNumbers []int, rightNumbers []int) int {
	var total int = 0

	// Clone numbers because they are by reference and we do not want to sort the originals.
	clonedLeftNumbers := cloneIntSlice(leftNumbers)
	clonedRightNumbers := cloneIntSlice(rightNumbers)

	sort.Ints(clonedLeftNumbers)
	sort.Ints(clonedRightNumbers)

	for index, leftNumber := range clonedLeftNumbers {
		rightNumber := clonedRightNumbers[index]
		lowestNumber := min(leftNumber, rightNumber)
		highestNumber := max(leftNumber, rightNumber)

		total += highestNumber - lowestNumber
	}

	return total
}

func calculateSimilarityScore(leftNumbers []int, rightNumbers []int) int {
	var total int = 0

	ocurrences := getOccurrencesCountPerNumber(rightNumbers)

	for _, number := range leftNumbers {
		occurenceCount, countExists := ocurrences[number]
		if countExists {
			total += number * occurenceCount
		}
	}

	return total
}

func cloneIntSlice(slice []int) []int {
	clonedSlice := make([]int, len(slice))

	copy(clonedSlice, slice)

	return clonedSlice
}

func getOccurrencesCountPerNumber(numbers []int) map[int]int {
	occurenceCountPerNumber := make(map[int]int)

	for _, number := range numbers {
		occurenceCountPerNumber[number]++
	}

	return occurenceCountPerNumber
}
