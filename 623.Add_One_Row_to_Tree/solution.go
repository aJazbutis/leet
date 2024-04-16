/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if root == nil {
		return root
	}
	if depth == 1 {
		return &TreeNode{Val: val, Left: root}
	}
	if depth == 2 {
		left, right := &TreeNode{Val: val, Left: root.Left}, &TreeNode{Val: val, Right: root.Right}
		root.Left, root.Right = left, right
		return root
	}
	addOneRow(root.Left, val, depth-1)
	addOneRow(root.Right, val, depth-1)
	return root
}
