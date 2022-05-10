package main

import (
	"fmt"
	"math/rand"
)

func pardition(nums []int, left int, right int) int {
	size := right - left + 1
	pivot := rand.Intn(size) + left
	nums[pivot], nums[right] = nums[right], nums[pivot]
	j := left
	for i := left; i < right; i++ {
		if nums[i] <= nums[right] {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	nums[j], nums[right] = nums[right], nums[j]
	return j
}

func quickSelect(nums []int, left int, right int, k int) int {
	pivot := pardition(nums, left, right)
	if pivot == k {
		return pivot
	}
	if pivot < k {
		return quickSelect(nums, pivot+1, right, k)
	}
	return quickSelect(nums, left, pivot-1, k)
}

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) <= k {
		return arr
	}
	least_k := quickSelect(arr, 0, len(arr)-1, k)
	return arr[:least_k]
}

func main() {
	test_cases := []interface{}{
		[]interface{}{
			[]int{3, 2, 1},
			2,
		},
		[]interface{}{
			[]int{0, 1, 2, 1},
			1,
		},
		[]interface{}{
			[]int{0, 0, 2, 3, 2, 1, 1, 2, 0, 4},
			10,
		},
	}
	for _, test_case := range test_cases {
		nums := test_case.([]interface{})[0].([]int)
		k := test_case.([]interface{})[1].(int)
		smallest_kths := getLeastNumbers(nums, k)
		fmt.Printf("数组：%v， 最小的%v个数：%v\n", nums, k, smallest_kths)
	}

}
