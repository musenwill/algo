package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString(0)
	lines := strings.Split(string(input), "\n")

	bagMap := make(map[string]int)
	fmt.Println(maxCount(bagMap, bricks(lines[0]), bagSize(lines[1])))
}

func maxCount(bagMap map[string]int, bricks []int, bagSize int) int {
	kna := fmt.Sprintf("%d-%d", len(bricks), bagSize)
	if n, ok := bagMap[kna]; ok {
		return n
	}

	if len(bricks) == 0 || bagSize <= 0 {
		bagMap[kna] = 0
		return 0
	}

	lastBrick := bricks[len(bricks)-1]

	kn1a1 := fmt.Sprintf("%d-%d", len(bricks)-1, bagSize - lastBrick)
	if _, ok := bagMap[kn1a1]; !ok {
		maxCount(bagMap,  bricks[:len(bricks)-1], bagSize - lastBrick)
	}
	n1a1 := bagMap[kn1a1]

	kna1 := fmt.Sprintf("%d-%d", len(bricks)-1, bagSize)
	if _, ok := bagMap[kna1]; !ok {
		maxCount(bagMap,  bricks[:len(bricks)-1], bagSize)
	}
	na1 := bagMap[kna1]

	if bagSize >= lastBrick {
		n1a1 += lastBrick
	} else {
		n1a1 = 0
	}
	n := max(n1a1, na1)
	bagMap[kna] = n

	return n
}

func bricks(line string) []int {
	var ns []int
	ss := strings.Split(line, " ")
	for _, s := range ss {
		b, _ := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
		ns = append(ns ,int(b))
	}
	return ns
}

func bagSize(line string) int {
	b, _ := strconv.ParseInt(strings.TrimSpace(line), 10, 64)
	return int(b)
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