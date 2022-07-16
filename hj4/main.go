package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()

	for i := 0; i < len(line); i += 8 {
		start := i
		end := i + 8
		if end > len(line) {
			end = len(line)
		}
		subString := string(line[start:end])
		fmt.Print(subString)

		padding := 8 - (end - start)
		if padding > 0 {
			pad := make([]byte, padding)
			for i := 0; i < len(pad); i++ {
				pad[i] = byte('0')
			}
			fmt.Print(string(pad))
		}
		fmt.Println()
	}
}
