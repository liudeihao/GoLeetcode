package main

func findMode(root *TreeNode) []int {
	var ans []int
	mostCnt, curCnt := 1, 1
	var prev *TreeNode
	var trav func(node *TreeNode)
	trav = func(node *TreeNode) {
		if node == nil {
			return
		}

		trav(node.Left)
		if prev == nil {
			ans = []int{node.Val}
		} else {
			if prev.Val == node.Val {
				curCnt++
				if curCnt > mostCnt {
					mostCnt = curCnt
					ans = []int{node.Val}
				} else if curCnt == mostCnt {
					ans = append(ans, node.Val)
				}
			} else {
				curCnt = 1
				if mostCnt == 1 {
					ans = append(ans, node.Val)
				}
			}
		}
		prev = node
		trav(node.Right)
	}
	trav(root)

	return ans
}
