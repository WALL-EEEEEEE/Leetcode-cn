package main

import (
	"fmt"
	"math/rand"
)

func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

func pardition(nums []int, l, r int) int {
	pivot := nums[r]
	var i, j int
	for i, j = l-1, l; j <= r; j++ {
		if nums[j] <= pivot {
			i++
			swap(nums, i, j)
		}
	}
	return i
}

func randPardition(nums []int, left, right int) int {
	var k int
	if right-left == 0 {
		k = left
	} else {
		k = rand.Intn(right-left) + left
	}
	swap(nums, k, right)
	return pardition(nums, left, right)
}
func quickSelect(nums []int, left, right, index int) int {
	pivot := randPardition(nums, left, right)
	if pivot == index {
		return nums[pivot]
	} else if pivot < index {
		return quickSelect(nums, pivot+1, right, index)
	}
	return quickSelect(nums, left, pivot-1, index)
}

func findKthLargest(nums []int, k int) int {
	var size = len(nums)
	var left, right int = 0, size - 1
	k = size - k
	return quickSelect(nums, left, right, k)
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
