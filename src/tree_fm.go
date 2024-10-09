package src

import (
	"fmt"

	"github.com/musenwill/algo/model"
	"github.com/musenwill/algo/util"
	"github.com/urfave/cli"
)

// construct binary tree from first and middle root order
func treeFM(c *cli.Context) error {
	treeFMCase("abdecf", "dbeafc")
	fmt.Println()
	treeFMCase("abcde", "edcba")
	fmt.Println()
	treeFMCase("abcdef", "edcbaf")
	return nil
}

func treeFMCase(firstRoot, middleRoot string) {
	util.Assert(len(firstRoot) == len(middleRoot))
	node := treeFMImp([]rune(firstRoot), []rune(middleRoot), 0, len(firstRoot)-1, 0, len(middleRoot)-1)

	fmt.Println(node.FirstRootDump())
	fmt.Println(node.MiddleRootDump())
	fmt.Println(node.LastRootDump())
}

func treeFMImp(firstRoot, middleRoot []rune, firstRootFirstIdx, firstRootLastIdx,
	middleRootFirstIdx, middleRootLastIdx int) *model.TreeNode {

	util.Assert(firstRootLastIdx-firstRootFirstIdx == middleRootLastIdx-middleRootFirstIdx)
	node := &model.TreeNode{
		Val:   firstRoot[firstRootFirstIdx],
		Left:  nil,
		Right: nil,
	}

	if firstRootFirstIdx == firstRootLastIdx {
		return node
	}

	rootIdx := 0
	for i := middleRootFirstIdx; i <= middleRootLastIdx; i++ {
		if middleRoot[i] == firstRoot[firstRootFirstIdx] {
			rootIdx = i
			break
		}
	}

	leftMiddleRootFirstIdx := middleRootFirstIdx
	leftMiddleRootLastIdx := rootIdx - 1
	rightMiddleRootFirstIdx := rootIdx + 1
	rightMiddleRootLastIdx := middleRootLastIdx

	leftFirstRootFirstIdx := firstRootFirstIdx + 1
	leftFirstRootLastIdx := leftFirstRootFirstIdx + leftMiddleRootLastIdx - leftMiddleRootFirstIdx
	rightFirstRootFirstIdx := leftFirstRootLastIdx + 1
	rightFirstRootLastIdx := firstRootLastIdx

	if leftMiddleRootLastIdx >= leftMiddleRootFirstIdx {
		node.Left = treeFMImp(firstRoot, middleRoot, leftFirstRootFirstIdx, leftFirstRootLastIdx, leftMiddleRootFirstIdx, leftMiddleRootLastIdx)
	}
	if rightMiddleRootLastIdx >= rightMiddleRootFirstIdx {
		node.Right = treeFMImp(firstRoot, middleRoot, rightFirstRootFirstIdx, rightFirstRootLastIdx, rightMiddleRootFirstIdx, rightMiddleRootLastIdx)
	}

	return node
}
