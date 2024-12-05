package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Invalid amount of arguments given, expected filepath as input")
		return
	}

	dataSetInBytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Failed to read the given file", os.Args[1], err)
	}

	dataSet := string(dataSetInBytes[:])

	safeReportsCount := calculateSafeReportsCount(dataSet)
	fmt.Println("Safe reports count:", safeReportsCount)
}

func calculateSafeReportsCount(dataSet string) int {
	var total int = 0

	reports := strings.Split(dataSet, "\n")
	for _, report := range reports {
		if isSafeReport(report) {
			total++
		}
	}

	return total
}

func isSafeReport(report string) bool {
	if report == "" {
		return false
	}

	values := convertStringSliceToIntSlice(strings.Split(report, " "))

	// If there are less than 2 items in the report, they are safe
	if len(values) < 2 {
		return true
	}

	valuesDescending := slices.Clone(values)
	sort.Sort(sort.Reverse(sort.IntSlice(valuesDescending)))

	// If the values are not the same as the ascending or descending sorted values it is a false report.
	if !slices.IsSorted(values) && !slices.Equal(values, valuesDescending) {
		return false
	}

	// Check that the steps between the values are at least 1 and max 3.
	for i := 1; i < len(valuesDescending); i++ {
		step := valuesDescending[i-1] - valuesDescending[i]

		if step < 1 || step > 3 {
			return false
		}
	}

	return true
}

func convertStringSliceToIntSlice(stringSlice []string) []int {
	var intSlice []int

	for _, valueAsString := range stringSlice {
		value, err := strconv.Atoi(valueAsString)
		if err != nil {
			continue
		}

		intSlice = append(intSlice, value)
	}

	return intSlice
}
