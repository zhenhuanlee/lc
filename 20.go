package main

import "fmt"

func main() {
	a := ""
	fmt.Println(isValid(a))
}

var pair = map[rune]rune{125: 123, 93: 91, 41: 40}

func isValid(s string) bool {
	stack := []rune{}

	for _, c := range s {
		switch c {
		case 123, 91, 40:
			stack = append(stack, c)
		default:
			l := len(stack) - 1
			if l < 0 {
				return false
			}
			last := stack[l]
			if last != pair[c] {
				return false
			}
			stack = stack[:l]
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}
