package main

import (
	"fmt"
	"math"
)

func maxProduct(nums []int) int {
	var ans int = math.MinInt32
	max_n, min_n := 1, 1
	len_nums := len(nums)
	for i := 0; i < len_nums; i++ {
		curr := nums[i]
		min_prod := min_n * curr
		max_prod := max_n * curr
		min_n = int(math.Min(math.Min(float64(max_prod), float64(curr)), float64(min_prod)))
		max_n = int(math.Max(math.Max(float64(max_prod), float64(curr)), float64(min_prod)))
		ans = int(math.Max(float64(max_n), float64(ans)))
	}
	return ans
}

func main() {
	var nums []int = []int{2, 3, -2, 4}
	//var nums []int = []int{2, 3}
	//var nums []int = []int{-2, 0, -1}
	//var nums []int = []int{-2}
	//var nums []int = []int{-1, -2, -9, -6}
	//max := maxProduct(nums)
	//max2 := maxProduct(nums2)
	//max3 := maxProduct(nums3)
	max := maxProduct(nums)
	//fmt.Println(max)
	//fmt.Println(max2)
	fmt.Println(max)
}
