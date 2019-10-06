package main

import "fmt"

func main() {
	fmt.Println(isAnagram("abc", "cba"))
}

func isAnagram(s string, t string) bool {
	m := make(map[rune]int)

	for _, a := range s {
		m[a]++
	}

	for _, b := range t {
		m[b]--
	}

	for _, t := range m {
		if t != 0 {
			return false
		}
	}

	return true
}
