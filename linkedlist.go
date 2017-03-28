package linkedlist

import "sync"

// LinkedList is a simple doubly-linked list
type LinkedList struct {
	mux sync.RWMutex

	head *Node
	tail *Node

	reporter bool
	len      int32
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

// prepend will prepend the list with a value, the reference Node is Returned
func (l *LinkedList) prepend(val interface{}) (n *Node) {
	n = newNode(nil, l.head, val)

	if l.head != nil {
		l.head.prev = n
	}

	if l.tail == nil {
		l.tail = n
	}

	l.head = n
	l.len++
	return
}

// append will append the list with a value, the reference Node is Returned
func (l *LinkedList) append(val interface{}) (n *Node) {
	n = newNode(l.tail, nil, val)

	if l.tail != nil {
		l.tail.next = n
	}

	if l.head == nil {
		l.head = n
	}

	l.tail = n
	l.len++
	return
}

// remove will remove a node from a list
func (l *LinkedList) remove(n *Node) {
	if n.prev != nil {
		n.prev.next = n.next
	}

	if n.next != nil {
		n.next.prev = n.prev
	}

	n.prev = nil
	n.next = nil
	n.val = nil
}

// forEach will iterate through each Node within the linked list
func (l *LinkedList) forEach(n *Node, fn ForEachFn) (ended bool) {
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

	return
}

// forEachRev will iterate through each Node within the linked list in reverse
func (l *LinkedList) forEachRev(n *Node, fn ForEachFn) (ended bool) {
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

	return
}

// Prepend will prepend the list with a value, the reference Node is Returned
func (l *LinkedList) Prepend(val interface{}) (n *Node) {
	l.write(func() {
		n = l.prepend(val)
	})

	return
}

// Append will append the list with a value, the reference Node is Returned
func (l *LinkedList) Append(val interface{}) (n *Node) {
	l.write(func() {
		n = l.append(val)
	})

	return
}

// Remove will remove a node from a list
func (l *LinkedList) Remove(n *Node) {
	l.write(func() {
		l.remove(n)
	})
}

// ForEach will iterate through each Node within the linked list
func (l *LinkedList) ForEach(n *Node, fn ForEachFn) (ended bool) {
	l.read(func() {
		ended = l.forEach(n, fn)
	})

	return
}

// ForEachRev will iterate through each Node within the linked list in reverse
func (l *LinkedList) ForEachRev(n *Node, fn ForEachFn) (ended bool) {
	l.read(func() {
		ended = l.forEachRev(n, fn)
	})

	return
}

// Map will return a mapped list
func (l *LinkedList) Map(fn MapFn) (nl *LinkedList) {
	if !l.reporter {
		nl = &LinkedList{reporter: true}
	} else {
		nl = l
	}

	l.read(func() {
		// Iterate through each item
		l.forEach(nil, func(n *Node, val interface{}) bool {
			if !l.reporter {
				nl.append(fn(val))
			} else {
				n.val = fn(val)
			}

			return false
		})
	})

	return
}

// Filter will return a filtered list
func (l *LinkedList) Filter(fn FilterFn) (nl *LinkedList) {
	if !l.reporter {
		nl = &LinkedList{reporter: true}
	} else {
		nl = l
	}

	l.read(func() {
		// Iterate through each item
		l.forEach(nil, func(n *Node, val interface{}) bool {
			if !l.reporter {
				if fn(val) {
					nl.append(val)
				}
			} else {
				if !fn(val) {
					nl.remove(n)
				}
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
		l.forEach(nil, func(_ *Node, val interface{}) bool {
			sum = fn(sum, val)
			return false
		})
	})

	return
}

// Slice will return a slice of the current linked list
func (l *LinkedList) Slice() (s []interface{}) {
	l.read(func() {
		s = make([]interface{}, 0, l.len)
		l.forEach(nil, func(_ *Node, val interface{}) bool {
			s = append(s, val)
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
