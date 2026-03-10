package main

type DNode struct {
	Key, Val   int
	Prev, Next *DNode
}

type LRUCache struct {
	mp    map[int]*DNode
	cap   int
	dummy *DNode // 双向循环链表
}

func Constructor(capacity int) LRUCache {
	dummy := &DNode{}
	dummy.Next, dummy.Prev = dummy, dummy
	return LRUCache{
		mp:    make(map[int]*DNode),
		cap:   capacity,
		dummy: dummy,
	}
}

func (lru *LRUCache) remove(node *DNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (lru *LRUCache) pushFront(node *DNode) {
	node.Prev = lru.dummy
	node.Next = lru.dummy.Next
	node.Prev.Next = node
	node.Next.Prev = node
}

func (lru *LRUCache) getNode(key int) *DNode {
	node, ok := lru.mp[key]
	if !ok {
		return nil
	}
	lru.remove(node)
	lru.pushFront(node)
	return node
}

func (lru *LRUCache) Get(key int) int {
	node := lru.getNode(key)
	if node == nil {
		return -1
	}
	return node.Val
}

func (lru *LRUCache) Put(key int, value int) {
	node := lru.getNode(key)
	if node != nil {
		node.Val = value
		return
	}
	if len(lru.mp) == lru.cap {
		backNode := lru.dummy.Prev
		delete(lru.mp, backNode.Key)
		lru.remove(backNode)
	}
	newp := &DNode{
		Key: key,
		Val: value,
	}
	lru.pushFront(newp)
	lru.mp[key] = newp
}

/*

type LRUCache struct {
	q   []int       // 先进先出的逻辑
	cap int         // 最大容量
	len int         // 当前容量
	m   map[int]int // 记录key被刷新的次数
	v   map[int]int // 记录值
}

func Constructor_146(cap int) LRUCache {
	return LRUCache{
		make([]int, 0, cap),
		cap,
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
		for lru.len == lru.cap {
			last := lru.q[0]
			lru.q = lru.q[1:]
			lru.m[last]--
			if lru.m[last] == 0 {
				lru.len--
				delete(lru.v, last)
			}
		}
		lru.len++
	}

	lru.v[key] = value
	lru.q = append(lru.q, key)
	lru.m[key]++
}

*/
