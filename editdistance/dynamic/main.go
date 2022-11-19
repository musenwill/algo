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

	line0 := []byte(lines[0])
	line1 := []byte(lines[1])
	fmt.Println(distance(line0, line1))
}

func distance(a, b []byte) int {
	cur := make([]int, len(a)+1)
	for i := 0; i < len(cur); i++ {
		cur[i] = i
	}

	var a1b1 int 
	for i := 0; i < len(b); i++ {
		a1b1 = cur[0]
		for j := 0; j < len(cur); j++ {
			if j == 0 {
				cur[0] = cur[0] + 1
			} else {
				var d int
				if b[i] == a[j-1] {
					d = a1b1
				} else {
					ab1 := cur[j]
					a1b := cur[j-1]
					d = min(a1b1, ab1, a1b) + 1
				}
				a1b1 = cur[j]
				cur[j] = d
			}
		}
	}

	return cur[len(cur)-1]
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