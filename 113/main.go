package main

// 113. Path Sum II
// Given the root of a binary tree and an integer targetSum, return all root-to-leaf paths where the sum of the node
// values in the path equals targetSum. Each path should be returned as a list of the node values, not node references.
//
// A root-to-leaf path is a path starting from the root and ending at any leaf node. A leaf is a node with no children.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	return checkPathSum(0, []int{}, root, &targetSum)
}

func checkPathSum(parentSum int, parentPath []int, node *TreeNode, targetSum *int) [][]int {
	var res [][]int
	if node == nil {
		return res
	}

	sum := parentSum + node.Val
	var path []int
	path = append(path, parentPath...)
	path = append(path, node.Val)
	if sum == *targetSum && node.Left == nil && node.Right == nil {
		return append(res, path)
	}
	res = append(res, checkPathSum(sum, path, node.Left, targetSum)...)
	res = append(res, checkPathSum(sum, path, node.Right, targetSum)...)
	return res
}
