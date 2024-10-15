package src

import (
	"fmt"

	"github.com/urfave/cli"
)

func editDistance(c *cli.Context) error {
	editDistanceCase("", "")
	editDistanceCase("apple", "")
	editDistanceCase("", "append")
	editDistanceCase("apple", "append")
	editDistanceCase("apple", "apple")
	editDistanceCase("abcde", "acde")
	editDistanceCase("abcde", "abcdef")
	editDistanceCase("abcde", "abcdd")
	return nil
}

func editDistanceCase(wordA, wordB string) {
	fmt.Printf("%d = editDistance(%s, %s)\n", editDistanceImp([]rune(wordA), []rune(wordB)), wordA, wordB)
}

func editDistanceImp(wordA, wordB []rune) int {
	if len(wordA) == 0 {
		return len(wordB)
	}
	if len(wordB) == 0 {
		return len(wordA)
	}

	distances := make([]int, len(wordB))

	pre, left := 0, 0
	for _, w := range wordA {
		for j, v := range wordB {
			up := distances[j]
			min := minOfTree(pre, left, up)
			if w != v {
				min++
			}
			pre = distances[j]
			distances[j] = min
			left = distances[j]
		}
	}

	return distances[len(wordB)-1]
}

func minOfTree(a, b, c int) int {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}
