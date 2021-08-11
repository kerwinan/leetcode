package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var result []int
	middle(root, &result)
	return result
}

func middle(t *TreeNode, out *[]int) {
	if t == nil {
		return
	}
	middle(t.Left, out)
	*out = append(*out, t.Val)
	middle(t.Right, out)
	return
}

func main() {

}
