package main

func twoSum(nums []int, target int) []int {
	res := make(map[int]int)
	for i, n := range nums {
		idx, ok := res[n]
		if ok {
			return []int{idx, i}
		}
		res[target-n] = i
	}
	return nil
}
