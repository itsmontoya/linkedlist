package linkedlist

func newNode(parent *LinkedList, prev, next *Node, val interface{}) *Node {
	return &Node{parent, prev, next, val}
}

// Node is a value container
type Node struct {
	parent *LinkedList

	prev *Node
	next *Node

	val interface{}
}

// Next will return the next Node
func (n *Node) Next() (nn *Node) {
	if n == nil {
		return
	}

	n.parent.read(func() {
		nn = n.next
	})

	return
}

// Prev will return the previous Node
func (n *Node) Prev() (pn *Node) {
	if n == nil {
		return
	}

	n.parent.read(func() {
		pn = n.prev
	})

	return
}

// Val will return the value of a Node
// Note: If your value is a pointer, please do not modify the value.
// 	If you need to modify the value, use the Update func
func (n *Node) Val() (val interface{}) {
	n.parent.read(func() {
		val = n.val
	})

	return
}

// Update will update the value of a particular Node
func (n *Node) Update(val interface{}) {
	n.parent.write(func() {
		n.val = val
	})
}
