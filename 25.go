package main

import (
	"fmt"
)

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	var prev, curr, next *ListNode = nil, head, head.Next
	i := k

	for curr != nil && i > 0 {
		curr.Next = prev
		prev = curr
		curr = next
		if curr != nil {
			next = curr.Next
		}
		i--
	}

	if curr != nil {
		head.Next = reverseKGroup(curr, k)
	} else if i > 0 {
		return reverseKGroup(prev, k-i)
	}

	return prev
}

func main() {
	a := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	r := reverseKGroup(a, 3)

	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}
