package main

import "fmt"

func main() {
	three := &TreeNode{Val: 3}
	two := &TreeNode{Val: 2, Left: three}
	root := &TreeNode{Val: 1, Right: two}
	fmt.Println(inorderTraversal(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 */
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		// get node from queue
		node := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		// if left exists
		if node.Left != nil {
			left := node.Left
			node.Left = nil
			// add node to the queue
			// add left to the queue
			queue = append(queue, node, left)
		} else {
			// process val
			result = append(result, node.Val)
			// process right
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return result
}

/**
 * Definition for a binary tree node.
 */
func inorderTraversalQ(root *TreeNode) []int {
	result := make([]int, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		// get node from queue
		node := queue[0]
		queue = queue[1:]
		// process val
		result = append(result, node.Val)
		// if left exists
		if node.Left != nil {
			// add left to the queue
			queue = append(queue, node.Left)
		}
		// process right
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}
