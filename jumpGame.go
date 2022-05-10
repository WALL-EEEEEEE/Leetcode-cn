package main

import (
	"fmt"
	"math"
)

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	end := 0
	maxPosition := 0
	steps := 0
	for i, num := range nums {
		maxPosition = int(math.Max(float64(i+num), float64(maxPosition)))
		if i == end && i < len(nums)-1 {
			end = maxPosition
			steps++
		}
	}
	return steps
}

func main() {
	test_cases := [][]int{
		[]int{2, 3, 1, 1, 4},
		[]int{2, 3, 0, 1, 4},
		[]int{0},
		[]int{1, 2},
		[]int{1, 2, 3},
		[]int{1, 1, 1, 1},
	}
	for _, test_case := range test_cases {
		least_jump := jump2(test_case)
		fmt.Printf("%+v 的最小跳跃数为: %v \n", test_case, least_jump)
	}
}
