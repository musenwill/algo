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
	mergeSort(input, 0, len(input)-1)
}

func mergeSort(input []byte, left, right int) {
	if left >= right {
		return
	}

	middle := (left + right) / 2
	mergeSort(input, left, middle)
	mergeSort(input, middle+1, right)

	// merge
	i, j := left, middle + 1
	for i < j && j <= right {
		if input[i] < input[j] {
			i++
			continue
		}
		tmp := input[j]
		for k := j; k > i; k-- {
			input[k] = input[k-1]
		}
		input[i] = tmp
		i++
		j++
	}
}
