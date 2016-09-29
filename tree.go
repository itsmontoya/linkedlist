package linkedlist

import "math"

func newTree(ll *LinkedList) (t tree) {
	t = append(t, newTrunk(ll))

	for li := t[len(t)-1]; len(li) > 3; li = t[len(t)-1] {
		t = append(t, newBranch(li))
	}

	invertTree(t)
	return
}

type tree []branch

type branch []leaf

type leaf struct {
	idx int
	val int
}

func newTrunk(ll *LinkedList) (out branch) {
	var n, cnt int
	l := len(ll.s)
	chunks := int(math.Sqrt(float64(l)))
	chunkSize := l / chunks
	out = make(branch, 0, chunks)

	ll.forEach(0, func(i item) (end bool) {
		if n == 0 {
			out = append(out, leaf{idx: cnt, val: i.idx})
			cnt++
		}

		if n == chunkSize {
			n = 0
			return
		}

		n++
		return
	})

	return
}

func newBranch(in branch) (out branch) {
	var n, cnt int
	l := len(in)
	chunks := int(math.Sqrt(float64(l)))
	chunkSize := l / chunks
	out = make(branch, 0, chunks)

	for _, v := range in {
		if n == 0 {
			v.idx = cnt
			out = append(out, v)
			cnt++
		}

		if n == chunkSize {
			n = 0
			continue
		}

		n++
	}

	return
}

func invertTree(t tree) {
	var (
		j   = len(t) - 1
		cap = len(t) / 2
	)

	for i := 0; i < cap; i++ {
		t[i], t[j] = t[j], t[i]
		j--
	}
}
