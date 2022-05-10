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

func wiggleSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	var mid int = len(nums) / 2
	if len(nums)%2 != 0 {
		mid += 1
	}
	quickSelect(nums, 0, len(nums)-1, mid)
	var tmp1, tmp2 []int = make([]int, mid), make([]int, len(nums)-mid)
	copy(tmp1, nums[:mid])
	copy(tmp2, nums[mid:])
	for i, j := 0, 0; i+j < len(nums); {
		if (i+j)%2 == 0 {
			nums[i+j] = tmp1[len(tmp1)-1-i]
			i++
		} else {
			nums[i+j] = tmp2[len(tmp2)-1-j]
			j++
		}
	}
}

func main() {
	test_cases := [][]int{
		/*
			[]int{1, 5, 1, 1, 6, 4},
			[]int{1, 3, 2, 2, 3, 1},
		*/
		[]int{2},
		[]int{2, 1},
		[]int{1, 1, 2, 1, 2, 2, 1},
		[]int{4, 5, 5, 6},
	}
	for _, test_case := range test_cases {
		var prepare_test_case []int = make([]int, len(test_case))
		copy(prepare_test_case, test_case)
		wiggleSort(prepare_test_case)
		fmt.Printf("数组：%+v, 的摆动序列为：%v\n", test_case, prepare_test_case)
	}
}
