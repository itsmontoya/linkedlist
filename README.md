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
// Need to write this still
```