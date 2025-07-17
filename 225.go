package main

type Queue []int

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}
func (q *Queue) Pop() int {
	top := (*q)[0]
	*q = (*q)[1:]
	return top
}
func (q *Queue) Peek() int {
	return (*q)[0]
}
func (q *Queue) Empty() bool {
	return q.Size() == 0
}
func (q *Queue) Size() int {
	return len(*q)
}

type MyStack struct {
	// 一个队列就够了
	q Queue
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.q.Push(x)
}

func (this *MyStack) Pop() int {
	size := this.q.Size()
	for size > 1 {
		size--
		this.q.Push(this.q.Pop())
	}
	return this.q.Pop()
}

func (this *MyStack) Top() int {
	size := this.q.Size()
	for size > 1 {
		size--
		this.q.Push(this.q.Pop())
	}
	top := this.q.Peek()
	this.q.Push(this.q.Pop())
	return top
}

func (this *MyStack) Empty() bool {
	return this.q.Empty()
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
