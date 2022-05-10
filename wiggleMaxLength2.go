package main

import (
	"fmt"
	"math"
)

func wiggleMaxLength2(nums []int) int {
	//动态规划解法
	num_size := len(nums)
	var up, down []int = make([]int, num_size), make([]int, num_size)
	if num_size < 2 {
		return len(nums)
	}
	up[0], down[0] = 1, 1
	for i := 1; i <= num_size-1; i++ {
		if nums[i] == nums[i-1] {
			down[i] = down[i-1]
			up[i] = up[i-1]
			continue
		}
		if nums[i] < nums[i-1] {
			down[i] = int(math.Max(float64(up[i-1]+1), float64(down[i-1])))
			up[i] = up[i-1]
		} else {
			up[i] = int(math.Max(float64(down[i-1]+1), float64(up[i-1])))
			down[i] = down[i-1]
		}
	}
	return int(math.Max(float64(up[num_size-1]), float64(down[num_size-1])))
}

func main() {
	test_cases := [][]int{
		[]int{1, 7, 4, 9, 2, 5},
		[]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]int{84},
		[]int{0, 0},
		[]int{1, 1, 1, 2, 2, 2, 1, 1, 1, 3, 3, 3, 2, 2, 2},
	}

	for _, test_case := range test_cases {
		wiggle_num := wiggleMaxLength2(test_case)
		fmt.Printf("数组：%+v, 的最大摆动序列为：%d\n", test_case, wiggle_num)
	}
}
