package main

import "fmt"

func main() {
	heights := []int{1, 8, 6, 2, 5, 100, 8, 3, 7}
	fmt.Println(maxArea(heights))
}

func maxArea(height []int) int {
	len := len(height)
	head, tail := 0, len-1
	max, area := 0, 0

	for head != tail {
		if height[head] > height[tail] {
			area = height[tail] * (tail - head)
			tail--
		} else {
			area = height[head] * (tail - head)
			head++
		}

		if area > max {
			max = area
		}
	}

	return max
}
