package main

func findKthLargest(nums []int, k int) int {
	shift := 10001
	card := make([]int, 2*shift)
	for _, num := range nums {
		card[num+shift]++
	}
	for i := 2*shift - 1; i >= 0; i-- {
		k -= card[i]
		if k <= 0 {
			return i - shift
		}
	}
	return -1
}
