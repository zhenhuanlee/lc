package main

import "testing"

func Test_game(t *testing.T) {
	g, a, e := []int{1, 2, 3}, []int{1, 2, 3}, 3
	if res := game(g, a); res != e {
		t.Fatalf("expect %v got %v", e, res)
	}
}

func game(guess []int, answer []int) int {
	sum := 0
	for i := 0; i < 3; i++ {
		sum += cpr(guess[i], answer[i])
	}
	return sum
}

func cpr(a, b int) int {
	if a == b {
		return 1
	}
	return 0
}
