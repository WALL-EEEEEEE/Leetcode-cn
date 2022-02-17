package main

import (
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	var left, right, pivot int = 0, len(nums) - 1, len(nums) - 1
	for left < right {
		for left < right && nums[left] < nums[pivot] {
			left++
		}
		tmp := nums[left]
		nums[left] = nums[pivot]
		nums[pivot] = tmp
	}

}

func main() {
	test_cases := []interface{}{
		[]interface{}{
			[]int{3, 2, 1, 5, 6, 4},
			2,
		},
		[]interface{}{
			[]int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			4,
		},
	}
	for _, test_case := range test_cases {
		nums := test_case.([]interface{})[0].([]int)
		k := test_case.([]interface{})[1].(int)
		kth := findKthLargest(nums, k)
		fmt.Printf("数组：%v， 第%v个最大值为：%v\n", nums, k, kth)
	}

}
