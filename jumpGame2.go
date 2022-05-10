package main

import "fmt"

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	steps := 0
	end := len(nums) - 1
	index := 0
	next_index := 0
	for index < end {
		step := nums[index]
		max_reach := 0
		for i := index + 1; i <= index+step && i < end; i++ {
			current_reach := nums[i] + i
			if current_reach == end {
				steps++
				return steps
			}
			if current_reach > max_reach {
				max_reach = current_reach
				next_index = i
			}
		}
		index = next_index
		steps++
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
		least_jump := jump(test_case)
		fmt.Printf("%+v 的最小跳跃数为: %v \n", test_case, least_jump)
	}
}
