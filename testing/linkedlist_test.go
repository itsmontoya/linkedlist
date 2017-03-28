package main

import (
	"fmt"
	"testing"

	. "github.com/itsmontoya/linkedlist"
	intlist "github.com/itsmontoya/linkedlist/typed/int"
	"github.com/joeshaw/gengen/generic"
)

func TestLinkedList(t *testing.T) {
	var (
		l   LinkedList
		err error
	)

	l.Append(0, 1, 2, 3, 4, 5, 6)

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

func TestMapFilterReduce(t *testing.T) {
	var l LinkedList
	l.Append(0, 1, 2, 3, 4, 5, 6)

	val := l.Map(testAddOne).Filter(testIsEven).Reduce(testAddInts)
	if val != 12 {
		t.Fatalf("expected %v and received %v", 12, val)
	}
}

func testIteration(l *LinkedList, start int) (err error) {
	cnt := start

	l.ForEach(nil, func(_ *Node, val generic.T) bool {
		if val.(int) != cnt {
			err = fmt.Errorf("invalid value, expected %d and received %d", cnt, val)
			return true
		}

		cnt++
		return false
	})

	cnt--

	l.ForEachRev(nil, func(_ *Node, val generic.T) bool {
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
	list := l.Map(func(val generic.T) (nval generic.T) {
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
	list := l.Filter(func(val generic.T) (ok bool) {
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
	val := l.Reduce(func(acc, val generic.T) (sum generic.T) {
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

func testAddOne(val generic.T) (nval generic.T) {
	nval = val.(int) + 1
	return
}

func testIsEven(val generic.T) (ok bool) {
	return val.(int)%2 == 0
}

func testAddInts(acc, val generic.T) (sum generic.T) {
	accV, _ := acc.(int)
	sum = accV + val.(int)
	return
}

func BenchmarkListAppend(b *testing.B) {
	var l LinkedList
	for i := 0; i < b.N; i++ {
		l.Append(i)
	}

	b.ReportAllocs()
}

func BenchmarkIntListAppend(b *testing.B) {
	var l intlist.LinkedList
	for i := 0; i < b.N; i++ {
		l.Append(i)
	}

	b.ReportAllocs()
}

func BenchmarkSliceAppend(b *testing.B) {
	s := make([]generic.T, 0, 32)
	for i := 0; i < b.N; i++ {
		s = append(s, i)
	}

	b.ReportAllocs()
}

func BenchmarkMapAppend(b *testing.B) {
	s := make(map[int]generic.T, 32)
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

func BenchmarkIntListPrepend(b *testing.B) {
	var l intlist.LinkedList
	for i := 0; i < b.N; i++ {
		l.Prepend(i)
	}

	b.ReportAllocs()
}

func BenchmarkSlicePrepend(b *testing.B) {
	s := make([]generic.T, 0, 32)
	for i := 0; i < b.N; i++ {
		s = append([]generic.T{i}, s...)
	}

	b.ReportAllocs()
}

func BenchmarkMapPrepend(b *testing.B) {
	s := make(map[int]generic.T, 32)
	for i := 0; i < b.N; i++ {
		s[i] = i
	}

	b.ReportAllocs()
}
