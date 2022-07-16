package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString(0)
	lines := strings.Split(data, "\n")

	firstLine := strings.ToLower(lines[0])
	secondLine := strings.ToLower(lines[1])

	c := []rune(secondLine)[0]
	fmt.Println(countChar(firstLine, c))
}

func countChar(str string, c rune) int {
	count := 0
	runes := []rune(str)
	for _, r := range runes {
		if r == c {
			count++
		}
	}
	return count
}
