package datastructure

import "sync"

type BTree struct {
	degree int
	length int
	root   *node
	cow    *copyOnWriteContext
}

type node struct {
	items    items
	children children
	cow      *copyOnWriteContext
}

type children []*node

type items []Item

type Item interface {
	Less(than Item) bool
}

type copyOnWriteContext struct {
	freelist *FreeList
}

type FreeList struct {
	mu       sync.Mutex
	freelist []*node
}
