package linkedlist

import (
	"fmt"
	"testing"
)

func TestReject(t *testing.T) {
	var err error
	l := New(4, ActionReject)
	if err = testAppend(l, 0, true); err != nil {
		t.Fatal(err)
	}

	if err = testAppend(l, 1, true); err != nil {
		t.Fatal(err)
	}

	if err = testAppend(l, 2, true); err != nil {
		t.Fatal(err)
	}

	if err = testAppend(l, 3, true); err != nil {
		t.Fatal(err)
	}

	if err = testAppend(l, 4, false); err != nil {
		t.Fatal(err)
	}

	if err = testAppend(l, 5, false); err != nil {
		t.Fatal(err)
	}

	if err = testIteration(l, 0); err != nil {
		t.Fatal(err)
	}

	if err = testMap(l, 0); err != nil {
		t.Fatal(err)
	}

	if err = testFilter(l, 0, true); err != nil {
		t.Fatal(err)
	}

	if err = testReduce(l, 0); err != nil {
		t.Fatal(err)
	}

	return
}

func TestPop(t *testing.T) {
	var err error
	l := New(4, ActionPop)
	if err = testAppend(l, 0, true); err != nil {
		t.Fatal(err)
	}

	if err = testAppend(l, 1, true); err != nil {
		t.Fatal(err)
	}

	if err = testAppend(l, 2, true); err != nil {
		t.Fatal(err)
	}

	if err = testAppend(l, 3, true); err != nil {
		t.Fatal(err)
	}

	if err = testAppend(l, 4, true); err != nil {
		t.Fatal(err)
	}

	if err = testAppend(l, 5, true); err != nil {
		t.Fatal(err)
	}

	if err = testIteration(l, 2); err != nil {
		t.Fatal(err)
	}

	if err = testMap(l, 2); err != nil {
		t.Fatal(err)
	}

	if err = testFilter(l, 0, false); err != nil {
		t.Fatal(err)
	}

	if err = testReduce(l, 2); err != nil {
		t.Fatal(err)
	}

	return
}

func TestGrow(t *testing.T) {
	var err error
	l := New(4, ActionGrow)
	if err = testAppend(l, 0, true); err != nil {
		t.Fatal(err)
	}

	if err = testAppend(l, 1, true); err != nil {
		t.Fatal(err)
	}

	if err = testAppend(l, 2, true); err != nil {
		t.Fatal(err)
	}

	if err = testAppend(l, 3, true); err != nil {
		t.Fatal(err)
	}

	if err = testAppend(l, 4, true); err != nil {
		t.Fatal(err)
	}

	if err = testAppend(l, 5, true); err != nil {
		t.Fatal(err)
	}

	if err = testIteration(l, 0); err != nil {
		t.Fatal(err)
	}

	if err = testMap(l, 0); err != nil {
		t.Fatal(err)
	}

	if err = testFilter(l, 0, true); err != nil {
		t.Fatal(err)
	}

	if err = testReduce(l, 0); err != nil {
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
		return fmt.Errorf("expected nil, received %v", n.Val())
	}

	if nv := n.Val(); nv.(int) != val {
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
	})

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
	})

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
	l := New(32, ActionGrow)
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
	l := New(32, ActionGrow)
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
