package main

import (
	"math"
	"testing"
)

func Test_reverse(t *testing.T) {
	num := 123
	if res := reverse(num); res != 321 {
		t.Errorf("expect 321 got %v", res)
	}

	num = -213
	if res := reverse(num); res != -312 {
		t.Errorf("expect -321 got %v", res)
	}

	num = 1
	if res := reverse(num); res != 1 {
		t.Errorf("expect 1 got %v", res)
	}

	num = 120
	if res := reverse(num); res != 21 {
		t.Errorf("expect 21 got %v", res)
	}

	num = 1534236469
	if res := reverse(num); res != 0 {
		t.Errorf("expect 0 got %v", res)
	}
}

func reverse(x int) int {
	res := 0

	for {
		y := x % 10
		x /= 10

		if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && y > 7) {
			return 0
		}
		if res < math.MinInt32/10 || (res == math.MinInt32/10 && y < -8) {
			return 0
		}

		res = res*10 + y

		if x == 0 {
			break
		}
	}

	return res
}
