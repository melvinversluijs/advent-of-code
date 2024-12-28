package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("File path is required as input for this script.")
		return
	}

	dataSetInBytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Failed reading file", os.Args[1], err)
		return
	}

	dataSet := strings.TrimSpace(string(dataSetInBytes[:]))

	totalXmasMatches := calculateXmasMatchCount(dataSet)
	fmt.Println("Total amount XMAS matches:", totalXmasMatches)

	totalCrossMasMatches := calculateCrossMasMatches(dataSet)
	fmt.Println("Total amount of X-MAS matches:", totalCrossMasMatches)
}

func calculateXmasMatchCount(dataSet string) int {
	total := 0
	rows := strings.Split(dataSet, "\n")
	maxY := len(rows)

	// Loop vertically
	for y := 0; y < maxY; y++ {
		maxX := len(rows[y])

		// Loop horizontally
		for x := 0; x < maxX; x++ {
			// Start finding matches from starting point X
			if !isX(rows[y][x]) {
				continue
			}

			// Right
			if canGoRight(x, maxX, 3) && rows[y][x:x+4] == "XMAS" {
				total++
			}

			// Left
			if canGoLeft(x, 3) && rows[y][x-3:x+1] == "SAMX" {
				total++
			}

			// Down
			if canGoDown(y, maxY, 3) && isM(rows[y+1][x]) && isA(rows[y+2][x]) && isS(rows[y+3][x]) {
				total++
			}

			// Up
			if canGoUp(y, 3) && isM(rows[y-1][x]) && isA(rows[y-2][x]) && isS(rows[y-3][x]) {
				total++
			}

			// Diagonal up / right
			if canGoUp(y, 3) && canGoRight(x, maxX, 3) && isM(rows[y-1][x+1]) && isA(rows[y-2][x+2]) && isS(rows[y-3][x+3]) {
				total++
			}

			// Diagonal up / left
			if canGoUp(y, 3) && canGoLeft(x, 3) && isM(rows[y-1][x-1]) && isA(rows[y-2][x-2]) && isS(rows[y-3][x-3]) {
				total++
			}

			// Diagonal down / right
			if canGoDown(y, maxY, 3) && canGoRight(x, maxX, 3) && isM(rows[y+1][x+1]) && isA(rows[y+2][x+2]) && isS(rows[y+3][x+3]) {
				total++
			}

			// Diagonal down / left
			if canGoDown(y, maxY, 3) && canGoLeft(x, 3) && isM(rows[y+1][x-1]) && isA(rows[y+2][x-2]) && isS(rows[y+3][x-3]) {
				total++
			}
		}
	}

	return total
}

func calculateCrossMasMatches(dataSet string) int {
	total := 0
	rows := strings.Split(dataSet, "\n")
	maxY := len(rows)

	// Loop vertically
	for y := 0; y < maxY; y++ {
		maxX := len(rows[y])

		// Loop horizontally
		for x := 0; x < maxX; x++ {
			// Start finding matches from starting point A
			if !isA(rows[y][x]) {
				continue
			}

			// Make sure there is room available 1 spot around the starting point, since we need a cross.
			if !canGoUp(y, 1) || !canGoDown(y, maxY, 1) || !canGoRight(x, maxX, 1) || !canGoLeft(x, 1) {
				continue
			}

			// M.S
			// .A.
			// M.S
			if isM(rows[y-1][x-1]) && isS(rows[y+1][x+1]) && isM(rows[y+1][x-1]) && isS(rows[y-1][x+1]) {
				total++
			}

			// S.S
			// .A.
			// M.M
			if isS(rows[y-1][x-1]) && isM(rows[y+1][x+1]) && isM(rows[y+1][x-1]) && isS(rows[y-1][x+1]) {
				total++
			}

			// M.M
			// .A.
			// S.S
			if isM(rows[y-1][x-1]) && isS(rows[y+1][x+1]) && isS(rows[y+1][x-1]) && isM(rows[y-1][x+1]) {
				total++
			}

			// S.M
			// .A.
			// S.M
			if isS(rows[y-1][x-1]) && isM(rows[y+1][x+1]) && isS(rows[y+1][x-1]) && isM(rows[y-1][x+1]) {
				total++
			}
		}
	}

	return total
}

func canGoRight(x int, maxX int, distance int) bool {
	return x+distance < maxX
}

func canGoLeft(x int, distance int) bool {
	return x-distance >= 0
}

func canGoDown(y int, maxY int, distance int) bool {
	return y+distance < maxY
}

func canGoUp(y int, distance int) bool {
	return y-distance >= 0
}

func isX(b byte) bool {
	return b == byte('X')
}

func isM(b byte) bool {
	return b == byte('M')
}

func isA(b byte) bool {
	return b == byte('A')
}

func isS(b byte) bool {
	return b == byte('S')
}
