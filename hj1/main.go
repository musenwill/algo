package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()
	fmt.Println(lastWordLen(string(line)))
}

func lastWordLen(str string) int {
	runes := []rune(str)
	wordLen := 0

	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] == rune(' ') {
			if wordLen > 0 {
				return wordLen
			} else {
				continue
			}
		} else {
			wordLen++
		}
	}

	return wordLen
}
