package src

import (
	"fmt"

	"github.com/urfave/cli"
)

func mergeSort(c *cli.Context) error {
	mergeSortCase([]int{2, 3, 1})
	mergeSortCase([]int{9, 8, 7, 6, 5, 4, 3, 2, 1})
	mergeSortCase([]int{8, 7, 6, 5, 4, 3, 2, 1})
	mergeSortCase([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	mergeSortCase([]int{1, 1, 1, 1, 1, 1, 1, 1, 1})
	mergeSortCase([]int{1, 1, 1, 1, 1, 1, 1, 1})
	mergeSortCase([]int{4, 4, 4, 4, 3, 3, 3, 2, 2, 1})
	mergeSortCase([]int{6, 8, 2, 5, 1, 9, 3, 4, 7})
	return nil
}

func mergeSortCase(array []int) {
	fmt.Println("origin: ", array)
	array = mergeSortImp(array)
	fmt.Println("sorted: ", array)
}

func mergeSortImp(array []int) []int {
	mergeSortRecur(array, 0, len(array)-1)
	return array
}

func mergeSortRecur(array []int, left, right int) {
	if left >= right {
		return
	}

	middle := (left + right) / 2
	mergeSortRecur(array, left, middle)
	mergeSortRecur(array, middle+1, right)

	i, j := left, middle+1
	for i < j && j <= right {
		if array[i] > array[j] {
			tmp := array[j]
			for k := j; k > i; k-- {
				array[k] = array[k-1]
			}
			array[i] = tmp
			j++
		}
		i++
	}
}
