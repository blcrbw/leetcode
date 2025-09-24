package main

// 460. LFU Cache
// Design and implement a data structure for a Least Frequently Used (LFU) cache.
//
// Implement the LFUCache class:
//    LFUCache(int capacity) Initializes the object with the capacity of the data structure.
//    int get(int key) Gets the value of the key if the key exists in the cache. Otherwise, returns -1.
//    void put(int key, int value) Update the value of the key if present, or inserts the key if not already present.
//   	When the cache reaches its capacity, it should invalidate and remove the least frequently used key before
//  	inserting a new item. For this problem, when there is a tie (i.e., two or more keys with the same frequency),
// 		the least recently used key would be invalidated.
// To determine the least frequently used key, a use counter is maintained for each key in the cache. The key with the
//	 smallest use counter is the least frequently used key.
// When a key is first inserted into the cache, its use counter is set to 1 (due to the put operation). The use counter
//	 for a key in the cache is incremented either a get or put operation is called on it.
// The functions get and put must each run in O(1) average time complexity.

import (
	"container/heap"
	"fmt"
	"time"
)

type Heap []*Cache

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	if h[i].tie == h[j].tie {
		return h[i].lastAccess < h[j].lastAccess
	}
	return h[i].tie < h[j].tie
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].heapIndex = i
	h[j].heapIndex = j
}

func (h *Heap) Push(x any) {
	item := x.(*Cache)
	item.heapIndex = len(*h)
	*h = append(*h, item)
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

type LFUCache struct {
	cap    int
	keyMap map[int]*Cache
	heap   *Heap
}

type Cache struct {
	value      int
	key        int
	tie        int
	lastAccess int64
	heapIndex  int
}

func Constructor(capacity int) LFUCache {
	var hp Heap
	heap.Init(&hp)

	return LFUCache{
		cap:    capacity,
		keyMap: make(map[int]*Cache, capacity),
		heap:   &hp,
	}
}

func (l *LFUCache) Get(key int) int {
	if cached, ok := l.keyMap[key]; ok {
		l.accessItem(cached)
		return cached.value
	} else {
		return -1
	}
}

func (l *LFUCache) isFull(key int) bool {
	return len(l.keyMap) >= l.cap
}

func (l *LFUCache) Put(key int, value int) {
	if cached, ok := l.keyMap[key]; ok {
		cached.value = value
		l.accessItem(cached)
	} else {
		if l.isFull(key) {
			l.deleteLast()
		}
		cached = &Cache{
			key:        key,
			value:      value,
			tie:        1,
			lastAccess: time.Now().UnixNano(),
		}

		l.keyMap[key] = cached
		heap.Push(l.heap, cached)
	}
}

func (l *LFUCache) deleteLast() {
	// Remove last item
	last := heap.Pop(l.heap).(*Cache)
	delete(l.keyMap, last.key)
}

func (l *LFUCache) accessItem(cached *Cache) {
	cached.tie++
	cached.lastAccess = time.Now().UnixNano()
	heap.Fix(l.heap, cached.heapIndex)
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

// []
// [[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
// [nil,nil,nil,1,nil,-1,nil,-1,3,4]
func test(commands []string, args [][]int, results []int) []int {

	if len(commands) != len(args) || len(commands) != len(results) {
		panic("Wrong number of arguments")
	}
	var lru LFUCache
	var res int
	ret := make([]int, 0)
	for i := 0; i < len(commands); i++ {
		res = 0
		if commands[i] == "LFUCache" {
			lru = Constructor(args[i][0])
		} else if commands[i] == "put" {
			lru.Put(args[i][0], args[i][1])
		} else if commands[i] == "get" {
			res = lru.Get(results[i])
			fmt.Printf("GET result: expected %d, got %d\n", results[i], res)
		}
		ret = append(ret, res)
	}
	return ret
}
