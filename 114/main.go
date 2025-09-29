package main

// 114. Flatten Binary Tree to Linked List
// Given the root of a binary tree, flatten the tree into a "linked list":
//    The "linked list" should use the same TreeNode class where the right child pointer points to the next node in the
//    list and the left child pointer is always null.
//    The "linked list" should be in the same order as a pre-order traversal of the binary tree.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	_, _ = orderNodes(root)
}

func orderNodes(node *TreeNode) (cur, last *TreeNode) {
	if node == nil || (node.Right == nil && node.Left == nil) {
		return node, node
	}
	// Fix order of the right-hand side first.
	_, last = orderNodes(node.Right)
	// Then fix order of the left-hand side and place it between current node and right subnode.
	if node.Left != nil {
		left, leftLast := orderNodes(node.Left)
		leftLast.Right = node.Right
		leftLast.Left = nil
		if last == nil {
			last = leftLast
		}
		node.Right = left
	}
	node.Left = nil
	return node, last
}
