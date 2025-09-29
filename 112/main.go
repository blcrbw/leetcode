package main

// 112. Path Sum
// Given the root of a binary tree and an integer targetSum, return true if the tree has a root-to-leaf path such that
// adding up all the values along the path equals targetSum.
// A leaf is a node with no children.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return pathSum(0, root, &targetSum)
}

func pathSum(parentSum int, root *TreeNode, targetSum *int) bool {
	if root == nil {
		return false
	}

	sum := parentSum + root.Val
	if sum == *targetSum && root.Left == nil && root.Right == nil {
		return true
	}
	return pathSum(sum, root.Left, targetSum) || pathSum(sum, root.Right, targetSum)
}
