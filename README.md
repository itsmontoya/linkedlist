# LinkedList [![GoDoc](https://godoc.org/github.com/itsmontoya/linkedlist?status.svg)](https://godoc.org/github.com/itsmontoya/linkedlist) ![Status](https://img.shields.io/badge/status-alpha-red.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/itsmontoya/linkedlist)](https://goreportcard.com/report/github.com/itsmontoya/linkedlist)
LinkedList is a simple doubly linked-list implementation which offers:
- Append
- Prepend
- Remove
- ForEach
- ForEachRev
- Map
- Filter
- Reduce

## Aren't linked lists bad?
It is true that in many situations, there is a better data structure to use than a linked list. While this is the case for many scenarios, it is not the case for ALL scenarios. Over the years, I've found situations where linked lists have proven extremely useful:

### I need to store data, and I'm not sure if the size beforehand
When you are unsure of your slice size, you have two options:
- Pre-allocate: Which often leads to over-allocating
- Start with empty slice: Causes lots of memcpy and allocations

With a linked list, you can grow your data structure as needed. As soon as the size has stabilized, the linked list can efficiently be converted to a slice.

### I need to prepend a lot
When you need to deal with lots of prepends, the common approach is to append your slice and reverse at the end. There are some caveats with this approach:
- You must wait to access your data until the appending is complete (so reverse can be called)
- You will encounter lots of memcpy during the reverse process, especially if your list is quite large

## Benchmarks
```bash
# go test --bench=.

# Generic LinkedList
BenchmarkListAppend-4          10000000         120 ns/op          40 B/op      2 allocs/op
BenchmarkListPrepend-4         10000000         118 ns/op          40 B/op      2 allocs/op
BenchmarkListFilter-4          50000000        28.3 ns/op          24 B/op      0 allocs/op

# Typed (int) LinkedList
BenchmarkIntListAppend-4       20000000         100 ns/op          32 B/op      1 allocs/op
BenchmarkIntListPrepend-4      20000000        93.2 ns/op          32 B/op      1 allocs/op

# Standard library
BenchmarkStdListAppend-4       10000000         238 ns/op          56 B/op      2 allocs/op
BenchmarkStdListPrepend-4      10000000         238 ns/op          56 B/op      2 allocs/op

# Slice
BenchmarkSliceAppend-4          3000000         430 ns/op          98 B/op      1 allocs/op
BenchmarkSlicePrepend-4           30000      369399 ns/op      243917 B/op      2 allocs/op
BenchmarkSliceFilter-4         20000000        74.2 ns/op          41 B/op      0 allocs/op

# Map
BenchmarkMapAppend-4            5000000         346 ns/op         106 B/op      1 allocs/op
BenchmarkMapPrepend-4           5000000         358 ns/op         106 B/op      1 allocs/op
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
