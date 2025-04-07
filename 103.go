package main

import (
	"slices"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	q := []*TreeNode{root}
	var ans [][]int
	reverse := false
	for len(q) > 0 {
		n := len(q)
		cur := make([]int, n)
		for i := range n {
			front := q[0]
			q = q[1:]
			cur[i] = front.Val
			if front.Left != nil {
				q = append(q, front.Left)
			}
			if front.Right != nil {
				q = append(q, front.Right)
			}
		}

		if reverse {
			slices.Reverse(cur)
		}
		reverse = !reverse
		ans = append(ans, cur)
	}
	return ans
}
