package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pardition(nums []int, l int, r int) int {
	size := r - l + 1
	rand.Seed(time.Now().Unix())
	pivot := rand.Intn(size) + l
	//fmt.Printf("size: %v, pivot: %v\n", size, pivot)
	//fmt.Printf("数组: %v, pivot: %v \n", nums, nums[pivot])
	nums[pivot], nums[r] = nums[r], nums[pivot]
	j := l
	for i := l; i < r; i++ {
		if nums[i] <= nums[r] {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	nums[j], nums[r] = nums[r], nums[j]
	//fmt.Printf("数组（分区后）：%v \n", nums)
	return j
}

func quickSelect(nums []int, l int, r int, k int) int {
	pivot := pardition(nums, l, r)
	/*
		for pivot != k {
			if pivot < k {
				pivot = pardition(nums, pivot+1, r)

			} else if pivot > k {
				pivot = pardition(nums, l, pivot-1)
			}
		}
		return nums[pivot]
	*/
	if pivot == k {
		return nums[pivot]
	}
	if pivot < k {
		return quickSelect(nums, pivot+1, r, k)
	}
	return quickSelect(nums, l, pivot-1, k)
}

func findKthLargest(nums []int, k int) int {
	k = len(nums) - k
	return quickSelect(nums, 0, len(nums)-1, k)
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
		[]interface{}{
			[]int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6},
			20,
		},
		[]interface{}{
			[]int{7, 6, 5, 4, 3, 2, 1},
			5,
		},
	}
	for _, test_case := range test_cases {
		nums := test_case.([]interface{})[0].([]int)
		k := test_case.([]interface{})[1].(int)
		kth := findKthLargest(nums, k)
		fmt.Printf("数组：%v， 第%v个最大值为：%v\n", nums, k, kth)
	}

}
