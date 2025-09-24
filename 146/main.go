package main

import "fmt"

// 146. LRU Cache
// Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.
// Implement the LRUCache class:
//    LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
//    int get(int key) Return the value of the key if the key exists, otherwise return -1.
//    void put(int key, int value) Update the value of the key if the key exists.
//  	Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation,
//  	evict the least recently used key.
// The functions get and put must each run in O(1) average time complexity.

type LRUCache struct {
	cap    int
	keyMap map[int]*Cache
	first  *Cache
	last   *Cache
}

type Cache struct {
	value int
	key   int
	prev  *Cache
	next  *Cache
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:    capacity,
		keyMap: make(map[int]*Cache, capacity),
	}
}

func (l *LRUCache) Get(key int) int {
	if cached, ok := l.keyMap[key]; ok {
		l.accessItem(cached)
		return cached.value
	} else {
		return -1
	}
}

func (l *LRUCache) isFull(key int) bool {
	return len(l.keyMap) >= l.cap
}

func (l *LRUCache) Put(key int, value int) {
	if cached, ok := l.keyMap[key]; ok {
		cached.value = value
		l.accessItem(cached)
	} else {
		if l.isFull(key) {
			// Remove last item
			delete(l.keyMap, l.last.key)
			if l.last.prev != nil {
				l.last.prev.next = nil
			}
			l.last = l.last.prev

		}
		// Add new one and bubble up to the first place
		cached = &Cache{
			key:   key,
			value: value,
		}

		l.keyMap[key] = cached
		l.accessItem(cached)

	}
}

func (l *LRUCache) accessItem(cached *Cache) {
	if l.first == cached {
		return
	}
	if cached.prev != nil {
		cached.prev.next = cached.next
	}
	if cached.next != nil {
		cached.next.prev = cached.prev
	}
	if l.first == nil {
		l.first = cached
		l.last = cached
		return
	}
	l.first.prev = cached
	cached.next = l.first
	if l.last == cached && cached.prev != nil {
		l.last = cached.prev
	}
	cached.prev = nil
	l.first = cached
	if l.last == nil {
		l.last = cached
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	cmds := []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"}
	args := [][]int{[]int{2}, []int{1, 1}, []int{2, 2}, []int{1}, []int{3, 3}, []int{2}, []int{4, 4}, []int{1}, []int{3}, []int{4}}
	res := []int{0, 0, 0, 1, 0, -1, 0, -1, 3, 4}
	test(cmds, args, res)
}

// []
// [[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
// [nil,nil,nil,1,nil,-1,nil,-1,3,4]
func test(commands []string, args [][]int, results []int) []int {

	if len(commands) != len(args) || len(commands) != len(results) {
		panic("Wrong number of arguments")
	}
	var lru LRUCache
	var res int
	ret := make([]int, 0)
	for i := 0; i < len(commands); i++ {
		res = 0
		if commands[i] == "LRUCache" {
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
