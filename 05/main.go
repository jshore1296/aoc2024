package main

import (
	"aoc2024/util"
	"fmt"
	"sort"
)

func main() {
	lines := util.ReadLines("input.txt")
	rules, updates := parseInput(lines)
	fmt.Println("Part 1: ", part1(rules, updates))
	fmt.Println("Part 2: ", part2(rules, updates))
}

func parseInput(lines []string) (map[int][]int, [][]int) {
	rules := make(map[int][]int)
	updates := make([][]int, 0)

	i := 0
	for {
		// parse rules
		if lines[i] == "" {
			i++
			break
		}
		vals := util.ParseInts(lines[i], "|")
		before, after := vals[0], vals[1]
		rules[before] = append(rules[before], after)
		i++
	}

	for ; i < len(lines); i++ {
		updates = append(updates, util.ParseInts(lines[i], ","))
	}
	return rules, updates
}

func isValidUpdate(rules map[int][]int, update []int) bool {
	seen := make(map[int]bool)

	for _, val := range update {
		mustComeAfter := rules[val]
		for _, badVal := range mustComeAfter {
			if seen[badVal] {
				return false
			}
		}
		seen[val] = true
	}
	return true
}

func part1(rules map[int][]int, updates [][]int) int {
	validUpdates := make([][]int, 0)
	for _, update := range updates {
		if isValidUpdate(rules, update) {
			validUpdates = append(validUpdates, update)
		}
	}

	total := 0
	for _, update := range validUpdates {
		total += update[len(update)/2]
	}
	return total
}

func part2(rules map[int][]int, updates [][]int) int {
	invalidUpdates := make([][]int, 0)
	for _, update := range updates {
		if !isValidUpdate(rules, update) {
			invalidUpdates = append(invalidUpdates, update)
		}
	}

	total := 0
	for _, update := range invalidUpdates {
		//fmt.Println(update)
		sort.Slice(update, func(i, j int) bool {
			// i < j if j appears in its rule list
			rule := rules[update[i]]
			for _, val := range rule {
				if val == update[j] {
					return true
				}
			}
			return false
		})
		//fmt.Println("Sorted: ", update)
		total += update[len(update)/2]
	}

	return total
}
