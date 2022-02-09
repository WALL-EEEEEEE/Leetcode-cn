package main

import (
	"fmt"
	"sort"
)

type Array struct {
	first       []int
	first_size  int
	second      []int
	second_size int
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	//首先将nums2数组，复制到nums1中，然后对nums1进行排序
	if m < 1 {
		copy(nums1, nums2)
	}
	copy(nums1[m:], nums2)
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
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
	}
	for _, test_case := range test_cases {
		fmt.Printf("数组1: %v, 数组2:%v, ", test_case.first, test_case.second)
		merge(test_case.first, test_case.first_size, test_case.second, test_case.second_size)
		fmt.Printf("合并数组：%v\n", test_case.first)
	}
}
