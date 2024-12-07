package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(name string) [][]int {
	file, err := os.Open(name)
	checkErr(err)

	scanner := bufio.NewScanner(file)
	lists := [][]int{
		{},
		{},
	}
	for scanner.Scan() {
		checkErr(scanner.Err())
		line := scanner.Text()
		vals := strings.Split(line, "   ")
		for i := 0; i < 2; i++ {
			val, err := strconv.ParseInt(vals[i], 10, 64)
			checkErr(err)
			lists[i] = append(lists[i], int(val))
		}
	}

	sort.Ints(lists[0])
	sort.Ints(lists[1])
	//fmt.Println(lists)

	return lists
}

func distance(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func part1(lists [][]int) int {
	total := 0
	for i := 0; i < len(lists[0]); i++ {
		total += distance(lists[0][i], lists[1][i])
	}
	return total
}

func part2(lists [][]int) int {
	frequency := map[int]int{}
	for _, val := range lists[1] {
		frequency[val] = frequency[val] + 1
	}

	similarity := 0
	for _, val := range lists[0] {
		similarity += val * frequency[val]
	}
	return similarity
}

func main() {
	lists := readFile("input.txt")
	fmt.Println(part1(lists))
	fmt.Println(part2(lists))
}
