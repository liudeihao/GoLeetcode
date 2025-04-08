package main

type LRUCache struct {
	q        []int       // 先进先出的逻辑
	capacity int         // 最大容量
	curCap   int         // 当前容量
	m        map[int]int // 记录key被刷新的次数
	v        map[int]int // 记录值
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		make([]int, 0, capacity),
		capacity,
		0,
		make(map[int]int),
		make(map[int]int),
	}
}

func (lru *LRUCache) Get(key int) int {
	v, ok := lru.v[key]
	if !ok {
		return -1
	}
	lru.q = append(lru.q, key)
	lru.m[key]++
	return v
}

func (lru *LRUCache) Put(key int, value int) {
	_, ok := lru.v[key]
	if !ok {
		for lru.curCap == lru.capacity {
			last := lru.q[0]
			lru.q = lru.q[1:]
			lru.m[last]--
			if lru.m[last] == 0 {
				lru.curCap--
				delete(lru.v, last)
			}
		}
		lru.curCap++
	}

	lru.v[key] = value
	lru.q = append(lru.q, key)
	lru.m[key]++
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
