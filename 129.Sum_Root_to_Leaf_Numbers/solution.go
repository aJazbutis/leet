/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sumNums(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return sum*10 + root.Val
	}
	return sumNums(root.Left, sum*10+root.Val) + sumNums(root.Right, sum*10+root.Val)
}

func sumNumbers(root *TreeNode) int {
	return sumNums(root, 0)
}
