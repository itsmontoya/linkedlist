package linkedlist

import ()

const (
	// ActionReject will reject the new entry if the list is full
	ActionReject Action = iota
	// ActionGrow will grow the list and accept the entry if the list if full
	ActionGrow
	// ActionPop will pop the opposing end and accept the entry if the list if full
	// Note: Append will pop the first entry, Prepend will pop the last entry
	ActionPop
)

// Action represents which action to take when a full list is encountered on insert
type Action uint8

// ForEachFn is the format of the function used to call ForEach
type ForEachFn func(n *Node, val interface{}) (end bool)

// MapFn is the format of the function used to call Map
type MapFn func(val interface{}) (nval interface{})

// FilterFn is the format of the function used to call Filter
type FilterFn func(val interface{}) (ok bool)

// ReduceFn is the format of the function used to call Reduce
type ReduceFn func(acc, val interface{}) (sum interface{})
