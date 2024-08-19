package src

import (
	"errors"
	"fmt"

	"github.com/urfave/cli"
)

func binarySearch(c *cli.Context) error {
	binarySearchCase([]int{}, 1)
	binarySearchCase([]int{1}, 1)
	binarySearchCase([]int{1, 2}, 1)
	binarySearchCase([]int{1, 2}, 2)
	binarySearchCase([]int{1, 2, 4}, 4)
	binarySearchCase([]int{1, 2, 4}, 2)
	binarySearchCase([]int{1, 2, 4}, 3)
	binarySearchCase([]int{1, 2, 4, 6, 8}, 4)
	return nil
}

func binarySearchCase(array []int, target int) {
	idx, err := binarySearchImp(array, target)
	if nil == err {
		fmt.Printf("(%d: %d) %v\n", target, idx, array)
	} else {
		fmt.Printf("(%d: %s) %v\n", target, err, array)
	}
}

func binarySearchImp(array []int, target int) (int, error) {
	if len(array) == 0 {
		return 0, errors.New("not found")
	}

	left, right := 0, len(array)-1
	middle := (left + right) / 2

	for middle >= 0 && middle <= right && left <= right {
		if array[middle] > target {
			right = middle - 1
			middle = (left + right) / 2
		} else if array[middle] < target {
			left = middle + 1
			middle = (left + right) / 2
		} else {
			return middle, nil
		}
	}

	return 0, errors.New("not found")
}
