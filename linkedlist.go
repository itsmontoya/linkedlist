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

// popHead will pop the head
func (l *LinkedList) popHead() (n *Node) {
	// Set popped node, if popped node is nil - return early
	if n = l.head; n == nil {
		return
	}

	// Set new head, if new head is nil - return early
	if l.head = n.next; l.head == nil {
		return
	}

	// Discard reference to popped node
	l.head.prev = nil
	// Decrement length
	l.len--
	return
}

// popTail will pop the tail
func (l *LinkedList) popTail() (n *Node) {
	// Set popped node, if popped node is nil - return early
	if n = l.tail; n == nil {
		return
	}

	// Set new tail, if new tail is nil - return early
	if l.tail = n.prev; l.tail == nil {
		return
	}

	// Discard reference to popped node
	l.tail.next = nil
	// Decrement length
	l.len--
	return
}

// Prepend will prepend the list with a value, the reference Node is Returned
// Note: If returned Node is nil, append was unsuccessful
func (l *LinkedList) Prepend(val interface{}) (n *Node) {
	l.write(func() {
		n = newNode(l, nil, l.head, val)

		if l.head != nil {
			l.head.prev = n
		}

		if l.tail == nil {
			l.tail = n
		}

		l.head = n
		l.len++
	})

	return
}

// Append will append the list with a value, the reference Node is Returned
// Note: If returned Node is nil, append was unsuccessful
func (l *LinkedList) Append(val interface{}) (n *Node) {
	l.write(func() {
		n = newNode(l, l.tail, nil, val)

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

// ForEachRev will iterate through each Node within the linked list in reverse
func (l *LinkedList) ForEachRev(n *Node, fn ForEachFn) (ended bool) {
	l.read(func() {
		if n == nil {
			// Provided Node is nil, set to tail
			n = l.tail
		}

		for n != nil {
			if fn(n, n.val) {
				ended = true
				return
			}

			n = n.prev
		}
	})

	return
}

// Map will return a mapped list
func (l *LinkedList) Map(fn MapFn) (list []interface{}) {
	l.read(func() {
		// Pre-allocate the list
		list = make([]interface{}, 0, l.len)

		// Iterate through each item
		l.ForEach(nil, func(_ *Node, val interface{}) bool {
			list = append(list, fn(val))
			return false
		})
	})

	return
}

// Filter will return a filtered list
func (l *LinkedList) Filter(fn FilterFn) (list []interface{}) {
	l.read(func() {
		// Pre-allocate the list
		list = make([]interface{}, 0, l.len)

		// Iterate through each item
		l.ForEach(nil, func(_ *Node, val interface{}) bool {
			if fn(val) {
				list = append(list, val)
			}

			return false
		})
	})

	return
}

// Reduce will return a reduced value
func (l *LinkedList) Reduce(fn ReduceFn) (sum interface{}) {
	l.read(func() {
		// Iterate through each item
		l.ForEach(nil, func(_ *Node, val interface{}) bool {
			sum = fn(sum, val)
			return false
		})
	})

	return
}

// Len will return the current lenght of the linked list
func (l *LinkedList) Len() (n int32) {
	l.read(func() {
		n = l.len
	})

	return
}
