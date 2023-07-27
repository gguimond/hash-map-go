package main

type Entry struct {
	key  string
	val  interface{}
	next *Entry
}

type HashMap struct {
	buckets  []*Entry
	nBuckets int
}

type HashMapInterface interface {
	hash(string) int
	getBucket(string) int
	set(string, interface{})
	get(string) interface{}
	delete(string)
}
