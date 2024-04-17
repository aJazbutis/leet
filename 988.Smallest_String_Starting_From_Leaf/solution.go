/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func reverse(str []rune) string {
	s := make([]rune, len(str))
	for i, v := range str {
		s[len(str)-1-i] = v
	}
	return string(s)
}

func dfs(root *TreeNode, cur []rune, ret *string) {
	cur = append(cur, rune('a'+root.Val))
	if root.Left == nil && root.Right == nil {
		current := reverse(cur)
		if *ret == "" || *ret > current {
			*ret = current
		}
	} else {
		if root.Left != nil {
			dfs(root.Left, cur, ret)
		}
		if root.Right != nil {
			dfs(root.Right, cur, ret)
		}
	}
	cur = cur[:len(cur)-1]
}

func smallestFromLeaf(root *TreeNode) string {
	if root == nil {
		return ""
	}
	ret := ""
	dfs(root, []rune{}, &ret)
	return ret
}
