package main

import (
	"testing"
)

func Test_hammingDistance(t *testing.T) {
	x, y, expect := 1, 4, 2
	if res := hammingDistance(x, y); res != expect {
		t.Fatalf("expect %v got %v", expect, res)
	}
}

func hammingDistance(x int, y int) int {
	foo := x ^ y
	count := 0

	for foo > 0 {
		if b := foo % 2; b == 1 {
			count++
		}

		foo >>= 1
	}

	return count
}
