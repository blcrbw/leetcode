package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	var prev, item, next *ListNode
	item, next = head, head.Next
	item.Next = nil
	for {
		prev = item
		item = next
		next = next.Next
		item.Next = prev
		if next == nil {
			return item
		}
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	inp := []int{1, 2, 3, 6, 4, 4, 8, 7, 9, 9484, 2, 5, 6, 4}

	item := create(inp)

	result := reverseList(&item)
	fmt.Println(result)
}

func create(inp []int) ListNode {
	var next ListNode
	if len(inp) > 1 {
		next = create(inp[1:])
		return ListNode{Val: inp[0], Next: &next}
	} else {
		var val int
		if len(inp) > 0 {
			val = inp[0]
		}
		return ListNode{Val: val}
	}

}
