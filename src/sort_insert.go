package src

import (
	"fmt"

	"github.com/urfave/cli"
)

func insertSort(c *cli.Context) error {
	insertSortCase([]int{9, 8, 7, 6, 5, 4, 3, 2, 1})
	insertSortCase([]int{8, 7, 6, 5, 4, 3, 2, 1})
	insertSortCase([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	insertSortCase([]int{1, 1, 1, 1, 1, 1, 1, 1, 1})
	insertSortCase([]int{1, 1, 1, 1, 1, 1, 1, 1})
	insertSortCase([]int{4, 4, 4, 4, 3, 3, 3, 2, 2, 1})
	insertSortCase([]int{6, 8, 2, 5, 1, 9, 3, 4, 7})
	return nil
}

func insertSortCase(array []int) {
	fmt.Println("origin: ", array)
	array = insertSortImp(array)
	fmt.Println("sorted: ", array)
}

func insertSortImp(array []int) []int {
	for i := 1; i < len(array); i++ {
		for j := 0; j < i; j++ {
			if array[j] > array[i] {
				tmp := array[j]
				array[j] = array[i]
				for ; j < i; j++ {
					array[j+1], tmp = tmp, array[j+1]
				}
				break
			}
		}
	}
	return array
}
