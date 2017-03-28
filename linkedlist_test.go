package linkedlist

import (
	//	"fmt"
	"testing"

	double "github.com/itsmontoya/linkedlist/double"
	single "github.com/itsmontoya/linkedlist/single"
)

func TestSingle(t *testing.T) {
	//	var (
	//		l   single.LinkedList
	///		err error
	//)

	//	l.Append(1)

}

func BenchmarkSLLAppend(b *testing.B) {
	var l single.LinkedList
	for i := 0; i < b.N; i++ {
		l.Append(i)
	}

	b.ReportAllocs()
}

func BenchmarkDLLAppend(b *testing.B) {
	var l double.LinkedList
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

func BenchmarkSLLPrepend(b *testing.B) {
	var l single.LinkedList
	for i := 0; i < b.N; i++ {
		l.Prepend(i)
	}

	b.ReportAllocs()
}

func BenchmarkDLLPrepend(b *testing.B) {
	var l double.LinkedList
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
