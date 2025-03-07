package main

import "math/rand"

type RandomizedSet struct {
	arr []int
	set map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		arr: []int{},
		set: map[int]int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.set[val]; ok {
		return false
	}
	this.set[val] = len(this.arr)
	this.arr = append(this.arr, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.set[val]
	if !ok {
		return false
	}
	lastIdx := len(this.arr) - 1
	// 这句是核心。。
	this.arr[idx] = this.arr[lastIdx]
	this.set[this.arr[idx]] = idx

	this.arr = this.arr[:lastIdx]
	delete(this.set, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	idx := rand.Intn(len(this.arr))
	return this.arr[idx]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
