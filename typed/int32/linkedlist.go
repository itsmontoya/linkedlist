package linkedlist

import "sync"

var zeroValue int32

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
func (l *LinkedList) prepend(val int32) (n *Node) {
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
func (l *LinkedList) append(val int32) (n *Node) {
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
	} else {
		l.head = n.next
		l.head.prev = nil
	}

	if n.next != nil {
		n.next.prev = n.prev
	} else {
		l.tail = n.prev
		l.tail.next = nil
	}

	n.prev = nil
	n.next = nil
	n.val = zeroValue
}

// forEach will iterate through each Node within the linked list
func (l *LinkedList) forEach(n *Node, fn ForEachFn) (ended bool) {
	var nn *Node
	if n == nil {
		// Provided Node is nil, set to head
		n = l.head
	}

	for n != nil {
		nn = n.next
		if fn(n, n.val) {
			ended = true
			return
		}

		n = nn
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
func (l *LinkedList) Prepend(vals ...int32) {
	l.write(func() {
		for _, val := range vals {
			l.prepend(val)
		}
	})

	return
}

// Append will append the list with a value, the reference Node is Returned
func (l *LinkedList) Append(vals ...int32) {
	l.write(func() {
		for _, val := range vals {
			l.append(val)
		}
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
		l.forEach(nil, func(n *Node, val int32) bool {
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
		l.forEach(nil, func(n *Node, val int32) bool {
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
func (l *LinkedList) Reduce(fn ReduceFn) (sum int32) {
	l.read(func() {
		// Iterate through each item
		l.forEach(nil, func(_ *Node, val int32) bool {
			sum = fn(sum, val)
			return false
		})
	})

	return
}

// Slice will return a slice of the current linked list
func (l *LinkedList) Slice() (s []int32) {
	l.read(func() {
		s = make([]int32, 0, l.len)
		l.forEach(nil, func(_ *Node, val int32) bool {
			s = append(s, val)
			return false
		})
	})

	return
}

// Val will return the value for a given Node
func (l *LinkedList) Val(n *Node) (val int32) {
	l.read(func() {
		val = n.val
	})

	return
}

// Update will update the value for a given Node
func (l *LinkedList) Update(n *Node, val int32) {
	l.write(func() {
		n.val = val
	})
}

// Len will return the current lenght of the linked list
func (l *LinkedList) Len() (n int32) {
	l.read(func() {
		n = l.len
	})

	return
}

func newNode(prev, next *Node, val int32) *Node {
	return &Node{prev, next, val}
}

// Node is a value container
type Node struct {
	prev *Node
	next *Node

	val int32
}

// ForEachFn is the format of the function used to call ForEach
type ForEachFn func(n *Node, val int32) (end bool)

// MapFn is the format of the function used to call Map
type MapFn func(val int32) (nval int32)

// FilterFn is the format of the function used to call Filter
type FilterFn func(val int32) (ok bool)

// ReduceFn is the format of the function used to call Reduce
type ReduceFn func(acc, val int32) (sum int32)
