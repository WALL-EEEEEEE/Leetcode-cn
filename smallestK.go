package main

import (
	"fmt"
)

func pardition(arr []int, l, r int) int {
	pivot := arr[r]
	var i, j int
	for i, j = l-1, l; j < r; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[r] = arr[r], arr[i+1]
	return i + 1
}

func smallestK(arr []int, k int) []int {
	left, right := 0, len(arr)-1
	var p, i, j int = left, left, right
	for p != k {
		p = pardition(arr, i, j)
		if p == k {
			return arr[:k]
		} else if p < k {
			i = p + 1
			j = right
		} else {
			j = p - 1
			i = left
		}

	}
	return arr[:k]
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
		kths := smallestK(nums, k)
		fmt.Printf("数组：%v， 最小的%v个值为：%v\n", nums, k, kths)
	}
}
