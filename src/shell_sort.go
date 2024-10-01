package src

import (
	"fmt"

	"github.com/urfave/cli"
)

func shellSort(c *cli.Context) error {
	shellSortCase([]int{9, 8, 7, 6, 5, 4, 3, 2, 1})
	shellSortCase([]int{8, 7, 6, 5, 4, 3, 2, 1})
	shellSortCase([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	shellSortCase([]int{1, 1, 1, 1, 1, 1, 1, 1, 1})
	shellSortCase([]int{1, 1, 1, 1, 1, 1, 1, 1})
	shellSortCase([]int{4, 4, 4, 4, 3, 3, 3, 2, 2, 1})
	shellSortCase([]int{6, 8, 2, 5, 1, 9, 3, 4, 7})
	return nil
}

func shellSortCase(array []int) {
	fmt.Println("origin: ", array)
	array = shellSortImp(array)
	fmt.Println("sorted: ", array)
}

// shell 排序等价于许多个小的插入排序
func shellSortImp(array []int) []int {
	for step := len(array) / 2; step > 0; step >>= 1 {
		for i := step; i < len(array); i++ {
			tmp := array[i]
			j := i
			for ; j >= step; j -= step {
				if array[j-step] > tmp {
					array[j] = array[j-step]
				} else {
					break
				}
			}
			array[j] = tmp
		}
	}

	return array
}
