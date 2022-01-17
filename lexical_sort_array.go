package main

import "fmt"

func lexical_sort_array(nums []int) []int {
	var ans []int
	var max_num int = 0
	var num int = 1
	var nums_map map[int]bool = make(map[int]bool)
	len_num := len(nums)
	for _, v := range nums {
		nums_map[v] = true
		if max_num > v {
			continue
		}
		max_num = v
	}
	for len(ans) < len_num {
		for num <= max_num && num != 0 {
			if _, ok := nums_map[num]; ok {
				ans = append(ans, num)
			}
			num = num * 10
		}
		for num%10 == 9 || num > max_num {
			num = num / 10
		}
		num += 1
	}
	return ans
}

func main() {
	num_arrs := [][]int{
		[]int{1, 10, 11, 21, 13, 32},
		[]int{5, 1003, 11, 21, 13, 23},
	}
	for _, num_arr := range num_arrs {
		fmt.Printf("数组：%v, ", num_arr)
		num_arr = lexical_sort_array(num_arr)
		fmt.Printf("排序后：%v\n", num_arr)

	}

}
