package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString(0)
	lines := strings.Split(string(input), "\n")

	disMap := make(map[string]int)
	line0 := []byte(lines[0])
	line1 := []byte(lines[1])
	distance(disMap, line0, line1)
	fmt.Println(disMap[fmt.Sprintf("%d-%d", len(line0), len(line1))])
}

func distance(disMap map[string]int, a, b []byte) {
	if len(a) * len(b) == 0 {
		disMap[fmt.Sprintf("%d-%d", len(a), len(b))] = max(len(a), len(b))
		return
	}

	if _, ok := disMap[fmt.Sprintf("%d-%d", len(a)-1, len(b)-1)]; !ok {
		distance(disMap, a[0:len(a)-1], b[0:len(b)-1])
	}
	da1b1 := disMap[fmt.Sprintf("%d-%d", len(a)-1, len(b)-1)]

	if a[len(a)-1] == b[len(b)-1] {
		disMap[fmt.Sprintf("%d-%d", len(a), len(b))] = da1b1
	} else {
		if _, ok := disMap[fmt.Sprintf("%d-%d", len(a)-1, len(b))]; !ok {
			distance(disMap, a[0:len(a)-1], b[0:len(b)])
		}
		da1b := disMap[fmt.Sprintf("%d-%d", len(a)-1, len(b))]

		if _, ok := disMap[fmt.Sprintf("%d-%d", len(a), len(b)-1)]; !ok {
			distance(disMap, a[0:len(a)], b[0:len(b)-1])
		}
		dab1 := disMap[fmt.Sprintf("%d-%d", len(a), len(b)-1)]

		disMap[fmt.Sprintf("%d-%d", len(a), len(b))] = min(da1b1, da1b, dab1) + 1
	}
}

func max(ns ...int) int {
	m := ns[0]

	for _, n := range ns {
		if n > m {
			m = n
		}
	}
	return m
}

func min(ns ...int) int {
	m := ns[0]

	for _, n := range ns {
		if n < m {
			m = n
		}
	}
	return m
}