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
	fmt.Println(treeLayerSum(node))
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

func treeLayerSum(root *model.TreeNode) []int {
	result := make([]int, 0)

	if root == nil {
		return result
	}

	layerA := make([]*model.TreeNode, 0)

	layerA = append(layerA, root)
	for len(layerA) > 0 {
		sum := 0
		layerB := make([]*model.TreeNode, 0)
		for _, node := range layerA {
			sum += int(node.Val)
			if node.Left != nil {
				layerB = append(layerB, node.Left)
			}
			if node.Right != nil {
				layerB = append(layerB, node.Right)
			}
		}
		result = append(result, sum)
		layerA = layerB
	}
	return result
}
