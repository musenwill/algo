package src

import (
	"fmt"

	"github.com/musenwill/algo/model"
	"github.com/urfave/cli"
)

func lru(c *cli.Context) error {
	// put 1
	// put 2
	// put 3
	// get 2
	// put 4
	// get 1
	// put 5
	// put 6
	// get 2
	// put 7
	// get 6
	// finally: 6 -> 7 -> 2 -> 5 -> 1

	lruC := newLruCache(5)
	lruC.Put(1, 100)
	fmt.Println("put (1, 100)")
	lruC.Get(1)
	fmt.Println("get 1")
	lruC.Put(2, 200)
	fmt.Println("put (2, 200)")
	fmt.Println(lruC.head.Dump())
	lruC.Put(3, 300)
	fmt.Println("put (3, 300)")
	fmt.Println(lruC.head.Dump())
	lruC.Get(2)
	fmt.Println("get 2")
	fmt.Println(lruC.head.Dump())
	lruC.Put(4, 400)
	fmt.Println("put (4, 400)")
	fmt.Println(lruC.head.Dump())
	lruC.Get(1)
	fmt.Println("get 1")
	fmt.Println(lruC.head.Dump())
	lruC.Put(5, 500)
	fmt.Println("put (5, 500)")
	fmt.Println(lruC.head.Dump())
	lruC.Put(6, 600)
	fmt.Println("put (6, 600)")
	fmt.Println(lruC.head.Dump())
	lruC.Get(2)
	fmt.Println("get 2")
	fmt.Println(lruC.head.Dump())
	lruC.Put(7, 700)
	fmt.Println("put (7, 700)")
	fmt.Println(lruC.head.Dump())
	lruC.Get(6)
	fmt.Println("get 6")
	fmt.Println(lruC.head.Dump())

	return nil
}

type lruCache struct {
	cap   int
	cache map[int]*model.DoubleLinkNode
	head  *model.DoubleLinkNode
	tail  *model.DoubleLinkNode
}

func newLruCache(cap int) *lruCache {
	return &lruCache{
		cap:   cap,
		cache: make(map[int]*model.DoubleLinkNode),
		head:  nil,
	}
}

func (L *lruCache) Put(key, val int) {
	if L.cache[key] == nil {
		node := &model.DoubleLinkNode{
			Key:  key,
			Val:  val,
			Prev: nil,
			Next: nil,
		}
		L.cache[key] = node

		if L.head == nil {
			L.head = node
			L.tail = node
		} else {
			L.head.Prev = node
			node.Next = L.head
			L.head = node
		}

		if len(L.cache) > L.cap {
			node := L.tail
			node.Prev.Next = nil
			node.Prev = nil
			node.Next = nil
			delete(L.cache, node.Key)
		}
	} else {
		node := L.cache[key]
		L.moveToHead(node)
	}
}

func (L *lruCache) Get(key int) int {
	node := L.cache[key]
	if node == nil {
		return 0
	}
	L.moveToHead(node)
	return node.Val
}

func (L *lruCache) moveToHead(node *model.DoubleLinkNode) {
	pre := node.Prev
	next := node.Next
	if pre != nil {
		pre.Next = next
		node.Prev = nil
		node.Next = L.head
		L.head.Prev = node
		L.head = node
	}
	if next != nil {
		if pre != nil {
			next.Prev = pre
		}
	} else {
		if pre == nil {
			L.tail = node
		} else {
			L.tail = pre
		}
	}
}
