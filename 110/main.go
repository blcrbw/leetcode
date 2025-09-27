package main

// 110. Balanced Binary Tree
// Given a binary tree, determine if it is height-balanced.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, res := depth(root)
	return res
}

func depth(n *TreeNode) (int, bool) {
	if n == nil {
		return 0, true
	}

	left, valid := depth(n.Left)
	if !valid {
		return 0, false
	}
	right, valid := depth(n.Right)
	if !valid {
		return 0, false
	}

	if right > left+1 || right < left-1 {
		return 0, false
	}
	return max(right, left) + 1, true
}
