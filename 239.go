package main

// MaxQueue 单调队列
type MaxQueue struct {
	queue []int
}

func (q *MaxQueue) Front() int {
	return q.queue[0]
}
func (q *MaxQueue) Back() int {
	return q.queue[len(q.queue)-1]
}
func (q *MaxQueue) Empty() bool {
	return len(q.queue) == 0
}
func (q *MaxQueue) Push(val int) {
	for !q.Empty() && val > q.Back() { // 这里是val > q.Back()而不是 >=, 不然Pop就坏了
		// pop_back，维护单调队列，这里区别于一般队列
		q.queue = q.queue[:len(q.queue)-1]
	}
	q.queue = append(q.queue, val)
}
func (q *MaxQueue) Pop(val int) {
	// 这里判断是因为它可能已经被pop_back了！
	if !q.Empty() && val == q.Front() {
		q.queue = q.queue[1:]
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	q := MaxQueue{}
	var ans []int
	for i := 0; i < k; i++ {
		q.Push(nums[i])
	}
	ans = append(ans, q.Front())

	for i := k; i < len(nums); i++ {
		q.Pop(nums[i-k])
		q.Push(nums[i])
		ans = append(ans, q.Front())
	}
	return ans
}
