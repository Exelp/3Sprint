package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := "hello world hello hello everyone"
	divString(s)
}
func divString(a string) {
	worldCount := make(map[string]int)
	worlds := strings.Split(a, " ")
	for _, i := range worlds {
		worldCount[i]++
	}
	sliceofCount := make([]int, 0, len(worldCount))
	keys := make([]string, 0, len(worldCount))
	for key, value := range worldCount {
		sliceofCount = append(sliceofCount, value)
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(b, c int) bool {
		return sliceofCount[b] > sliceofCount[c]
	})
	for e := range keys {
		fmt.Printf("%s: %d\n", keys[e], sliceofCount[e])
	}
}
