package linkedlist

import (
//	"fmt"
)

// New returns a new linked list
func New() (ll *LinkedList) {
	ll = &LinkedList{
		head: -1,
		tail: -1,

		s: make([]item, 0, 32),
	}
	return
}

// LinkedList is a linked list
type LinkedList struct {
	head int
	tail int

	s []item
}

func (ll *LinkedList) forEach(fn func(i item) (end bool)) (ended bool) {
	i := ll.s[ll.head]
	for {
		if fn(i) {
			return true
		}

		if i.next == -1 {
			break
		}

		i = ll.s[i.next]
	}

	return
}

func (ll *LinkedList) putFirst(k string, v interface{}) {
	ll.s = append(ll.s, newItem(k, v))
	ll.head = 0
	ll.tail = 0
}

func (ll *LinkedList) prepend(i item) {
	i.next = ll.head
	i.idx = len(ll.s)

	ll.s = append(ll.s, i)
	ll.s[ll.head].prev = i.idx
	ll.head = i.idx
}

func (ll *LinkedList) append(i item) {
	i.prev = ll.tail
	i.idx = len(ll.s)

	ll.s = append(ll.s, i)
	ll.s[ll.tail].next = i.idx
	ll.tail = i.idx
}

// Get will get
func (ll *LinkedList) Get(k string) (v interface{}) {
	ll.ForEach(func(ik string, iv interface{}) (end bool) {
		if ik != k {
			return
		}

		v = iv
		end = true
		return
	})

	return
}

// Put will put
func (ll *LinkedList) Put(k string, v interface{}) {
	if ll.head == -1 {
		ll.putFirst(k, v)
		return
	}

	if k < ll.s[ll.head].key {
		ll.prepend(newItem(k, v))
		return
	}

	if k > ll.s[ll.tail].key {
		ll.append(newItem(k, v))
		return
	}

	ll.forEach(func(i item) (end bool) {
		if k > i.key {
			return
		}

		end = true
		if k == i.key {
			ll.s[i.idx].val = v
			return
		}

		ni := newItem(k, v)
		ni.idx = len(ll.s)
		ni.next = i.idx
		ni.prev = i.prev

		ll.s = append(ll.s, ni)
		ll.s[i.prev].next = ni.idx
		ll.s[i.idx].prev = ni.idx

		return
	})
}

// ForEach will iterate through each item until it reaches the end OR the end boolean is returned as true
func (ll *LinkedList) ForEach(fn func(k string, v interface{}) (end bool)) (ended bool) {
	ll.forEach(func(i item) (end bool) {
		return fn(i.key, i.val)
	})

	return
}

func newItem(k string, v interface{}) item {
	return item{
		key:  k,
		val:  v,
		prev: -1,
		next: -1,
	}
}

type item struct {
	key string
	val interface{}

	prev int
	next int
	idx  int
}
