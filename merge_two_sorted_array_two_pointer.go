package main

import (
	"fmt"
)

type Array struct {
	first       []int
	first_size  int
	second      []int
	second_size int
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	//双指针法，建立一个大小为m+n的临时数组，然后再采用双指针法，两个指针分别指向nums1和nums2的开头， 分别遍历
	var pointer1, pointer2 int = 0, 0
	if n < 1 {
		return
	}
	if m < 1 {
		copy(nums1, nums2)
		return
	}
	var temp []int = make([]int, m, m)
	copy(temp, nums1[:m])
	for pointer1 < m && pointer2 < n {
		if temp[pointer1] <= nums2[pointer2] {
			nums1[pointer1+pointer2] = temp[pointer1]
			pointer1++
		} else {
			nums1[pointer1+pointer2] = nums2[pointer2]
			pointer2++
		}
	}
	if pointer2 < n {
		for _, v := range nums2[pointer2:] {
			nums1[pointer1+pointer2] = v
			pointer2++
		}
	}
	if pointer1 < m {
		for _, v := range temp[pointer1:] {
			nums1[pointer1+pointer2] = v
			pointer1++
		}
	}
}

func main() {
	test_cases := []Array{
		Array{
			first:       []int{1, 2, 3, 0, 0, 0},
			first_size:  3,
			second:      []int{2, 5, 6},
			second_size: 3,
		},
		Array{
			first:       []int{1},
			first_size:  1,
			second:      []int{},
			second_size: 0,
		},
		Array{
			first:       []int{0},
			first_size:  0,
			second:      []int{1},
			second_size: 1,
		},
		Array{
			first:       []int{4, 0, 0, 0, 0, 0},
			first_size:  1,
			second:      []int{1, 2, 3, 5, 6},
			second_size: 5,
		},
		Array{
			first:       []int{2, 0},
			first_size:  1,
			second:      []int{1},
			second_size: 1,
		},
	}
	for _, test_case := range test_cases {
		fmt.Printf("数组1: %v, 数组2:%v, ", test_case.first, test_case.second)
		merge(test_case.first, test_case.first_size, test_case.second, test_case.second_size)
		fmt.Printf("合并数组：%v\n", test_case.first)
	}
}
