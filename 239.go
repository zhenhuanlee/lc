package main

import "fmt"

type node struct {
	prev  *node
	val   int
	next  *node
	index int
}

func main() {
	// nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	nums := []int{1, -1}
	r := maxSlidingWindow(nums, 1)
	fmt.Println(r)
}

func maxSlidingWindow(nums []int, k int) []int {
	var head, tail *node
	res := []int{}

	for i, n := range nums {
		if head == nil {
			head = &node{val: n, index: i}
		}

		if tail == nil {
			tail = head
		} else {
			fmt.Println("xxx", n)
			tail.next = &node{
				prev:  tail,
				val:   n,
				index: i,
			}
			tail = tail.next
		}

		for tail.prev != nil && tail.val > tail.prev.val {
			if tail.prev.prev == nil {
				head = tail
				head.prev = nil
				break
			}
			p := tail.prev.prev
			p.next = tail
			tail.prev = p
		}

		if head.index+k <= i {
			if head.next == nil {
				head = &node{val: n, index: i}
			} else {
				head = head.next
				head.prev = nil
			}
		}

		if tail.index >= k-1 {
			res = append(res, head.val)
		}

		fmt.Println(i, n, head, tail)
	}

	return res
}
