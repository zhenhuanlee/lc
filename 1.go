package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(nums, target))

}

func twoSum(nums []int, target int) []int {
	m := map[int]int{}

	for i, n := range nums {
		if m[target-n] == 0 {
			m[n] = i + 1
			continue
		}
		return []int{m[target-n] - 1, i}
	}

	return []int{}
}
