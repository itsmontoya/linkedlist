package linkedlist

func newNode(prev, next *Node, val interface{}) *Node {
	return &Node{prev, next, val}
}

// Node is a value container
type Node struct {
	prev *Node
	next *Node

	val interface{}
}
