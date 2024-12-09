package main

import (
	"aoc2024/util"
	"fmt"
	"strconv"
)

func main() {
	line := util.ReadLines("input.txt")[0]

	fmt.Println("Part 1: ", part1(line))
	fmt.Println("Part 2: ", part2(line))
}

func explode(line string) []int {
	result := []int{}

	label := 0
	for i, c := range line {
		length, _ := strconv.ParseInt(string(c), 10, 64)
		if i%2 == 0 {
			// file
			for j := 0; j < int(length); j++ {
				result = append(result, label)
			}
			label++
		} else {
			// free space
			for j := 0; j < int(length); j++ {
				result = append(result, -1)
			}
		}
	}
	return result
}

func resort(data []int) {
	start := 0
	end := len(data) - 1
	for start < end {
		if data[start] > -1 {
			start++
			continue
		}
		if data[end] == -1 {
			end--
			continue
		}
		data[start], data[end] = data[end], data[start]
		start++
		end--
	}
}

func checksum(data []int) int {
	total := 0

	for i, val := range data {
		if val == -1 {
			continue
		}
		total += i * val
	}
	return total
}

func part1(line string) int {
	data := explode(line)
	resort(data)

	return checksum(data)
}

func getLargestFileNumber(data []int) int {
	for i := len(data) - 1; i > 0; i-- {
		if data[i] != -1 {
			return data[i]
		}
	}
	panic("nope")
}

func getFileIdx(data []int, fileNum int) (idx, length int) {
	for i, val := range data {
		if val == fileNum {
			if idx == 0 {
				idx = i
			}
			length++
		}
	}
	return idx, length
}

func findEmptySpace(data []int, length int) (idx int) {
	currentLength := 0
	for i := 0; i < len(data); i++ {
		if data[i] == -1 {
			if idx == 0 {
				idx = i
			}
			currentLength++
			if currentLength >= length {
				return idx
			}
		} else {
			idx = 0
			currentLength = 0
		}
	}
	return 0
}

func packFiles(data []int) {
	largestFile := getLargestFileNumber(data)

	for i := largestFile; i > 0; i-- {
		idx, length := getFileIdx(data, i)

		blankIdx := findEmptySpace(data, length)
		if blankIdx != 0 && blankIdx < idx {
			for j := 0; j < length; j++ {
				data[idx], data[blankIdx] = data[blankIdx], data[idx]
				idx++
				blankIdx++
			}
		}
	}
}

func part2(line string) int {
	data := explode(line)

	packFiles(data)

	return checksum(data)
}
