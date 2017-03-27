package linkedlist

import "sync"

type single struct {
	mux  sync.RWMutex
	head *Node
	len  int32
}
