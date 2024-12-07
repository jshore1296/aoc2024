package main

import (
	"aoc2024/util"
	"fmt"
)

func main() {
	lines := util.ReadLines("input.txt")
	w := parseInput(lines)

	fmt.Println("Part 1: ", part1(w))
}

type world struct {
	grid          [][]rune // rows, then columns
	guardPosition [2]int   // row, then col.

	//row, then col e.g. -1, 0 = north; 0, 1 = east; 1, 0 = south; 0, -1 = west
	guardDirection [2]int
}

func (w *world) turnRight() {
	w.guardDirection[0], w.guardDirection[1] = w.guardDirection[1], w.guardDirection[0]*-1
}

func (w *world) markGuardSpot() {
	w.grid[w.guardPosition[0]][w.guardPosition[1]] = 'X'
}

func (w *world) advance() {
	newRow := w.guardPosition[0] + w.guardDirection[0]
	newCol := w.guardPosition[1] + w.guardDirection[1]
	if w.onGrid(newRow, newCol) && w.grid[newRow][newCol] == '#' {
		w.turnRight()
	} else {
		w.markGuardSpot()
		w.guardPosition = [2]int{newRow, newCol}
	}
}

func (w *world) guardOnGrid() bool {
	return w.onGrid(w.guardPosition[0], w.guardPosition[1])
}

func (w *world) onGrid(row, col int) bool {
	if row < 0 || row >= len(w.grid) {
		return false
	}
	if col < 0 || col >= len(w.grid[0]) {
		return false
	}
	return true
}

func (w *world) spotsVisited() int {
	total := 0
	for _, row := range w.grid {
		for _, val := range row {
			if val == 'X' {
				total++
			}
		}
	}
	return total
}

func parseInput(lines []string) *world {
	res := &world{
		grid: make([][]rune, len(lines)),
	}
	for rowIdx, row := range lines {
		for colIdx, data := range row {
			res.grid[rowIdx] = append(res.grid[rowIdx], data)
			if data == '^' {
				res.grid[rowIdx][colIdx] = 'X'
				res.guardPosition[0] = rowIdx
				res.guardPosition[1] = colIdx
				res.guardDirection = [2]int{-1, 0}
			}
		}
	}
	return res
}

func part1(w *world) int {
	for w.guardOnGrid() {
		w.advance()
	}
	return w.spotsVisited()
}
