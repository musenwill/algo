package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString(0)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		line = strings.ToLower(strings.TrimSpace(line))
		if len(line) == 0 {
			continue
		}
		fmt.Println(toDecimal(line))
	}
}

func toDecimal(hex string) int {
	runes := []rune(hex)
	runes = runes[2:]

	sum := 0
	for _, r := range runes {
		sum *= 16
		sum += HEX[r]
	}
	return sum
}

var HEX = map[rune]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'a': 10,
	'b': 11,
	'c': 12,
	'd': 13,
	'e': 14,
	'f': 15,
}
