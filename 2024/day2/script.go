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

	safeReportsCount, safeReportsCountWithDampener := calculateSafeReportsCount(dataSet)
	fmt.Println("Safe reports count:", safeReportsCount)
	fmt.Println("Safe reports count with dampener:", safeReportsCountWithDampener)
}

func calculateSafeReportsCount(dataSet string) (int, int) {
	var total int = 0
	var totalWithDampener = 0

	reports := strings.Split(dataSet, "\n")
	for _, report := range reports {
		if report == "" {
			continue
		}

		values := convertStringSliceToIntSlice(strings.Split(report, " "))

		// If there are less than 2 items in the report, the report can be skipped
		if len(values) < 2 {
			continue
		}

		if isSafeReport(values) {
			total++
			continue
		}

		if isSafeWithDampener(values) {
			totalWithDampener++
		}
	}

	return total, total + totalWithDampener
}

func isSafeReport(report []int) bool {
	reportDescending := slices.Clone(report)
	sort.Sort(sort.Reverse(sort.IntSlice(reportDescending)))

	// If the values are not the same as the ascending or descending sorted values it is a false report.
	if !slices.IsSorted(report) && !slices.Equal(report, reportDescending) {
		return false
	}

	// Check that the steps between the values are at least 1 and max 3.
	for i := 1; i < len(reportDescending); i++ {
		step := reportDescending[i-1] - reportDescending[i]

		if step < 1 || step > 3 {
			return false
		}
	}

	return true
}

func isSafeWithDampener(report []int) bool {
	for index := range report {
		clonedReport := slices.Clone(report)
		dampenedReport := append(clonedReport[:index], clonedReport[index+1:]...)

		// If a dampened version is safe, return true.
		if isSafeReport(dampenedReport) {
			return true
		}
	}

	return false
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
