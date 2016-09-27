package linkedlist

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	ll := New()
	ll.Put("b", "my")
	ll.Put("e", "are")
	ll.Put("a", "Hello")
	ll.Put("d", "how")
	ll.Put("g", "today?")
	ll.Put("f", "you")
	ll.Put("c", "friend")

	ll.ForEach(func(k string, v interface{}) (end bool) {
		fmt.Println(k, v)
		return
	})
}
