package linkedlist

import "sync"

// LinkedList is a simple doubly-linked list
type LinkedList struct {
	mux sync.RWMutex

	head *Node
	tail *Node

	len int32
}

func (l *LinkedList) read(fn func()) {
	l.mux.RLock()
	fn()
	l.mux.RUnlock()
}

func (l *LinkedList) write(fn func()) {
	l.mux.Lock()
	fn()
	l.mux.Unlock()
}

// Prepend will prepend the list with a value, the reference Node is Returned
func (l *LinkedList) Prepend(val interface{}) (n *Node) {
	l.write(func() {
		n = newNode(l.head, val)
		if l.tail == nil {
			l.tail = n
		}

		l.head = n
		l.len++
	})

	return
}

// Append will append the list with a value, the reference Node is Returned
func (l *LinkedList) Append(val interface{}) (n *Node) {
	l.write(func() {
		n = newNode(nil, val)

		if l.tail != nil {
			l.tail.next = n
		}

		if l.head == nil {
			l.head = n
		}

		l.tail = n
		l.len++

	})

	return
}

// ForEach will iterate through each Node within the linked list
func (l *LinkedList) ForEach(n *Node, fn ForEachFn) (ended bool) {
	l.read(func() {
		if n == nil {
			// Provided Node is nil, set to head
			n = l.head
		}

		for n != nil {
			if fn(n, n.val) {
				ended = true
				return
			}

			n = n.next
		}
	})

	return
}

// Val will return a Node's value
func (l *LinkedList) Val(n *Node) (val interface{}) {
	l.read(func() {
		val = n.val
	})

	return
}

// Update will update a Node's value
func (l *LinkedList) Update(n *Node, val interface{}) {
	l.write(func() {
		n.val = val
	})
}

// Remove will remove a node
// Note: Removals on a singly linked list are much slower than doubly linked lists
// If you intend to remove on a regular basis, please consider using a doubly linked list
func (l *LinkedList) Remove(n *Node) {
	var pn *Node
	l.ForEach(nil, func(cn *Node, _ interface{}) (end bool) {
		if cn != n {
			pn = cn
			return
		}

		if pn != nil {
			pn.next = n.next
		}

		// Match is found, do stuff
		return true
	})
}

// Len will return the current lenght of the linked list
func (l *LinkedList) Len() (n int32) {
	l.read(func() {
		n = l.len
	})

	return
}

func newNode(next *Node, val interface{}) *Node {
	return &Node{next: next, val: val}
}

// Node is a value container
type Node struct {
	val interface{}

	next *Node
}

// ForEachFn is the format of the function used to call ForEach
type ForEachFn func(n *Node, val interface{}) (end bool)
