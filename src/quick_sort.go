package src

import (
	"fmt"

	"github.com/urfave/cli"
)

func quickSort(c *cli.Context) error {
	quickSortCase([]int{2, 3, 1})
	quickSortCase([]int{9, 8, 7, 6, 5, 4, 3, 2, 1})
	quickSortCase([]int{8, 7, 6, 5, 4, 3, 2, 1})
	quickSortCase([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	quickSortCase([]int{1, 1, 1, 1, 1, 1, 1, 1, 1})
	quickSortCase([]int{1, 1, 1, 1, 1, 1, 1, 1})
	quickSortCase([]int{4, 4, 4, 4, 3, 3, 3, 2, 2, 1})
	quickSortCase([]int{6, 8, 2, 5, 1, 9, 3, 4, 7})
	return nil
}

func quickSortCase(array []int) {
	fmt.Println("origin: ", array)
	array = quickSortImp(array)
	fmt.Println("sorted: ", array)
}

func quickSortImp(array []int) []int {
	quickSortRecur(array, 0, len(array)-1)
	return array
}

func quickSortRecur(array []int, left, right int) {
	if left >= right {
		return
	}
	middle := (left + right) / 2
	array[right], array[middle] = array[middle], array[right]

	i, j := left, right-1
	for i <= j {
		if array[i] <= array[right] {
			i++
		} else {
			array[i], array[j] = array[j], array[i]
			j--
		}
	}
	array[i], array[right] = array[right], array[i]

	quickSortRecur(array, left, i-1)
	quickSortRecur(array, i+1, right)
}
