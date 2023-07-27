package main

import "fmt"

func newHashMap() *HashMap {
	return &HashMap{
		nBuckets: 4,
		buckets:  make([]*Entry, 4),
	}
}

func (h HashMap) hash(s string) int {
	x := 5381
	for _, ch := range s {
		x = 33*x + int(ch)
	}
	return x
}

func (h HashMap) getBucket(key string) int {
	return h.hash(key) % h.nBuckets
}

func (h HashMap) set(key string, val interface{}) {
	bucket := h.getBucket(key)
	fmt.Print(bucket)
	v := h.buckets[bucket]

	for v != nil {
		if v.key == key {
			v.val = val
			return
		}
		v = v.next
	}

	newVal := &Entry{
		key:  key,
		val:  val,
		next: h.buckets[bucket],
	}
	h.buckets[bucket] = newVal
}

func (h HashMap) get(key string) interface{} {
	bucket := h.getBucket(key)
	v := h.buckets[bucket]

	for v != nil {
		if v.key == key {
			return v.val
		}
		v = v.next
	}

	return nil
}

func (h HashMap) delete(key string) {
	bucket := h.getBucket(key)
	var prev *Entry = nil
	var v *Entry = h.buckets[bucket]
	for v != nil {
		if v.key == key {
			if prev == nil {
				h.buckets[bucket] = v.next
			} else {
				prev.next = v.next
			}
		}
		prev = v
		v = v.next
	}
}

func main() {
	h := newHashMap()
	a := 5
	h.set("item a", a)
	h.set("item b", a)
	fmt.Printf("%d\n", h.get("item a"))
	fmt.Printf("%d\n", h.get("item b"))
	h.delete("item a")
	h.delete("item b")
}
