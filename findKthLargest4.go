package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pardition(nums []int, l int, r int) int {
	size := r - l + 1
	rand.Seed(time.Now().Unix())
	pivot := rand.Intn(size)
	fmt.Printf("数组: %v, pivot: %v \n", nums, nums[pivot])
	nums[pivot], nums[r] = nums[r], nums[pivot]
	j := l
	for i := l; i < r; i++ {
		if nums[i] < nums[r] {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	nums[j], nums[r] = nums[r], nums[j]
	fmt.Printf("数组（分区后）：%v \n", nums)
	return j
}

func quickSelect(nums []int, l int, r int, k int) int {
	pivot := pardition(nums, l, r)
	k = len(nums) - k
	for pivot != k {
		if pivot < k {
			pivot = pardition(nums, pivot+1, r)

		} else if pivot > k {
			pivot = pardition(nums, l, pivot-1)
		}
	}
	return nums[pivot]
}

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, k)
}

func main() {
	test_cases := []interface{}{
		[]interface{}{
			[]int{3, 2, 1, 5, 6, 4},
			2,
		},
		/*
			[]interface{}{
				[]int{3, 2, 3, 1, 2, 4, 5, 5, 6},
				4,
			},
		*/
	}
	for _, test_case := range test_cases {
		nums := test_case.([]interface{})[0].([]int)
		k := test_case.([]interface{})[1].(int)
		kth := findKthLargest(nums, k)
		fmt.Printf("数组：%v， 第%v个最大值为：%v\n", nums, k, kth)
	}

}
