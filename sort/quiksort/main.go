package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _, _ := reader.ReadLine()
	sort(input)
	fmt.Println(string(input))
}

func sort(input []byte) {
	quickSort(input, 0, len(input)-1)
}

func quickSort(input []byte, left, right int) {
	if left >= right {
		return
	}

	median := input[left]
	i, j := left + 1, right

	for i <= j {
		if input[i] <= median {
			i++
			continue
		}
		if input[j] > median {
			j--
			continue
		}
		input[i], input[j] = input[j], input[i]
	}

	input[left], input[i-1] = input[i-1], input[left]

	quickSort(input, left, i-2)
	quickSort(input, i, right)
}