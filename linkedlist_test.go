package linkedlist

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	var (
		l   LinkedList
		err error
	)

	if err = testAppend(&l, 0, true); err != nil {
		t.Fatal(err)
	}

	if err = testAppend(&l, 1, true); err != nil {
		t.Fatal(err)
	}

	if err = testAppend(&l, 2, true); err != nil {
		t.Fatal(err)
	}

	if err = testAppend(&l, 3, true); err != nil {
		t.Fatal(err)
	}

	if err = testAppend(&l, 4, true); err != nil {
		t.Fatal(err)
	}

	if err = testAppend(&l, 5, true); err != nil {
		t.Fatal(err)
	}

	if err = testIteration(&l, 0); err != nil {
		t.Fatal(err)
	}

	if err = testMap(&l, 0); err != nil {
		t.Fatal(err)
	}

	if err = testFilter(&l, 0, true); err != nil {
		t.Fatal(err)
	}

	if err = testReduce(&l, 0); err != nil {
		t.Fatal(err)
	}

	return
}

func testAppend(l *LinkedList, val int, expectValue bool) (err error) {
	n := l.Append(val)
	if !expectValue && n == nil {
		return
	}

	if !expectValue && n != nil {
		return fmt.Errorf("expected nil, received %v", n.val)
	}

	if nv := n.val; nv.(int) != val {
		return fmt.Errorf("invalid value, expected %d and received %d", val, nv)
	}

	return
}

func testIteration(l *LinkedList, start int) (err error) {
	cnt := start

	l.ForEach(nil, func(_ *Node, val interface{}) bool {
		if val.(int) != cnt {
			err = fmt.Errorf("invalid value, expected %d and received %d", cnt, val)
			return true
		}

		cnt++
		return false
	})

	cnt--

	l.ForEachRev(nil, func(_ *Node, val interface{}) bool {
		if val.(int) != cnt {
			err = fmt.Errorf("invalid value, expected %d and received %d", cnt, val)
			return true
		}

		cnt--
		return false
	})

	return
}

func testMap(l *LinkedList, start int) (err error) {
	list := l.Map(func(val interface{}) (nval interface{}) {
		nval = val.(int) * 2
		return
	}).Slice()

	for i := 0; i < len(list); i++ {
		v := list[i]
		ev := (i + start) * 2
		if v != ev {
			return fmt.Errorf("invalid value, expected %d and received %d", ev, v)
		}
	}

	return
}

func testFilter(l *LinkedList, tgt int, expected bool) (err error) {
	list := l.Filter(func(val interface{}) (ok bool) {
		return val.(int) == tgt
	}).Slice()

	expectedLen := 1
	if !expected {
		expectedLen = 0
	}

	if ll := len(list); ll != expectedLen {
		err = fmt.Errorf("invalid list length, expected %d and received %d", expectedLen, ll)
	}

	return
}

func testReduce(l *LinkedList, start int) (err error) {
	var cv int
	len := int(l.Len())
	val := l.Reduce(func(acc, val interface{}) (sum interface{}) {
		accV, _ := acc.(int)
		sum = accV + val.(int)
		return
	}).(int)

	for i := start; i < len+start; i++ {
		cv += i
	}

	if val != cv {
		err = fmt.Errorf("invalid value, expected %d and received %d", cv, val)
	}

	return
}

func BenchmarkListAppend(b *testing.B) {
	var l LinkedList
	for i := 0; i < b.N; i++ {
		l.Append(i)
	}

	b.ReportAllocs()
}

func BenchmarkSliceAppend(b *testing.B) {
	s := make([]interface{}, 0, 32)
	for i := 0; i < b.N; i++ {
		s = append(s, i)
	}

	b.ReportAllocs()
}

func BenchmarkMapAppend(b *testing.B) {
	s := make(map[int]interface{}, 32)
	for i := 0; i < b.N; i++ {
		s[i] = i
	}

	b.ReportAllocs()
}

func BenchmarkListPrepend(b *testing.B) {
	var l LinkedList
	for i := 0; i < b.N; i++ {
		l.Prepend(i)
	}

	b.ReportAllocs()
}

func BenchmarkSlicePrepend(b *testing.B) {
	s := make([]interface{}, 0, 32)
	for i := 0; i < b.N; i++ {
		s = append([]interface{}{i}, s...)
	}

	b.ReportAllocs()
}

func BenchmarkMapPrepend(b *testing.B) {
	s := make(map[int]interface{}, 32)
	for i := 0; i < b.N; i++ {
		s[i] = i
	}

	b.ReportAllocs()
}
