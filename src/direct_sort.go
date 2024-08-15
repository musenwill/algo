package src

import (
	"fmt"

	"github.com/urfave/cli"
)

func directSort(c *cli.Context) error {
	directSortCase([]int{9, 8, 7, 6, 5, 4, 3, 2, 1})
	directSortCase([]int{8, 7, 6, 5, 4, 3, 2, 1})
	directSortCase([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	directSortCase([]int{1, 1, 1, 1, 1, 1, 1, 1, 1})
	directSortCase([]int{1, 1, 1, 1, 1, 1, 1, 1})
	directSortCase([]int{4, 4, 4, 4, 3, 3, 3, 2, 2, 1})
	directSortCase([]int{6, 8, 2, 5, 1, 9, 3, 4, 7})
	return nil
}

func directSortCase(array []int) {
	fmt.Println("origin: ", array)
	array = directSortImp(array)
	fmt.Println("sorted: ", array)
}

func directSortImp(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array
}
