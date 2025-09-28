package main

// 101. Symmetric Tree
// Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)
}

func isMirror(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left != nil && right != nil && left.Val == right.Val && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left) {
		return true
	}
	return false
}
