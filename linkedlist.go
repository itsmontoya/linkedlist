package linkedlist

var (
	//DEBUG ONLY
	hop = 0
)

// New returns a new linked list
func New(sz int) (ll *LinkedList) {
	ll = &LinkedList{
		head: -1,
		tail: -1,

		s:        make([]item, 0, sz),
		nextGrow: sz * 2,
	}
	return
}

// LinkedList is a linked list
type LinkedList struct {
	head int
	tail int

	s []item
	t tree

	nextGrow int
	putCnt   int
}

func (ll *LinkedList) forEach(idx int, fn func(i item) (end bool)) (ended bool) {
	i := ll.s[idx]
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

func (ll *LinkedList) getIndex(k string) (idx int, ok bool) {
	var bi int
	for _, b := range ll.t {
		for ; bi < len(b); bi++ {
			lf := b[bi]
			hop++
			key := ll.s[lf.val].key

			switch {
			case k > key:
				idx = lf.val
				bi = lf.idx
			case k == key:
				idx = lf.val
				ok = true
				return
			case k < key:
				break
			}
		}
	}

	ll.forEach(idx, func(i item) (end bool) {
		hop++
		idx = i.idx

		switch {
		case k > i.key:
		case k == i.key:
			ok = true
			end = true
		case k < i.key:
			end = true
		}

		return
	})

	return
}

// Get will get
func (ll *LinkedList) Get(k string) (v interface{}) {
	idx, ok := ll.getIndex(k)
	if !ok {
		return
	}

	v = ll.s[idx].val
	hop = 0
	return
}

// Put will put
func (ll *LinkedList) Put(k string, v interface{}) {
	var (
		idx int
		ok  bool
		ni  item
		ri  item
	)

	if ll.head == -1 {
		ll.putFirst(k, v)
		ll.buildTree()
		goto END
	}

	if k > ll.s[ll.tail].key {
		ll.append(newItem(k, v))
		goto END
	}

	if k < ll.s[ll.head].key {
		ll.prepend(newItem(k, v))
		goto END
	}

	if idx, ok = ll.getIndex(k); ok {
		ll.s[idx].val = v
		goto END
	}

	ri = ll.s[idx]
	ni = newItem(k, v)

	ni.idx = len(ll.s)
	ni.next = idx
	ni.prev = ri.prev

	ll.s = append(ll.s, ni)
	ll.s[ri.prev].next = ni.idx
	ll.s[ri.idx].prev = ni.idx

END:
	if ll.putCnt++; ll.putCnt == ll.nextGrow {
		ll.buildTree()
		ll.nextGrow *= 2
	}
	hop = 0
}

// ForEach will iterate through each item until it reaches the end OR the end boolean is returned as true
func (ll *LinkedList) ForEach(fn func(k string, v interface{}) (end bool)) (ended bool) {
	ll.forEach(ll.head, func(i item) (end bool) {
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

func (ll *LinkedList) buildTree() {
	ll.t = newTree(ll)
}
