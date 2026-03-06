package main

func maxSlidingWindow(nums []int, k int) []int {
    maxQ := []int{}
    empty := func() bool {
        return len(maxQ) == 0
    }
    front := func() int {
        return maxQ[0]
    }
    back := func() int {
        return maxQ[len(maxQ)-1]
    }
    pop := func() {
        maxQ = maxQ[1:]
    }
    pop_back := func() {
        maxQ = maxQ[:len(maxQ)-1]
    }
    push := func(i int) {
        for !empty() && nums[back()] <= nums[i] {
            pop_back()
        }
        maxQ = append(maxQ, i)
    }
    for i := range k {
        push(i)
    }
    ans := []int{nums[front()]}
    for i := k; i < len(nums); i++ {
        if front() == i-k {
            pop()
        }
        push(i)
        ans = append(ans, nums[front()])
    }
    return ans
}
