package src

import (
	"fmt"

	"github.com/musenwill/algo/model"
	"github.com/urfave/cli"
)

func reverseLink(c *cli.Context) error {
	reverseLinkCase(nil)
	reverseLinkCase(model.BuildSingleLink(1))
	reverseLinkCase(model.BuildSingleLink(1, 2))
	reverseLinkCase(model.BuildSingleLink(1, 2, 3))
	reverseLinkCase(model.BuildSingleLink(1, 2, 3, 4))
	reverseLinkCase(model.BuildSingleLink(1, 2, 3, 4, 5))
	return nil
}

func reverseLinkCase(head *model.SingleLinkNode) {
	fmt.Printf("origin : %s\n", head.Dump())
	head = reverseLinkImp(head)
	fmt.Printf("reverse: %s\n", head.Dump())
}

func reverseLinkImp(head *model.SingleLinkNode) *model.SingleLinkNode {
	var pre *model.SingleLinkNode = nil
	var next *model.SingleLinkNode = nil
	cur := head

	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
