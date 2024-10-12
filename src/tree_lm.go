package src

import (
	"fmt"

	"github.com/musenwill/algo/model"
	"github.com/musenwill/algo/util"
	"github.com/urfave/cli"
)

// construct binary tree from last and middle root order
func treeLM(c *cli.Context) error {
	treeLMCase("debfca", "dbeafc")
	fmt.Println()
	treeLMCase("edcba", "edcba")
	fmt.Println()
	treeLMCase("edcbfa", "edcbaf")
	return nil
}

func treeLMCase(firstRoot, middleRoot string) {
	util.Assert(len(firstRoot) == len(middleRoot))
	node := treeLMImp([]rune(firstRoot), []rune(middleRoot), 0, len(firstRoot)-1, 0, len(middleRoot)-1)

	fmt.Println(node.FirstRootDump())
	fmt.Println(node.MiddleRootDump())
	fmt.Println(node.LastRootDump())
	fmt.Println(treeLayerSum(node))
}

func treeLMImp(lastRoot, middleRoot []rune, lastRootFirstIdx, lastRootLastIdx,
	middleRootFirstIdx, middleRootLastIdx int) *model.TreeNode {

	util.Assert(lastRootLastIdx-lastRootFirstIdx == middleRootLastIdx-middleRootFirstIdx)
	node := &model.TreeNode{
		Val:   lastRoot[lastRootLastIdx],
		Left:  nil,
		Right: nil,
	}

	if lastRootFirstIdx == lastRootLastIdx {
		return node
	}

	rootIdx := 0
	for i := middleRootFirstIdx; i <= middleRootLastIdx; i++ {
		if middleRoot[i] == lastRoot[lastRootLastIdx] {
			rootIdx = i
			break
		}
	}

	leftMiddleRootFirstIdx := middleRootFirstIdx
	leftMiddleRootLastIdx := rootIdx - 1
	rightMiddleRootFirstIdx := rootIdx + 1
	rightMiddleRootLastIdx := middleRootLastIdx

	leftLastRootFirstIdx := lastRootFirstIdx
	leftLastRootLastIdx := leftLastRootFirstIdx + leftMiddleRootLastIdx - leftMiddleRootFirstIdx
	rightLastRootFirstIdx := leftLastRootLastIdx + 1
	rightLastRootLastIdx := lastRootLastIdx - 1

	if leftMiddleRootLastIdx >= leftMiddleRootFirstIdx {
		node.Left = treeLMImp(lastRoot, middleRoot, leftLastRootFirstIdx, leftLastRootLastIdx, leftMiddleRootFirstIdx, leftMiddleRootLastIdx)
	}
	if rightMiddleRootLastIdx >= rightMiddleRootFirstIdx {
		node.Right = treeLMImp(lastRoot, middleRoot, rightLastRootFirstIdx, rightLastRootLastIdx, rightMiddleRootFirstIdx, rightMiddleRootLastIdx)
	}

	return node
}
