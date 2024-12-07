package main

import (
	"aoc2024/util"
	"fmt"
)

func main() {
	lines := util.ReadLines("input.txt")

	fmt.Println("Part 1: ", part1(wordSearch(lines)))
	fmt.Println("Part 2: ", part2(wordSearch(lines)))
}

type wordSearch []string

var lettersExpected = []byte{'M', 'A', 'S'}

func (w wordSearch) xMas(row, col int, rowDiff, colDiff int) bool {
	if w[row][col] != 'X' {
		return false
	}

	for i := 0; i < 3; i++ {
		row += rowDiff
		col += colDiff

		if row < 0 || row >= len(w) {
			// off the board
			return false
		}
		if col < 0 || col >= len(w[0]) {
			// off the board
			return false
		}
		if w[row][col] != lettersExpected[i] {
			return false
		}
	}

	return true
}

func (w wordSearch) findXmasFrom(row, col int) (found int) {
	if w[row][col] != 'X' {
		return
	}

	for _, rowDiff := range []int{-1, 0, 1} {
		for _, colDiff := range []int{-1, 0, 1} {
			if w.xMas(row, col, rowDiff, colDiff) {
				found++
			}
		}
	}

	return found
}

func part1(ws wordSearch) int {
	total := 0

	for rowIdx, row := range ws {
		for colIdx := range row {
			total += ws.findXmasFrom(rowIdx, colIdx)
		}
	}

	return total
}

func (w wordSearch) isXMasV2(row, col int) bool {
	if w[row][col] != 'A' {
		return false
	}
	if row == 0 || row == len(w)-1 {
		return false
	}
	if col == 0 || col == len(w[0])-1 {
		return false
	}

	// no X or A allowed at the diagonals
	for _, rowDiff := range []int{-1, 1} {
		for _, colDiff := range []int{-1, 1} {
			val := w[row+rowDiff][col+colDiff]
			if val == 'X' || val == 'A' {
				return false
			}
		}
	}

	// all M or S's left, now we just have to make sure the two diagonal pairs aren't equal, which implies one M and one S
	topLeft := w[row-1][col-1]
	topRight := w[row-1][col+1]
	bottomLeft := w[row+1][col-1]
	bottomRight := w[row+1][col+1]

	if topLeft == bottomRight {
		return false
	}

	if topRight == bottomLeft {
		return false
	}

	return true

}

func part2(ws wordSearch) int {
	total := 0

	for rowIdx, row := range ws {
		for colIdx := range row {
			if ws.isXMasV2(rowIdx, colIdx) {
				total++
			}
		}
	}

	return total
}
