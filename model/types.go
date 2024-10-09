package model

import (
	"fmt"
	"strings"
)

type SingleLinkNode struct {
	Val  int
	Next *SingleLinkNode
}

func (L *SingleLinkNode) NthNode(n int) *SingleLinkNode {
	count := 0
	for L != nil {
		if count == n {
			return L
		}
		L = L.Next
		count++
	}
	return nil
}

func (L *SingleLinkNode) TailNode() *SingleLinkNode {
	for L != nil && L.Next != nil {
		L = L.Next
	}
	return L
}

func (L *SingleLinkNode) Dump() string {
	var sb strings.Builder

	for L != nil {
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

type DoubleLinkNode struct {
	Key  int
	Val  int
	Prev *DoubleLinkNode
	Next *DoubleLinkNode
}

func (D *DoubleLinkNode) Dump() string {
	var sb strings.Builder

	for D != nil {
		sb.WriteString(fmt.Sprintf("(%d, %d) -> ", D.Key, D.Val))
		D = D.Next
	}
	sb.WriteString("nil")
	return sb.String()
}

type TreeNode struct {
	Val   rune
	Left  *TreeNode
	Right *TreeNode
}

func (T *TreeNode) FirstRootDump() string {
	var sb strings.Builder
	T.firstRootDump(&sb)
	return sb.String()
}

func (T *TreeNode) firstRootDump(sb *strings.Builder) {
	if T == nil {
		return
	}
	sb.WriteRune(T.Val)
	if T.Left != nil {
		T.Left.firstRootDump(sb)
	}
	if T.Right != nil {
		T.Right.firstRootDump(sb)
	}
}

func (T *TreeNode) MiddleRootDump() string {
	var sb strings.Builder
	T.middleRootDump(&sb)
	return sb.String()
}

func (T *TreeNode) middleRootDump(sb *strings.Builder) {
	if T == nil {
		return
	}
	if T.Left != nil {
		T.Left.middleRootDump(sb)
	}
	sb.WriteRune(T.Val)
	if T.Right != nil {
		T.Right.middleRootDump(sb)
	}
}

func (T *TreeNode) LastRootDump() string {
	var sb strings.Builder
	T.lastRootDump(&sb)
	return sb.String()
}

func (T *TreeNode) lastRootDump(sb *strings.Builder) {
	if T == nil {
		return
	}

	if T.Left != nil {
		T.Left.lastRootDump(sb)
	}
	if T.Right != nil {
		T.Right.lastRootDump(sb)
	}
	sb.WriteRune(T.Val)
}
