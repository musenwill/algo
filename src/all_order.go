package src

import (
	"fmt"

	"github.com/urfave/cli"
)

func allOrder(c *cli.Context) error {
	allOrderCase("abc")
	fmt.Println()
	allOrderCase("abcd")
	return nil
}

func allOrderCase(set string) {
	allOrderImp([]rune(set), 0, len(set)-1)
}

func allOrderImp(set []rune, firstIdx, lastIdx int) {
	if firstIdx >= lastIdx {
		fmt.Println(string(set))
		return
	}

	for i := firstIdx; i <= lastIdx; i++ {
		set[firstIdx], set[i] = set[i], set[firstIdx]
		allOrderImp(set, firstIdx+1, lastIdx)
		set[firstIdx], set[i] = set[i], set[firstIdx]
	}
}
