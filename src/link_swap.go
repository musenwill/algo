package src

import (
	"fmt"

	"github.com/musenwill/algo/model"
	"github.com/urfave/cli"
)

func swapLink(c *cli.Context) error {
	swapLinkCase(nil)
	swapLinkCase(model.BuildSingleLink(1))
	swapLinkCase(model.BuildSingleLink(1, 2))
	swapLinkCase(model.BuildSingleLink(1, 2, 3))
	swapLinkCase(model.BuildSingleLink(1, 2, 3, 4))
	swapLinkCase(model.BuildSingleLink(1, 2, 3, 4, 5))
	return nil
}

func swapLinkCase(head *model.SingleLinkNode) {
	fmt.Printf("origin : %s\n", head.Dump())
	head = swapLinkImp(head)
	fmt.Printf("swapped: %s\n", head.Dump())
}

func swapLinkImp(head *model.SingleLinkNode) *model.SingleLinkNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next

	var pre *model.SingleLinkNode = nil
	nodeA := head
	nodeB := head.Next
	for nodeA != nil && nodeB != nil {
		next := nodeB.Next
		nodeA.Next = next
		nodeB.Next = nodeA
		if pre != nil {
			pre.Next = nodeB
		}

		pre = nodeA
		nodeA = next
		if nodeA != nil {
			nodeB = nodeA.Next
		} else {
			nodeB = nil
		}
	}

	return newHead
}
