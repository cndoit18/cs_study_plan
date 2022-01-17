package lru

import "container/list"

type LRUCache struct {
	cache   map[int]int
	indexer map[int]*list.Element
	ll      *list.List
	max     int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache:   map[int]int{},
		ll:      list.New(),
		max:     capacity,
		indexer: map[int]*list.Element{},
	}
}

func (this *LRUCache) Get(key int) int {
	if value, ok := this.cache[key]; ok {
		this.ll.MoveToFront(this.indexer[key])
		return value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; ok {
		this.ll.MoveToFront(this.indexer[key])
	} else {
		this.indexer[key] = this.ll.PushFront(key)
	}
	this.cache[key] = value
	if this.ll.Len() > this.max {
		key := this.ll.Remove(this.ll.Back()).(int)
		delete(this.cache, key)
		delete(this.indexer, key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
