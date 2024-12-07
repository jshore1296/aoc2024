package util

import (
	"strconv"
	"strings"
)

func ParseInts(in, sep string) (out []int) {
	for _, str := range strings.Split(in, sep) {
		i, err := strconv.ParseInt(str, 10, 64)
		checkErr(err)
		out = append(out, int(i))
	}
	return out
}
