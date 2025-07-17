package main

func maxSlidingWindow(nums []int, k int) []int {
	deque := []int{0}
	push := func(x int) {
		deque = append(deque, x)
	}
	top := func() int {
		return deque[len(deque)-1]
	}
	pop_back := func() {
		deque = deque[:len(deque)-1]
	}
	empty := func() bool {
		return len(deque) == 0
	}
	for i := 1; i < k; i++ {
		for !empty() && nums[top()] <= nums[i] {
			pop_back()
		}
		push(i)
	}
	ans := []int{nums[deque[0]]}
	for i := k; i < len(nums); i++ {
		for !empty() && (top() < i-k+1 || nums[top()] <= nums[i]) {
			pop_back()
		}
		push(i)
		for deque[0] < i-k+1 {
			deque = deque[1:]
		}
		ans = append(ans, nums[deque[0]])
	}
	return ans
}
