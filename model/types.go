package model

import (
	"fmt"
	"strings"
)

type SingleLinkNode struct {
	Val  int
	Next *SingleLinkNode
}

func (L *SingleLinkNode) Dump() string {
	var sb strings.Builder

	for nil != L {
		sb.WriteString(fmt.Sprintf("%d -> ", L.Val))
		L = L.Next
	}
	sb.WriteString("nil")
	return sb.String()
}

func BuildSingleLink(values ...int) *SingleLinkNode {
	var head *SingleLinkNode = nil
	var cur *SingleLinkNode = nil

	for _, val := range values {
		node := new(SingleLinkNode)
		node.Val = val
		node.Next = nil

		if head == nil {
			head = node
			cur = node
		} else {
			cur.Next = node
			cur = node
		}
	}

	return head
}
