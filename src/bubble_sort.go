package src

import (
	"fmt"

	"github.com/urfave/cli"
)

func bubbleSort(c *cli.Context) error {
	bubbleSortCase([]int{9, 8, 7, 6, 5, 4, 3, 2, 1})
	bubbleSortCase([]int{8, 7, 6, 5, 4, 3, 2, 1})
	bubbleSortCase([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	bubbleSortCase([]int{1, 1, 1, 1, 1, 1, 1, 1, 1})
	bubbleSortCase([]int{1, 1, 1, 1, 1, 1, 1, 1})
	bubbleSortCase([]int{4, 4, 4, 4, 3, 3, 3, 2, 2, 1})
	bubbleSortCase([]int{6, 8, 2, 5, 1, 9, 3, 4, 7})
	return nil
}

func bubbleSortCase(array []int) {
	fmt.Println("origin: ", array)
	array = bubbleSortImp(array)
	fmt.Println("sorted: ", array)
}

func bubbleSortImp(array []int) []int {
	for i := len(array) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	return array
}
