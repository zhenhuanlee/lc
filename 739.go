package main

import "fmt"

func main() {
	temperatures := []int{23, 24, 25, 21, 19, 22, 26, 23}
	a := dailyTemperatures(temperatures)
	fmt.Println(a)
}

type listnode struct {
	tem   int
	index int
	next  *listnode
}

func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	var curr *listnode

	for i, t := range T {
		// when curr is nil
		if curr == nil {
			curr = &listnode{
				tem:   t,
				index: i,
			}
			continue
		}

		//
		for curr != nil && curr.tem < t {
			res[curr.index] = i - curr.index
			curr = curr.next
		}
		curr = &listnode{
			tem:   t,
			index: i,
			next:  curr,
		}

	}

	return res
}
