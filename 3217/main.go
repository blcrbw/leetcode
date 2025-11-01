package main

// 3217. Delete Nodes From Linked List Present in Array
//
// You are given an array of integers nums and the head of a linked list.
// Return the head of the modified linked list after removing all nodes
// from the linked list that have a value that exists in nums.

type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	l := make(map[int]interface{})
	for _, n := range nums {
		l[n] = nil
	}

	var res, prev *ListNode

	for {
		_, ok := l[head.Val]
		if !ok {
			if res == nil {
				res = head
			}
			if prev != nil {
				prev.Next = head
			}
			prev = head
		} else if prev != nil {
			prev.Next = head.Next
		}
		if head.Next == nil {
			break
		}
		head = head.Next

	}

	return res
}
