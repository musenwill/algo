package src

import (
	"fmt"

	"github.com/urfave/cli"
)

func topK(c *cli.Context) error {
	array := []int{1, 4, 6, 2, 3, 8, 9, 5, 7, 10}
	result := topKImp(array, 3)
	fmt.Println(array)
	fmt.Println(result)

	return nil
}

func topKImp(array []int, k int) []int {
	heap := TopKHeapCreate(k)
	result := make([]int, 0, k)

	for _, n := range array {
		heap.Put(n)
	}
	for heap.Len() > 0 {
		result = append(result, heap.Pop())
	}
	return result
}

type topKHeap struct {
	cap   int
	array []int
}

func TopKHeapCreate(cap int) *topKHeap {
	return &topKHeap{
		cap:   cap,
		array: make([]int, 0, cap),
	}
}

func (h *topKHeap) Put(n int) {
	if len(h.array) >= h.cap {
		if n <= h.array[0] {
			return
		}
		h.array[0] = n

		curIdx := 1
		for {
			left, right := curIdx*2, curIdx*2+1
			if right <= len(h.array) && h.array[right-1] < h.array[left-1] && h.array[right-1] < h.array[curIdx-1] {
				h.array[right-1], h.array[curIdx-1] = h.array[curIdx-1], h.array[right-1]
				curIdx = right
			} else if left <= len(h.array) && h.array[left-1] < h.array[curIdx-1] {
				h.array[left-1], h.array[curIdx-1] = h.array[curIdx-1], h.array[left-1]
				curIdx = left
			} else {
				break
			}
		}
	} else {
		h.array = append(h.array, n)
		curIdx := len(h.array)
		for curIdx >= 2 {
			parent := curIdx / 2
			if h.array[curIdx-1] < h.array[parent-1] {
				h.array[curIdx-1], h.array[parent-1] = h.array[parent-1], h.array[curIdx-1]
				curIdx = parent
			} else {
				break
			}
		}
	}
}

func (h *topKHeap) Pop() int {
	result := h.array[0]

	curIdx := 1
	for {
		left, right := curIdx*2, curIdx*2+1
		if right <= len(h.array) && h.array[right-1] < h.array[left-1] {
			h.array[curIdx-1] = h.array[right-1]
			curIdx = right
		} else if left <= len(h.array) {
			h.array[curIdx-1] = h.array[left-1]
			curIdx = left
		} else {
			break
		}
	}
	h.array = h.array[:len(h.array)-1]

	return result
}

func (h *topKHeap) Len() int {
	return len(h.array)
}
