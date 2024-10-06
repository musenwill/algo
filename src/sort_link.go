package src

import (
	"fmt"

	"github.com/musenwill/algo/model"
	"github.com/urfave/cli"
)

func sortLink(c *cli.Context) error {
	sortLinkCase(nil)
	sortLinkCase(model.BuildSingleLink(1))
	sortLinkCase(model.BuildSingleLink(1, 2))
	sortLinkCase(model.BuildSingleLink(2, 1))
	sortLinkCase(model.BuildSingleLink(2, 2))
	sortLinkCase(model.BuildSingleLink(3, 2, 5, 4, 1))
	sortLinkCase(model.BuildSingleLink(1, 2, 3, 4, 5))
	sortLinkCase(model.BuildSingleLink(5, 4, 3, 2, 1))
	sortLinkCase(model.BuildSingleLink(3, 2, 2, 2, 1))
	return nil
}

func sortLinkCase(head *model.SingleLinkNode) {
	fmt.Printf("origin: %s\n", head.Dump())
	head = sortLinkImp(head)
	fmt.Printf("sorted: %s\n", head.Dump())
}

func sortLinkImp(head *model.SingleLinkNode) *model.SingleLinkNode {
	newHead := head
	cur := head
	var curPre *model.SingleLinkNode = nil

	for cur != nil {
		tmp := cur.Next
		tmpPre := cur
		for tmp != nil {
			if tmp.Val < cur.Val {
				if cur.Next == tmp {
					cur.Next = tmp.Next
					tmp.Next = cur
				} else {
					tmp.Next, cur.Next = cur.Next, tmp.Next
					tmpPre.Next = cur
				}
				if curPre != nil {
					curPre.Next = tmp
				} else {
					newHead = tmp
				}
				tmp, cur = cur, tmp
			}
			tmpPre = tmp
			tmp = tmp.Next
		}
		curPre = cur
		cur = cur.Next
	}

	return newHead
}
