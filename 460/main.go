package main

// 430. Flatten a Multilevel Doubly Linked List
// You are given a doubly linked list, which contains nodes that have a next pointer, a previous pointer, and an additional child pointer. This child pointer may or may not point to a separate doubly linked list, also containing these special nodes. These child lists may have one or more children of their own, and so on, to produce a multilevel data structure as shown in the example below.
// Given the head of the first level of the list, flatten the list so that all the nodes appear in a single-level, doubly linked list. Let curr be a node with a child list. The nodes in the child list should appear after curr and before curr.next in the flattened list.
// Return the head of the flattened list. The nodes in the list must have all of their child pointers set to null.

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	cur := root
	for {
		if cur.Child != nil {
			pasteBetween(cur, cur.Child, cur.Next)
		}

		if cur.Next == nil {
			break
		} else {
			cur = cur.Next
		}
	}
	return root
}

func pasteBetween(prev, cur, next *Node) {
	cur.Prev = prev
	prev.Next = cur
	prev.Child = nil
	for {
		if cur.Child != nil {
			pasteBetween(cur, cur.Child, cur.Next)
		}
		if cur.Next != nil {
			cur = cur.Next
		} else {
			break
		}
	}
	cur.Next = next
	if next != nil {
		next.Prev = cur
	}
}

func test(data []int) []int {
	root := buildList(data)

	if root == nil {
		return nil
	}

	result := flatten(root)

	return toSlice(result)
}

func buildList(data []int) *Node {
	level := 0
	lIdx := 0
	var root *Node
	var prev *Node
	var first *Node
	var parent *Node
	for _, val := range data {
		if val > 0 {
			node := &Node{Val: val}
			if root == nil {
				root = node
			}
			if prev == nil {
				parent = first
				for lIdx > 0 {
					if parent.Next != nil {
						parent = parent.Next
					}
					lIdx--
				}
				if parent != nil {
					parent.Child = node
				}
				first = node
			} else {
				prev.Next = node
				node.Prev = prev
			}
			prev = node
		} else {
			if prev != nil {
				prev = nil
				level++
				lIdx = 0
			} else {
				lIdx++
			}

		}
	}
	return root
}

func toSlice(root *Node) []int {
	ret := make([]int, 0)
	ret = append(ret, root.Val)
	for root.Next != nil {
		root = root.Next
		ret = append(ret, root.Val)
	}
	return ret
}
