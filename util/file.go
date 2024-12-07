package util

import (
	"bufio"
	"os"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadLines(file string) []string {
	f, err := os.Open(file)
	checkErr(err)

	scanner := bufio.NewScanner(f)
	out := make([]string, 0)

	for scanner.Scan() {
		out = append(out, scanner.Text())
	}

	return out
}
