package src

import (
	"fmt"

	"github.com/musenwill/algo/model"
	"github.com/urfave/cli"
)

func circleLink(c *cli.Context) error {
	linkHashCircleCase(nil)
	linkHashCircleCase(model.BuildSingleLink(1))
	linkHashCircleCase(model.BuildSingleLink(1, 2))
	linkHashCircleCase(model.BuildSingleLink(1, 2, 3))
	linkHashCircleCase(model.BuildSingleLink(1, 2, 3, 4))
	linkHashCircleCase(model.BuildSingleLink(1, 2, 3, 4, 5, 6, 7, 8, 9))

	{
		head := model.BuildSingleLink(1)
		tail := head.TailNode()
		tail.Next = head
		linkHashCircleCase(head)
	}
	{
		head := model.BuildSingleLink(1, 2)
		tail := head.TailNode()
		tail.Next = head
		linkHashCircleCase(head)
	}
	{
		head := model.BuildSingleLink(1, 2, 3)
		tail := head.TailNode()
		tail.Next = head
		linkHashCircleCase(head)
	}
	{
		head := model.BuildSingleLink(1, 2, 3)
		node := head.NthNode(1)
		tail := head.TailNode()
		tail.Next = node
		linkHashCircleCase(head)
	}
	{
		head := model.BuildSingleLink(1, 2, 3, 4)
		tail := head.TailNode()
		tail.Next = head
		linkHashCircleCase(head)
	}
	{
		head := model.BuildSingleLink(1, 2, 3, 4, 5, 6, 7, 8, 9)
		tail := head.TailNode()
		tail.Next = head
		linkHashCircleCase(head)
	}
	{
		head := model.BuildSingleLink(1, 2, 3, 4, 5, 6, 7, 8, 9)
		node := head.NthNode(1)
		tail := head.TailNode()
		tail.Next = node
		linkHashCircleCase(head)
	}
	{
		head := model.BuildSingleLink(1, 2, 3, 4, 5, 6, 7, 8, 9)
		node := head.NthNode(2)
		tail := head.TailNode()
		tail.Next = node
		linkHashCircleCase(head)
	}
	{
		head := model.BuildSingleLink(1, 2, 3, 4, 5, 6, 7, 8, 9)
		node := head.NthNode(3)
		tail := head.TailNode()
		tail.Next = node
		linkHashCircleCase(head)
	}
	return nil
}

func linkHashCircleCase(head *model.SingleLinkNode) {
	fmt.Printf("has circle %v\n", linkHasCircleImp(head))
}

func linkHasCircleImp(head *model.SingleLinkNode) bool {
	step2 := head
	step3 := head
	hasCircle := false

	for step2 != nil && step3 != nil {
		if step2.Next != nil && step2.Next.Next != nil {
			step2 = step2.Next.Next
		} else {
			break
		}
		if step3.Next != nil && step3.Next.Next != nil && step3.Next.Next.Next != nil {
			step3 = step3.Next.Next.Next
		} else {
			break
		}

		if step2 == step3 {
			hasCircle = true
			break
		}
	}

	return hasCircle
}
