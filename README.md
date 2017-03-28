# LinkedList [![GoDoc](https://godoc.org/github.com/itsmontoya/linkedlist?status.svg)](https://godoc.org/github.com/itsmontoya/linkedlist) ![Status](https://img.shields.io/badge/status-alpha-red.svg)
LinkedList is a simple doubly linked-list implementation which offers:
- Append
- Prepend
- ForEach
- ForEachRev
- Map
- Filter
- Reduce

## Benchmarks
```bash
## Append benchmarks
# Generic list
BenchmarkListAppend-4          10000000         153 ns/op          40 B/op      2 allocs/op
# Typed list (int)
BenchmarkIntListAppend-4       20000000         119 ns/op          32 B/op      1 allocs/op
# Standard library
BenchmarkStdListAppend-4       10000000         246 ns/op          56 B/op      2 allocs/op
BenchmarkSliceAppend-4          3000000         428 ns/op          98 B/op      1 allocs/op
BenchmarkMapAppend-4            5000000         322 ns/op         106 B/op      1 allocs/op

## Prepend benchmarks
# Generic list
BenchmarkListPrepend-4         10000000         158 ns/op          40 B/op      2 allocs/op
# Typed list (int)
BenchmarkIntListPrepend-4      10000000         121 ns/op          32 B/op      1 allocs/op
# Standard library
BenchmarkStdListPrepend-4      10000000         254 ns/op          56 B/op      2 allocs/op
BenchmarkSlicePrepend-4           30000      381157 ns/op      243917 B/op      3 allocs/op
BenchmarkMapPrepend-4           5000000         346 ns/op         106 B/op      1 allocs/op
```

## Usage
```go
package main

import (
	"fmt"

	"github.com/itsmontoya/linkedlist"
)

func main() {
	helloWorld()
	numbers()
}

func helloWorld() {
	l := linkedlist.New(32, linkedlist.ActionGrow)
	l.Append("hello")
	l.Append("world!")

	rv := l.Reduce(func(acc, val interface{}) (sum interface{}) {
		if accV, ok := acc.(string); ok {
			return accV + " " + val.(string)
		}

		return val
	})

	fmt.Println(rv)
}

func numbers() {
	l := linkedlist.New(32, linkedlist.ActionGrow)
	l.Append(0)
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)

	for _, val := range l.Map(doubleNumber) {
		fmt.Println(val)
	}
}

func doubleNumber(val interface{}) (nval interface{}) {
	nval = val.(int) * 2
	return
}

```