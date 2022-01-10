package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	y := 0
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}
	for x != 0 {
		y = y*10 + x%10
		x = x / 10
	}
	if y > math.MaxInt32 || y < math.MinInt32 {
		return 0
	}
	return y
}

func main() {
	tests := []int{
		123,
		-123,
		120,
		0,
	}
	for _, test_in := range tests {
		test_out := reverse(test_in)
		fmt.Printf("%d 逆转为：%d\n", test_in, test_out)
	}
}
