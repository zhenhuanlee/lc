package main

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func Test_threeSumClosest(t *testing.T) {
	nums, target, expect := []int{-1, 2, 1, -4}, 1, 2
	if res := threeSumClosest(nums, target); res != 2 {
		t.Fatalf("expect %v got %v", expect, res)
	}
}

func threeSumClosest(nums []int, target int) int {
	len := len(nums)
	a, b, c, res := 0, 1, len-1, 0
	gap := math.MaxInt32

	sort.Ints(nums)

	for {
		foo := nums[a] + nums[b] + nums[c]

		fmt.Println(foo)

		if diff := abs(foo - target); diff < gap {
			gap = diff
			res = foo
		}

		if foo > target {
			c--
		} else {
			b++
		}

		if b == c {
			a++
			b = a + 1
			c = len - 1
			if a > len-3 {
				break
			}
		}
	}

	return res
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
