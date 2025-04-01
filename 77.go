package main

/*
所有回溯问题都遵循同一套代码框架：

def backtrack(路径, 选择列表):
	# print(f“当前路径: {路径}, 剩余选择: {选择列表}")
    if 满足终止条件:
        保存结果（如添加到结果列表）
        return

    for 选择 in 选择列表:
        if 选择不满足约束条件:  # 剪枝优化
            continue
        做选择（将选择加入路径）
        backtrack(新路径, 新选择列表)  # 递归
        撤销选择（从路径中移除）

*/

func combine(n int, k int) [][]int {
	var ans [][]int
	var path []int
	var backtrack func(int)
	backtrack = func(start int) {
		if len(path) == k {
			newCombine := make([]int, k)
			copy(newCombine, path)
			ans = append(ans, newCombine)
		}
		for i := start + 1; i <= n; i++ {
			path = append(path, i)
			backtrack(i)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return ans
}
