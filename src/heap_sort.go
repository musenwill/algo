package src

import (
	"fmt"

	"github.com/urfave/cli"
)

func heapSort(c *cli.Context) error {
	heapSortCase([]int{9, 8, 7, 6, 5, 4, 3, 2, 1})
	heapSortCase([]int{8, 7, 6, 5, 4, 3, 2, 1})
	heapSortCase([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	heapSortCase([]int{1, 1, 1, 1, 1, 1, 1, 1, 1})
	heapSortCase([]int{1, 1, 1, 1, 1, 1, 1, 1})
	heapSortCase([]int{4, 4, 4, 4, 3, 3, 3, 2, 2, 1})
	heapSortCase([]int{6, 8, 2, 5, 1, 9, 3, 4, 7})
	return nil
}

func heapSortCase(array []int) {
	fmt.Println("origin: ", array)
	array = heapSortImp(array)
	fmt.Println("sorted: ", array)
}

func heapSortImp(array []int) []int {
	array = heapBuild(array)
	sorted := make([]int, 0, len(array))
	var popped int

	for len(array) > 0 {
		array, popped = heapPop(array)
		sorted = append(sorted, popped)
	}

	return sorted
}

func heapBuild(array []int) []int {
	for heapIdx := 2; heapIdx <= len(array); heapIdx++ {
		for treeIdx := heapIdx; treeIdx > 1; treeIdx >>= 1 {
			if array[treeIdx-1] < array[treeIdx/2-1] {
				array[treeIdx-1], array[treeIdx/2-1] = array[treeIdx/2-1], array[treeIdx-1]
			} else {
				break
			}
		}
	}
	return array
}

func heapPop(array []int) ([]int, int) {
	heapLen := len(array)
	popped := array[0]

	parent := 1
	for parent <= heapLen {
		left, right := parent*2, parent*2+1
		if right <= heapLen && array[right-1] < array[left-1] {
			array[parent-1] = array[right-1]
			parent = right
		} else if left <= heapLen {
			array[parent-1] = array[left-1]
			parent = left
		} else {
			break
		}
	}

	return array[:heapLen-1], popped
}
