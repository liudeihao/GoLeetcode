package main

func flatten(root *TreeNode) {
	var prev *TreeNode
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		// 右、左、根 遍历。头插法
		if root == nil {
			return
		}
		dfs(root.Right)
		dfs(root.Left)
		root.Right = prev
		root.Left = nil
		prev = root
	}
	dfs(root)
}

/*
// 朴素想法：记录下来。
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	var list []*TreeNode
	var prevOrd func(*TreeNode)
	prevOrd = func(node *TreeNode) {
		list = append(list, node)
		if node.Left != nil {
			prevOrd(node.Left)
		}
		if node.Right != nil {
			prevOrd(node.Right)
		}
	}
	prevOrd(root)
	for i := 0; i < len(list)-1; i++ {
		list[i].Left = nil
		list[i].Right = list[i+1]
	}
}

*/

/*
可以用Moris遍历优化，以下是Moris遍历的代码（并非这题的答案）
// Moris遍历
cur := root
for cur != nil {
	if cur.Left == nil {
		cur = cur.Right // 对于prev.Right之前设置过了的，也不需要特殊处理，而是接下来敬请见证。
		continue
	}
	prev := cur.Left
	for prev.Right != nil && prev.Right != cur {
		prev = prev.Right
	}
	if prev.Right == nil {
		// 是第一次遍历cur
		prev.Right = cur
		cur = cur.Left
	} else { // 即prev.Right == cur
		// 是第二次遍历cur
		prev.Right = nil
		cur = cur.Right
	}
}
*/
