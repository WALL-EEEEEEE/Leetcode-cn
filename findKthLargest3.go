package main

import (
	"fmt"
	"math/rand"
)

func pardition(nums []int, l, r int) int {
	/*左侧单向*/
	pivot := nums[r]
	var i, j int
	for i, j = l-1, l; j < r; j++ {
		if nums[j] < pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}

func pardition2(nums []int, left, right int) int {
	/*右侧单向*/
	pivot := nums[left]
	j := left
	for i := left + 1; i <= right; i++ {
		// 小于 pivot 的元素都被交换到前面
		if nums[i] < pivot {
			j++
			nums[j], nums[i] = nums[i], nums[j]
		}
	}
	// 在之前遍历的过程中，满足 [left+1, j] < pivot，并且 (j, i] >= pivot
	nums[j], nums[left] = nums[left], nums[j]
	// 交换以后 [left, j-1] < pivot, nums[j] = pivot, [j+1, right] >= pivot
	return j
}

func randPardition(nums []int, left, right int) int {
	var k int
	if right-left == 0 {
		k = left
	} else {
		k = rand.Intn(right-left) + left
	}
	nums[k], nums[right] = nums[right], nums[k]
	return pardition2(nums, left, right)
}

func quickSelect(nums []int, left, right, index int) int {
	pivot := randPardition(nums, left, right)
	for pivot != index {
		if pivot < index {
			pivot = randPardition(nums, pivot+1, right)
		} else {
			pivot = randPardition(nums, left, pivot-1)
		}
	}
	return nums[pivot]
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
			[]int{3, 2, 1, 5, 3, 6, 4},
			2,
		},
		[]interface{}{
			[]int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			4,
		},
		[]interface{}{
			[]int{
				1, 3, 5, 20, 0,
			},
			3,
		},
		[]interface{}{
			[]int{
				7, 9, 5, 20, 0, 6,
			},
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
