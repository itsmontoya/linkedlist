package linkedlist

// ForEachFn is the format of the function used to call ForEach
type ForEachFn func(n *Node, val interface{}) (end bool)

// MapFn is the format of the function used to call Map
type MapFn func(val interface{}) (nval interface{})

// FilterFn is the format of the function used to call Filter
type FilterFn func(val interface{}) (ok bool)

// ReduceFn is the format of the function used to call Reduce
type ReduceFn func(acc, val interface{}) (sum interface{})
