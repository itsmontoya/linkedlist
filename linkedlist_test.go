package linkedlist

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

var (
	simpleTest = getSorted(1024)

	setKey   string
	setValue interface{}
)

type kv struct {
	Key string
	Val interface{}
}

func TestMain(m *testing.M) {
	sc := m.Run()
	os.Exit(sc)
}

func TestBasic(t *testing.T) {
	var (
		val interface{}
		cnt int64 = 1024 * 1024
		ll        = New(int(cnt))
	)

	sl := getSorted(cnt)

	head := sl[0]
	mid := sl[len(sl)/2]
	tail := sl[len(sl)-1]

	for i, kv := range sl {
		ll.Put(kv.Key, kv.Val)
	}

	if val = ll.Get(head.Key); val != head.Val {
		t.Error("Invalid value for a", val)
	}

	if val = ll.Get(mid.Key); val != mid.Val {
		t.Error("Invalid value for b", val)
	}

	if val = ll.Get(tail.Key); val != tail.Val {
		t.Error("Invalid value for c", val)
	}

	ll.buildTree()
	if val = ll.Get(head.Key); val != head.Val {
		t.Error("Invalid value for a", val)
	}

	if val = ll.Get(mid.Key); val != mid.Val {
		t.Error("Invalid value for b", val)
	}

	if val = ll.Get(tail.Key); val != tail.Val {
		t.Error("Invalid value for c", val)
	}
}

func TestMedium(t *testing.T) {
	//	ll := getPopulatedLL()
	//	fmt.Println("Oh yes?", ll.Get(strconv.FormatInt(time.Now().Unix(), 10)))
}

func BenchmarkPut(b *testing.B) {
	b.StopTimer()
	ll := New(32)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		for _, kv := range simpleTest {
			ll.Put(kv.Key, kv.Val)
		}
	}

	b.ReportAllocs()
}

func BenchmarkMapPut(b *testing.B) {
	b.StopTimer()
	m := make(map[string]interface{})
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		for _, kv := range simpleTest {
			m[kv.Key] = kv.Val
		}
	}

	b.ReportAllocs()
}

func BenchmarkGet(b *testing.B) {
	b.StopTimer()
	ll := getPopulatedLL()
	b.StartTimer()

	var val interface{}
	for i := 0; i < b.N; i++ {
		for _, kv := range simpleTest {
			val = ll.Get(kv.Key)
		}
	}

	setValue = val
	b.ReportAllocs()
}

func BenchmarkMapGet(b *testing.B) {
	b.StopTimer()
	m := getPopulatedMap()
	b.StartTimer()

	var val interface{}
	for i := 0; i < b.N; i++ {
		for _, kv := range simpleTest {
			val = m[kv.Key]
		}
	}

	setValue = val
	b.ReportAllocs()
}

func BenchmarkForEach(b *testing.B) {
	b.StopTimer()
	ll := getPopulatedLL()
	b.StartTimer()

	var (
		key string
		val interface{}
	)

	for i := 0; i < b.N; i++ {
		ll.ForEach(func(k string, v interface{}) (end bool) {
			key = k
			val = v
			return
		})
	}

	setKey = key
	setValue = val
	b.ReportAllocs()
}

func BenchmarkMapForEach(b *testing.B) {
	b.StopTimer()
	m := getPopulatedMap()
	b.StartTimer()

	var (
		key string
		val interface{}
	)

	for i := 0; i < b.N; i++ {
		for k, v := range m {
			key = k
			val = v
		}
	}

	setKey = key
	setValue = val
	b.ReportAllocs()
}

func getPopulatedLL() *LinkedList {
	ll := New(32)

	for _, kv := range simpleTest {
		ll.Put(kv.Key, kv.Val)
	}

	ll.buildTree()
	return ll
}

func getPopulatedMap() map[string]interface{} {
	m := make(map[string]interface{})

	for _, kv := range simpleTest {
		m[kv.Key] = kv.Val
	}

	return m
}

func getSorted(n int64) (s []kv) {
	var i int64
	now := time.Now().Unix()
	for i = 0; i < n; i++ {
		v := now + i
		s = append(s, kv{strconv.FormatInt(v, 10), v})
	}

	return
}

func timedFn(name string, fn func()) (duration int64) {
	s := time.Now().UnixNano()
	fn()
	e := time.Now().UnixNano()
	duration = e - s
	fmt.Printf("%s took %v nanoseconds!\n", name, duration)
	return
}

func getRand(n int64) (s []kv) {
	s = make([]kv, int(n))
	src := getSorted(n)
	perm := rand.Perm(int(n))

	for i, v := range perm {
		s[v] = src[i]
	}

	return
}

func comparisons() {
	ll := New(32)
	m := make(map[string]interface{})
	example := getSorted(1024 * 1024)
	randExample := getRand(1024 * 12)

	fmt.Println("Putting (sorted):")
	timedFn("ll", func() {
		for _, kv := range example {
			ll.Put(kv.Key, kv.Val)
		}
	})

	timedFn("map", func() {
		for _, kv := range example {
			m[kv.Key] = kv.Val
		}
	})

	ll = New(32)
	m = make(map[string]interface{})
	fmt.Print("\nPutting (random):\n")

	timedFn("ll", func() {
		for _, kv := range randExample {
			ll.Put(kv.Key, kv.Val)
		}
	})

	timedFn("map", func() {
		for _, kv := range randExample {
			m[kv.Key] = kv.Val
		}
	})

	fmt.Print("\nIteration:\n")
	timedFn("ll", func() {
		ll.ForEach(func(k string, v interface{}) (end bool) {
			setKey = k
			setValue = v
			return
		})
	})

	timedFn("map", func() {
		for k, v := range m {
			setKey = k
			setValue = v
		}
	})
}
