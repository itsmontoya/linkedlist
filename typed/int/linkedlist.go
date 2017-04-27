// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package linkedlist

var (
	zeroVal int
	zeroSum int
)

// LinkedList is a simple doubly-linked list
type LinkedList struct {
	head *Node
	tail *Node

	reporter bool
	len      int32
}

// prepend will prepend the list with a value, the reference node is Returned
func (l *LinkedList) prepend(val int) (n *Node) {
	n = newNode(nil, l.head, val)

	if l.head != nil {
		// Head exists, set the previous value to our new node
		l.head.prev = n
	}

	if l.tail == nil {
		// This is the first item, so it will be the head AND the tail
		l.tail = n
	}

	// Set head as our new node
	l.head = n
	// Increment node count
	l.len++
	return
}

// append will append the list with a value, the reference node is Returned
func (l *LinkedList) append(val int) (n *Node) {
	n = newNode(l.tail, nil, val)

	if l.tail != nil {
		// Tail exists, set the next value to our new node
		l.tail.next = n
	}

	if l.head == nil {
		// This is the first item, so it will be the head AND the tail
		l.head = n
	}

	// Set tail as our new node
	l.tail = n
	// Increment node count
	l.len++
	return
}

// mapCopy will return a copied and mapped list
func (l *LinkedList) mapCopy(fn MapFn) (nl *LinkedList) {
	nl = &LinkedList{reporter: true}
	// Iterate through each item
	l.ForEach(nil, func(n *Node, val int) bool {
		nl.append(fn(val))
		return false
	})

	return
}

// mapModify will return a copied and mapped list
func (l *LinkedList) mapModify(fn MapFn) (nl *LinkedList) {
	nl = l
	// Iterate through each item
	l.ForEach(nil, func(n *Node, val int) bool {
		n.val = fn(val)
		return false
	})

	return
}

// filterCopy will return a copied and filtered list
func (l *LinkedList) filterCopy(fn FilterFn) (nl *LinkedList) {
	nl = &LinkedList{reporter: true}
	// Iterate through each item
	l.ForEach(nil, func(_ *Node, val int) bool {
		if fn(val) {
			nl.append(val)
		}

		return false
	})

	return
}

// filterModify will modify and return filtered list
func (l *LinkedList) filterModify(fn FilterFn) (nl *LinkedList) {
	nl = l
	// Iterate through each item
	l.ForEach(nil, func(n *Node, val int) bool {
		if !fn(val) {
			l.Remove(n)
		}

		return false
	})

	return
}

// Prepend will prepend the list with a value, the reference Node is Returned
func (l *LinkedList) Prepend(vals ...int) {
	// Iterate through provided values
	for _, val := range vals {
		l.prepend(val)
	}

	return
}

// Append will append the list with a value, the reference Node is Returned
func (l *LinkedList) Append(vals ...int) {
	// Iterate through provided values
	for _, val := range vals {
		l.append(val)
	}

	return
}

// Remove will remove a node from a list
func (l *LinkedList) Remove(n *Node) {
	if n.prev != nil {
		// Set previous node's next as our current next node
		n.prev.next = n.next
	} else {
		// We have no previous, which means this is the head node
		// Set head as the node which proceeds this one
		if l.head = n.next; l.head != nil {
			// Remove the previous value from our new head
			l.head.prev = nil
		}
	}

	if n.next != nil {
		// Set next node's previous as our current previous node
		n.next.prev = n.prev
	} else {
		// We have no next, which means this is the tail node
		// Set tail as the node which precedes this one
		if l.tail = n.prev; l.tail != nil {
			// Remove the next value from our new tail
			l.tail.next = nil
		}
	}

	// Set node to zero values
	n.prev = nil
	n.next = nil
	n.val = zeroVal
	// Decrement node count
	l.len--
}

// ForEach will iterate through each node within the linked list
func (l *LinkedList) ForEach(n *Node, fn ForEachFn) (ended bool) {
	if n == nil {
		// Provided node is nil, set to head
		n = l.head
	}

	// Next node
	var nn *Node
	// Iterate until n equals nil
	for n != nil {
		// Set next node
		nn = n.next
		// Call provided func
		if fn(n, n.val) {
			// Func returned true, return with ended as true
			return true
		}

		// Set n as the next node
		n = nn
	}

	return false
}

// ForEachRev will iterate through each node within the linked list in reverse
func (l *LinkedList) ForEachRev(n *Node, fn ForEachFn) (ended bool) {
	if n == nil {
		// Provided node is nil, set to tail
		n = l.tail
	}

	// Previous node
	var pn *Node
	// Iterate until n equals nil
	for n != nil {
		// Set previous node
		pn = n.prev
		// Call provided func
		if fn(n, n.val) {
			// Func returned true, return with ended as true
			return true
		}

		// Set n as the previous node
		n = pn
	}

	return false
}

// Map will return a mapped list
func (l *LinkedList) Map(fn MapFn) (nl *LinkedList) {
	if l.reporter {
		return l.mapModify(fn)
	}

	return l.mapCopy(fn)
}

// Filter will return a filtered list
func (l *LinkedList) Filter(fn FilterFn) (nl *LinkedList) {
	if l.reporter {
		return l.filterModify(fn)
	}

	return l.filterCopy(fn)
}

// Reduce will return a reduced value
func (l *LinkedList) Reduce(fn ReduceFn) (sum int) {
	// Iterate through each item
	l.ForEach(nil, func(_ *Node, val int) bool {
		sum = fn(sum, val)
		return false
	})

	return
}

// Slice will return a slice of the current linked list
func (l *LinkedList) Slice() (s []int) {
	s = make([]int, 0, l.len)
	l.ForEach(nil, func(_ *Node, val int) bool {
		s = append(s, val)
		return false
	})

	return
}

// Val will return the value for a given node
func (l *LinkedList) Val(n *Node) (val int) {
	return n.val
}

// Update will update the value for a given node
func (l *LinkedList) Update(n *Node, val int) {
	n.val = val
}

// Len will return the current length of the linked list
func (l *LinkedList) Len() (n int32) {
	return l.len
}

func newNode(prev, next *Node, val int) *Node {
	return &Node{prev, next, val}
}

// Node is a value container
type Node struct {
	prev *Node
	next *Node

	val int
}

// ForEachFn is the format of the function used to call ForEach
type ForEachFn func(n *Node, val int) (end bool)

// MapFn is the format of the function used to call Map
type MapFn func(val int) (nval int)

// FilterFn is the format of the function used to call Filter
type FilterFn func(val int) (ok bool)

// ReduceFn is the format of the function used to call Reduce
type ReduceFn func(acc, val int) (sum int)
