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
	total := calculateMatchCount(dataSet)

	fmt.Println("Total amount of matches:", total)
}

func calculateMatchCount(dataSet string) int {
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
			if canGoRight(x, maxX) && rows[y][x:x+4] == "XMAS" {
				total++
			}

			// Left
			if canGoLeft(x) && rows[y][x-3:x+1] == "SAMX" {
				total++
			}

			// Down
			if canGoDown(y, maxY) && isM(rows[y+1][x]) && isA(rows[y+2][x]) && isS(rows[y+3][x]) {
				total++
			}

			// Up
			if canGoUp(y) && isM(rows[y-1][x]) && isA(rows[y-2][x]) && isS(rows[y-3][x]) {
				total++
			}

			// Diagonal up / right
			if canGoUp(y) && canGoRight(x, maxX) && isM(rows[y-1][x+1]) && isA(rows[y-2][x+2]) && isS(rows[y-3][x+3]) {
				total++
			}

			// Diagonal up / left
			if canGoUp(y) && canGoLeft(x) && isM(rows[y-1][x-1]) && isA(rows[y-2][x-2]) && isS(rows[y-3][x-3]) {
				total++
			}

			// Diagonal down / right
			if canGoDown(y, maxY) && canGoRight(x, maxX) && isM(rows[y+1][x+1]) && isA(rows[y+2][x+2]) && isS(rows[y+3][x+3]) {
				total++
			}

			// Diagonal down / left
			if canGoDown(y, maxY) && canGoLeft(x) && isM(rows[y+1][x-1]) && isA(rows[y+2][x-2]) && isS(rows[y+3][x-3]) {
				total++
			}
		}
	}

	return total
}

func canGoRight(x int, maxX int) bool {
	return x+3 < maxX
}

func canGoLeft(x int) bool {
	return x-3 >= 0
}

func canGoDown(y int, maxY int) bool {
	return y+3 < maxY
}

func canGoUp(y int) bool {
	return y-3 >= 0
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
