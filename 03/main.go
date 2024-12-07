package main

import (
	"aoc2024/util"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	input := util.ReadLines("input.txt")[0]
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}

var (
	mulInstr      = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	mulOrDoOrDont = regexp.MustCompile(`(?:mul\((\d+),(\d+)\)|do\(\)|don't\(\))`)
)

func part1(input string) int {
	matches := mulInstr.FindAllStringSubmatch(input, -1)

	answer := 0
	for _, match := range matches {
		x, _ := strconv.ParseInt(match[1], 10, 64)
		y, _ := strconv.ParseInt(match[2], 10, 64)
		answer += int(x * y)

	}
	return answer
}

func part2(input string) int {
	matches := mulOrDoOrDont.FindAllStringSubmatch(input, -1)

	enabled := true
	answer := 0
	for _, match := range matches {
		if match[0] == "do()" {
			enabled = true
			continue
		} else if match[0] == "don't()" {
			enabled = false
			continue
		} else if enabled {
			x, _ := strconv.ParseInt(match[1], 10, 64)
			y, _ := strconv.ParseInt(match[2], 10, 64)
			answer += int(x * y)
		}
	}
	return answer
}
