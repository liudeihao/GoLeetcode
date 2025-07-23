package main

func minCameraCover(root *TreeNode) int {
    var ans int
    var traverse func(*TreeNode) int
    traverse = func(r *TreeNode) int {
        // 0：该节点无覆盖
        // 1：本节点有摄像头
        // 2：本节点有覆盖
        if r == nil {
            return 2
        }
        left := traverse(r.Left)
        right := traverse(r.Right)
        if left == 2 && right == 2 {
            return 0
        }
        if left == 0 || right == 0 {
            ans++
            return 1
        }
        if left == 1 || right == 1 {
            return 2
        }
        return -1
    }
    if traverse(root) == 0 {
        ans++
    }
    return ans
}
