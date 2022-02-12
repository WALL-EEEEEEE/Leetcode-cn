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
	var pointer1, pointer2 int = m, n
	for pointer1 > 0 && pointer2 > 0 {
		if nums1[pointer1-1] <= nums2[pointer2-1] {
			nums1[pointer1+pointer2-1] = nums2[pointer2-1]
			pointer2--
		} else {
			nums1[pointer1+pointer2-1] = nums1[pointer1-1]
			pointer1--
		}
	}
	for pointer1 != 0 {
		nums1[pointer1-1] = nums1[pointer1-1]
		pointer1--
	}
	for pointer2 != 0 {
		nums1[pointer2-1] = nums2[pointer2-1]
		pointer2--
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
