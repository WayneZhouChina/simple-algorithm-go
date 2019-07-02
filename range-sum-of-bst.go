package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}

	var ret int
	if root.Val < L {
		ret = rangeSumBST(root.Right, L, R)
	} else if root.Val > R {
		ret = rangeSumBST(root.Left, L, R)
	} else {
		ret = root.Val
		ret += rangeSumBST(root.Right, L, R)
		ret += rangeSumBST(root.Left, L, R)
	}

	return ret
}
