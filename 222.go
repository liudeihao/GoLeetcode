package main

import "sort"

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	}
	// 计算树的高度
	height := 0
	for node := root; node != nil; node = node.Left {
		height++
	}
	// 最大可能的节点编号（2^height - 1）
	maxNodes := 1 << height
	// 查找第一个不存在的节点编号
	firstMissing := sort.Search(maxNodes, func(k int) bool {
		if k == 0 { // 根节点始终存在
			return false
		}
		// 从根节点解码路径
		node := root
		bits := 1 << (height - 2) // 路径解码掩码（去掉最高位）
		for bits > 0 {
			if k&bits == 0 {
				node = node.Left
			} else {
				node = node.Right
			}
			bits >>= 1
		}
		return node == nil
	})
	return firstMissing - 1
}
