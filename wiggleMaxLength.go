package main

import (
	"fmt"
)

func wiggleMaxLength(nums []int) int {
	var len_w int = 1
	var flag bool = false
	if len(nums) < 2 {
		return len(nums)
	}
	var end int = len(nums) - 1
	for i := end; i > 0; i-- {
		diff := nums[i] - nums[i-1]
		if i == end {
			flag = diff > 0
			if diff != 0 {
				len_w += 1
			} else {
				end = i - 1
			}
			continue
		}
		if diff == 0 {
			continue
		} else if flag != (diff > 0) {
			len_w += 1
			flag = diff > 0
			continue
		}
	}
	return len_w
}

func main() {
	test_cases := [][]int{
		/*
			[]int{1, 7, 4, 9, 2, 5},
			[]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{84},
			[]int{0, 0},
		*/
		[]int{1, 1, 1, 2, 2, 2, 1, 1, 1, 3, 3, 3, 2, 2, 2},
	}

	for _, test_case := range test_cases {
		wiggle_num := wiggleMaxLength(test_case)
		fmt.Printf("数组：%+v, 的最大摆动序列为：%d\n", test_case, wiggle_num)
	}
}
