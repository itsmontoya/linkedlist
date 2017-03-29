# LinkedList [![GoDoc](https://godoc.org/github.com/itsmontoya/linkedlist?status.svg)](https://godoc.org/github.com/itsmontoya/linkedlist) ![Status](https://img.shields.io/badge/status-alpha-red.svg)
LinkedList is a simple doubly linked-list implementation which offers:
- Append
- Prepend
- Remove
- ForEach
- ForEachRev
- Map
- Filter
- Reduce

## Benchmarks
```bash
# Generic LinkedList
BenchmarkListAppend-4          10000000         153 ns/op          40 B/op      2 allocs/op
BenchmarkListPrepend-4         10000000         158 ns/op          40 B/op      2 allocs/op

# Typed (int) LinkedList
BenchmarkIntListAppend-4       20000000         119 ns/op          32 B/op      1 allocs/op
BenchmarkIntListPrepend-4      10000000         121 ns/op          32 B/op      1 allocs/op

# Standard library
BenchmarkStdListAppend-4       10000000         246 ns/op          56 B/op      2 allocs/op
BenchmarkStdListPrepend-4      10000000         254 ns/op          56 B/op      2 allocs/op

# Slice
BenchmarkSliceAppend-4          3000000         428 ns/op          98 B/op      1 allocs/op
BenchmarkSlicePrepend-4           30000      381157 ns/op      243917 B/op      3 allocs/op

# Map
BenchmarkMapAppend-4            5000000         322 ns/op         106 B/op      1 allocs/op
BenchmarkMapPrepend-4           5000000         346 ns/op         106 B/op      1 allocs/op
```

## Usage
```go
package main

import (
	"fmt"

	"github.com/itsmontoya/linkedlist/typed/int"
)

func main() {
	var l linkedlist.LinkedList
	// Populate list values
	l.Append(0, 1, 2, 3, 4, 5, 6)

	// Create new list with map applied
	nl := l.Map(addOne)
	// Set mapped value
	mapped := nl.Slice()

	// Filter new list
	nl.Filter(isEven)

	// Set filtered and reduced values
	filtered := nl.Slice()
	reduced := nl.Reduce(addInts)

	// Note - This can also be done shorthand as such:
	// val := l.Map(addOne).Filter(isEven).Reduce(addInts)

	fmt.Printf("Original list: %v\n", l.Slice())
	fmt.Printf("Slice with map applied: %v\n", mapped)
	fmt.Printf("Slice with map and filter applied: %v\n", filtered)
	fmt.Printf("Result of reduction: %v\n", reduced)
}

func addOne(val int) (nval int) {
	return val + 1
}

func isEven(val int) (ok bool) {
	return val%2 == 0
}

func addInts(acc, val int) (sum int) {
	return acc + val
}

```