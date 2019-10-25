package main

import (
	"testing"
)

func longestPalindrome(s string) string {
	len := len(s)
	if len < 2 {
		return s
	}
	res := [3]int{}

	for i := range s {
		f1 := getStartAndEnd(s, i, i)
		if i+1 < len && s[i] == s[i+1] {
			f2 := getStartAndEnd(s, i, i+1)
			if f2[0] > f1[0] {
				f1 = f2
			}
		}

		if f1[0] > res[0] {
			res = f1
		}
	}

	return s[res[1] : res[2]+1]
}

func getStartAndEnd(str string, s, e int) [3]int {
	for s > 0 && e+1 < len(str) && str[s-1] == str[e+1] {
		s--
		e++
	}

	return [3]int{e - s, s, e}
}

func TestLongestPalindrome(t *testing.T) {
	var (
		str string
		res string
	)
	str = "babad"
	res = longestPalindrome(str)
	if (res != "bab") && (res != "aba") {
		t.Fatalf("expect `%v` got `%v`", "bab", res)
	}

	str = "cbbd"
	res = longestPalindrome(str)
	if res != "bb" {
		t.Fatalf("expect `%v` got `%v`", "bb", res)
	}

	str = ""
	res = longestPalindrome(str)
	if res != "" {
		t.Fatalf("expect `%v` got `%v`", "", res)
	}

	str = "ccc"
	res = longestPalindrome(str)
	if res != "ccc" {
		t.Fatalf("expect `%v` got `%v`", "ccc", res)
	}
}
