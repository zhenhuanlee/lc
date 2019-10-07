package main

import "fmt"

func main() {
	s := ""

	r := lengthOfLongestSubstring(s)
	fmt.Println(r)
}

func lengthOfLongestSubstring(s string) int {
	max, foo := 0, 0
	for i := range s {
		m := make(map[byte]int)

		for j := i; j < len(s); j++ {
			if m[s[j]] != 0 {
				foo = j - i
				goto NEXT
			}
			m[s[j]] = 1
		}
		foo = len(s) - i

	NEXT:

		if max < foo {
			max = foo
		}
	}

	return max
}
