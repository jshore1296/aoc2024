package main

import (
	"aoc2024/util"
	"fmt"
)

func main() {
	lines := util.ReadLines("input.txt")
	nodeLocations := parseInput(lines)

	fmt.Println("Part 1: ", part1(nodeLocations, len(lines), len(lines[0])))
	fmt.Println("Part 2: ", part2(nodeLocations, len(lines), len(lines[0])))

}

func parseInput(lines []string) map[rune][][2]int {
	locations := make(map[rune][][2]int)

	for rowIdx, row := range lines {
		for colIdx, val := range row {
			if val != '.' {
				locations[val] = append(locations[val], [2]int{rowIdx, colIdx})
			}
		}
	}

	return locations
}

func calculateDistances(coords [][2]int, row, col int) [][2]int {
	res := make([][2]int, len(coords))
	for i, coord := range coords {
		res[i] = [2]int{coord[0] - row, coord[1] - col}
	}
	return res
}

func oneDistanceIsDouble(distances [][2]int) bool {
	for _, d1 := range distances {
		for _, d2 := range distances {
			if d1 == d2 {
				continue
			}
			if d1[0]*2 == d2[0] && d1[1]*2 == d2[1] {
				return true
			}
		}
	}
	return false
}

func isAntiNode(nodeLocations map[rune][][2]int, row, col int) bool {
	for _, coords := range nodeLocations {
		distances := calculateDistances(coords, row, col)
		if oneDistanceIsDouble(distances) {
			return true
		}
	}
	return false
}

func part1(nodeLocations map[rune][][2]int, rows, cols int) int {
	total := 0

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if isAntiNode(nodeLocations, row, col) {
				total++
			}
		}
	}

	return total
}

func twoDistancesFormLine(distances [][2]int) bool {
	for _, d1 := range distances {
		for _, d2 := range distances {
			if d1 == d2 {
				continue
			}
			// two distances form a line if x1/y1 == x2/y2, or if x1*y2 == x2*y1
			if d1[0]*d2[1] == d1[1]*d2[0] {
				return true
			}
		}
	}
	return false
}

func isNewAntiNode(nodeLocations map[rune][][2]int, row, col int) bool {
	for _, coords := range nodeLocations {
		distances := calculateDistances(coords, row, col)
		if twoDistancesFormLine(distances) {
			return true
		}
	}
	return false
}

func part2(nodeLocations map[rune][][2]int, rows, cols int) int {
	total := 0

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if isNewAntiNode(nodeLocations, row, col) {
				total++
			}
		}
	}

	return total
}
