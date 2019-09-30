package main

import (
	"fmt"
)

func main() {
	// a := &ListNode{
	// 	Val: 2,
	// 	Next: &ListNode{
	// 		Val: 4,
	// 		Next: &ListNode{
	// 			Val: 3,
	// 		},
	// 	},
	// }

	// b := &ListNode{
	// 	Val: 5,
	// 	Next: &ListNode{
	// 		Val: 6,
	// 		Next: &ListNode{
	// 			Val: 4,
	// 			// Next: &ListNode{
	// 			// 	Val: 9,
	// 			// },
	// 		},
	// 	},
	// }
	a := &ListNode{
		Val: 1,
	}
	b := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
		},
	}

	c := addTwoNumbers(a, b)

	for c != nil {
		fmt.Println(c.Val)
		c = c.Next
	}
}

// ListNode list node
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res, prev *ListNode = &ListNode{Val: 0}, nil
	head := res
	for l1 != nil || l2 != nil || prev.Val >= 10 {
		if l1 == nil {
			l1 = &ListNode{Val: 0}
		}
		if l2 == nil {
			l2 = &ListNode{Val: 0}
		}

		fmt.Println(prev)
		sum := l1.Val + l2.Val + res.Val
		foo, too := sum%10, sum/10

		res.Val = foo
		res.Next = &ListNode{Val: too}

		fmt.Println(l1, l2, res, prev)

		prev = res
		l1, l2, res = l1.Next, l2.Next, res.Next
	}

	if res.Val == 0 {
		prev.Next = nil
	}
	return head
}
