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
# Append benchmark
BenchmarkListAppend-4         10000000         258 ns/op          56 B/op      2 allocs/op
BenchmarkSliceAppend-4         5000000         393 ns/op          92 B/op      1 allocs/op
BenchmarkMapAppend-4           5000000         329 ns/op         106 B/op      1 allocs/op

# Prepend benchmark
BenchmarkListPrepend-4        10000000         253 ns/op          56 B/op      2 allocs/op
BenchmarkSlicePrepend-4          30000      375657 ns/op      243917 B/op      3 allocs/op
BenchmarkMapPrepend-4          5000000         341 ns/op         106 B/op      1 allocs/op
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