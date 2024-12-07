package main

import (
	"aoc2024/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := util.ReadLines("input.txt")
	input := parseInput(lines)

	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}

type calibration struct {
	target int64
	vals   []int
}

func parseInput(lines []string) []calibration {
	res := make([]calibration, 0, len(lines))
	for _, line := range lines {
		c := calibration{}
		parts := strings.Split(line, ": ")
		c.target, _ = strconv.ParseInt(parts[0], 10, 64)
		c.vals = util.ParseInts(parts[1], " ")
		res = append(res, c)
	}
	return res
}

func (c calibration) couldBeTrue() bool {
	//fmt.Println("calibrating: ", c)
	return eval(c.vals[0], int(c.target), c.vals[1:])
}

func (c calibration) couldBeTrue2() bool {
	//fmt.Println("calibrating: ", c)
	return eval2(c.vals[0], int(c.target), c.vals[1:])
}

func eval(current, target int, remaining []int) bool {
	//fmt.Printf("current: %d, target: %d, remaining: %v\n", current, target, remaining)
	if len(remaining) == 0 {
		return current == target
	}

	return eval(current+remaining[0], target, remaining[1:]) || eval(current*remaining[0], target, remaining[1:])
}

func eval2(current, target int, remaining []int) bool {
	//fmt.Printf("current: %d, target: %d, remaining: %v\n", current, target, remaining)
	if len(remaining) == 0 {
		return current == target
	}

	concat, _ := strconv.ParseInt(fmt.Sprintf("%d%d", current, remaining[0]), 10, 64)
	return eval2(current+remaining[0], target, remaining[1:]) || eval2(current*remaining[0], target, remaining[1:]) || eval2(int(concat), target, remaining[1:])
}

func part1(input []calibration) int {
	total := 0
	for _, c := range input {
		if c.couldBeTrue() {
			total += int(c.target)
		}
	}
	return total
}

func part2(input []calibration) int {
	total := 0
	for _, c := range input {
		if c.couldBeTrue2() {
			total += int(c.target)
		}
	}
	return total
}
