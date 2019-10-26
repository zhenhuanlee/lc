package main

import (
	"math"
	"testing"
)

func Test_myAtoi(t *testing.T) {
	m := map[string]int{
		// "42":                  42,
		"   -42":               -42,
		"4193 with words":      4193,
		"words and 987":        0,
		"-91283472332":         -2147483648,
		"   +0 123":            0,
		"9223372036854775808":  2147483647,
		"-9223372036854775808": -2147483648,
		"-   234":              0,
		"0-1":                  0,
	}

	for k, v := range m {
		if res := myAtoi(k); res != v {
			t.Fatalf("expect %v got %v", v, res)
		}
	}

}

func myAtoi(str string) int {
	foo, too, flag := 0, 0, 0

	for _, c := range str {
		switch c {
		case 32:
			if flag != 0 || too != 0 {
				goto END
			}
		case 43:
			if too != 0 || flag == 1 {
				goto END
			}
			too = 1
		case 45:
			if too != 0 || flag == 1 {
				goto END
			}
			too = -1
		case 48, 49, 50, 51, 52, 53, 54, 55, 56, 57:
			flag = 1
			foo = foo*10 + int(c) - 48
			if max := math.MaxInt32; foo > max && (too+1 > 0) {
				return max
			}
			if min := -math.MinInt32; foo > min && too == -1 {
				return -min
			}
		default:
			goto END
		}
	}

END:

	if too == 0 {
		too = 1
	}

	return too * foo
}
