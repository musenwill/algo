package src

import (
	"fmt"

	"github.com/musenwill/algo/model"
	"github.com/urfave/cli"
)

func mergeLink(c *cli.Context) error {
	mergeLinkCase(nil, nil)
	mergeLinkCase(nil, model.BuildSingleLink(1, 2))
	mergeLinkCase(model.BuildSingleLink(1, 2), nil)
	mergeLinkCase(model.BuildSingleLink(1, 3, 5, 7, 9), model.BuildSingleLink(2, 4, 6, 8))
	mergeLinkCase(model.BuildSingleLink(1, 2, 3, 4), model.BuildSingleLink(5, 6, 7, 8))
	mergeLinkCase(model.BuildSingleLink(5, 6, 7, 8), model.BuildSingleLink(1, 2, 3, 4))
	mergeLinkCase(model.BuildSingleLink(1, 1, 3, 3), model.BuildSingleLink(2, 2, 3, 3))
	return nil
}

func mergeLinkCase(linkA, linkB *model.SingleLinkNode) {
	fmt.Printf("linkA: %s\n", linkA.Dump())
	fmt.Printf("linkB: %s\n", linkB.Dump())
	merged := mergeLinkImp(linkA, linkB)
	fmt.Printf("merged: %s\n", merged.Dump())
}

func mergeLinkImp(linkA, linkB *model.SingleLinkNode) *model.SingleLinkNode {
	var newHead *model.SingleLinkNode = nil
	var newLink *model.SingleLinkNode = nil

	for linkA != nil && linkB != nil {
		if linkA.Val <= linkB.Val {
			if newLink == nil {
				newHead = linkA
				newLink = linkA
			} else {
				newLink.Next = linkA
				newLink = linkA
			}
			linkA = linkA.Next
		} else {
			if newLink == nil {
				newHead = linkB
				newLink = linkB
			} else {
				newLink.Next = linkB
				newLink = linkB
			}
			linkB = linkB.Next
		}
		newLink.Next = nil
	}
	if linkA != nil {
		if newLink == nil {
			newHead = linkA
		} else {
			newLink.Next = linkA
		}
	} else {
		if newLink == nil {
			newHead = linkB
		} else {
			newLink.Next = linkB
		}
	}
	return newHead
}
