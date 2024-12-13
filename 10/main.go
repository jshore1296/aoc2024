package main

import (
	"aoc2024/util"
	"fmt"
)

func main() {
	lines := util.ReadLines("input.txt")

	tm := parseInput(lines)
	fmt.Println("Part 1: ", part1(tm))
	fmt.Println("Part 2: ", part2(tm))
}

type topoMap struct {
	grid [][]int

	uniq  map[[2]int]bool
	total int
}

func parseInput(lines []string) topoMap {
	tm := topoMap{
		grid: make([][]int, len(lines)),
	}

	for rowIdx, row := range lines {
		tm.grid[rowIdx] = util.ParseInts(row, "")
	}

	return tm
}

func (tm topoMap) nextIndices(rowIdx, colIdx int) [][2]int {
	val := tm.grid[rowIdx][colIdx]
	res := make([][2]int, 0)
	if rowIdx > 0 {
		if tm.grid[rowIdx-1][colIdx] == val+1 {
			res = append(res, [2]int{rowIdx - 1, colIdx})
		}
	}
	if rowIdx < len(tm.grid)-1 {
		if tm.grid[rowIdx+1][colIdx] == val+1 {
			res = append(res, [2]int{rowIdx + 1, colIdx})
		}
	}
	if colIdx > 0 {
		if tm.grid[rowIdx][colIdx-1] == val+1 {
			res = append(res, [2]int{rowIdx, colIdx - 1})
		}
	}
	if colIdx < len(tm.grid[0])-1 {
		if tm.grid[rowIdx][colIdx+1] == val+1 {
			res = append(res, [2]int{rowIdx, colIdx + 1})
		}
	}
	return res
}

func (tm topoMap) score1(rowIdx, colIdx int) int {
	tm.uniq = make(map[[2]int]bool)
	tm.total = 0

	tm.innerScore(rowIdx, colIdx)
	return len(tm.uniq)
}

func (tm topoMap) score2(rowIdx, colIdx int) int {
	tm.uniq = make(map[[2]int]bool)
	tm.total = 0

	tm.innerScore(rowIdx, colIdx)
	return tm.total
}

func (tm *topoMap) innerScore(rowIdx, colIdx int) {
	if tm.grid[rowIdx][colIdx] == 9 {
		tm.uniq[[2]int{rowIdx, colIdx}] = true
		tm.total++
		return
	}
	for _, nextIdx := range tm.nextIndices(rowIdx, colIdx) {
		tm.innerScore(nextIdx[0], nextIdx[1])
	}
}

func part1(tm topoMap) int {
	total := 0

	for rowIdx, row := range tm.grid {
		for colIdx, val := range row {
			if val == 0 {
				s := tm.score1(rowIdx, colIdx)
				total += s
			}
		}
	}
	return total
}

func part2(tm topoMap) int {
	total := 0

	for rowIdx, row := range tm.grid {
		for colIdx, val := range row {
			if val == 0 {
				s := tm.score2(rowIdx, colIdx)
				total += s
			}
		}
	}
	return total
}
