package main

import (
	"aoc2024/util"
	"fmt"
)

func main() {
	lines := util.ReadLines("input.txt")

	data := make([][]int, 0)
	for _, l := range lines {
		data = append(data, util.ParseInts(l, " "))
	}

	fmt.Println("Part 1: ", part1(data))
	fmt.Println("Part 2: ", part2(data))
}

func part1(data [][]int) int {
	safe := 0

	for _, report := range data {
		if isSafe(report) {
			safe++
		}
	}

	return safe
}

func part2(data [][]int) int {
	safe := 0

	for _, report := range data {
		//fmt.Println(report)
		if isSafe(report) {
			safe++
			//fmt.Println("Safe!")
			continue
		}

		for i := range report {
			edited := make([]int, 0)
			edited = append(edited, report[0:i]...)
			edited = append(edited, report[i+1:]...)
			if isSafe(edited) {
				safe++
				//fmt.Println("Safe, edited: ", edited)
				break
			}
		}

	}
	return safe
}

func isSafe(report []int) bool {
	increasing := report[1] > report[0]
	for i := 1; i < len(report); i++ {
		if report[i] == report[i-1] {
			return false
		}
		if report[i] > report[i-1] != increasing {
			return false
		}
		if max(report[i]-report[i-1], report[i-1]-report[i]) > 3 {
			return false
		}
	}
	return true
}
